/*
 * node.c - Basic A2S Cache program
 */

#include <vlib/vlib.h>
#include <vnet/vnet.h>
#include <vnet/pg/pg.h>
#include <vppinfra/error.h>
#include <vnet/ip/ip.h>
#include <vnet/udp/udp.h>
#include <vnet/ethernet/ethernet.h>
#include <vppinfra/hash.h>

#include <a2s/a2s.h>

typedef struct {
    u32 next_index;
    u32 sw_if_index;
} a2s_trace_t;

typedef enum
{
  A2S_DROP,
  A2S_NEXT_NODE,
  A2S_N_NEXT,
} a2s_reply_next_t;

#ifndef CLIB_MARCH_VARIANT

/* packet trace format function */
static u8 *format_a2s_trace(u8 *s, va_list *args) {
    s = format(s, "a2s: sw_if_index %d, next_index %d",
               va_arg(*args, u32), va_arg(*args, u32));
    return s;
}

vlib_node_registration_t a2s_node;

#endif /* CLIB_MARCH_VARIANT */

#define foreach_a2s_error \
_(PROCESSED, "A2S packets processed")

static const u8 a2s_info_query[] = {
    0xFF, 0xFF, 0xFF, 0xFF, 0x54, 0x53, 0x6F, 0x75, 0x72, 0x63, 0x65, 0x20, 0x45, 0x6E, 0x67, 0x69
};

static const u8 a2s_player_request[] = {
    0xFF, 0xFF, 0xFF, 0xFF, 0x55, 0x00, 0x00, 0x00, 0x00
};

static const u8 a2s_rules_request[] = {
    0xFF, 0xFF, 0xFF, 0xFF, 0x56, 0x00, 0x00, 0x00, 0x00
};

static const u8 challenge_response[] = {
    0xFF,0xFF,0xFF,0xFF,0x41
};

typedef enum {
#define _(sym,str) A2S_ERROR_##sym,
    foreach_a2s_error
#undef _
    A2S_N_ERROR,
} a2s_error_t;


static inline u32 create_cookie(clib_time_t clib_time, udp_header_t *udp0) {
    // Use the lower 8 bits of the total_cpu_time
    u32 time_part = (u32)((u64)(clib_time.total_cpu_time / clib_time.seconds_per_clock) & 0xFF);

    // Combine src_port and dst_port
    u32 combined_ports = (udp0->src_port << 16) | udp0->dst_port;

    // Combine ports with the lower 8 bits of time
    u32 combined = time_part + combined_ports;

    return combined;
}

static inline bool check_cookie(clib_time_t clib_time, udp_header_t *udp0, u32 cookie) {
    // Combine dst_port and src_port into a single 32-bit value
    u32 combined_ports = (udp0->src_port << 16) | udp0->dst_port;

    // Subtract the combined ports from the cookie to retrieve the time part
    u32 cookie_time_part = cookie - combined_ports;

    // Extract the current time part
    u32 current_time_part = (u32)((u64)(clib_time.total_cpu_time / clib_time.seconds_per_clock) & 0xFF);

    // Calculate the time difference, accounting for wraparound
    u32 time_difference = current_time_part - cookie_time_part;

    // Check if the time difference is within 5 seconds
    return time_difference <= 5;
}

static inline void send_challenge(a2s_main_t *mp, vlib_buffer_t *b0) {
    ip4_header_t *ip0 = vlib_buffer_get_current(b0);
    udp_header_t *udp0 = ip4_next_header(ip0);
    u8 *payload = (u8 *)(udp0 + 1);
    u16 host_payload_length = clib_net_to_host_u16(udp0->length) - 8;

    // Swap ip addresses
    u32 tmp_ip = ip0->src_address.as_u32;
    ip0->src_address.as_u32 = ip0->dst_address.as_u32;
    ip0->dst_address.as_u32 = tmp_ip;

    // Set other values
    ip0->ttl = 64;
    ip0->length = clib_host_to_net_u16 (sizeof(ip4_header_t) + sizeof(udp_header_t) + 9);
    ip0->checksum = ip4_header_checksum(ip0);

    // Change size of the packet so its big enough for a 9 byte payload only
    vlib_buffer_chain_increase_length(b0, b0, 9 - host_payload_length);

    // Create cookie and append it to the challenge response header
    u32 cookie = create_cookie(mp->clib_time,udp0);
    clib_memcpy(payload,challenge_response,sizeof(challenge_response));
    clib_memcpy(payload+5,&cookie,4);

    // Swap ports
    u16 tmp_port = udp0->src_port;
    udp0->src_port = udp0->dst_port;
    udp0->dst_port = tmp_port;

    // Set other values
    udp0->checksum = 0;
    udp0->length = clib_host_to_net_u16(sizeof(udp_header_t) + 9);

    vnet_buffer(b0)->sw_if_index[VLIB_TX] = 0;
}

static inline void send_data(a2s_main_t *mp, vlib_main_t *vm, vlib_buffer_t *b0, u8 *data, u16 length) {
    ip4_header_t *ip0 = vlib_buffer_get_current(b0);
    udp_header_t *udp0 = ip4_next_header(ip0);
    u8 *payload = (u8 *)(udp0 + 1);
    u32 buffer_size = vlib_buffer_get_default_data_size(vm);
    vlib_buffer_t *last_b0 = b0;

    // Swap ip addresses
    u32 tmp_ip = ip0->src_address.as_u32;
    ip0->src_address.as_u32 = ip0->dst_address.as_u32;
    ip0->dst_address.as_u32 = tmp_ip;

    // Set other values
    ip0->ttl = 64;
    ip0->flags_and_fragment_offset = 0;
    ip0->length = clib_host_to_net_u16(sizeof(ip4_header_t) + sizeof(udp_header_t) + length);
    ip0->checksum = ip4_header_checksum(ip0);

    // Swap ports
    u16 tmp_port = udp0->src_port;
    udp0->src_port = udp0->dst_port;
    udp0->dst_port = tmp_port;

    // Set other values
    udp0->checksum = 0;
    udp0->length = clib_host_to_net_u16(sizeof(udp_header_t) + length);

    // Copy initial data into the first buffer
    u32 bytes_to_copy = clib_min(length, buffer_size - sizeof(ip4_header_t) - sizeof(udp_header_t));
    clib_memcpy(payload, data, bytes_to_copy);
    vlib_buffer_chain_increase_length(b0, last_b0, bytes_to_copy);
    length -= bytes_to_copy;
    data += bytes_to_copy;

    // Allocate and chain additional buffers if necessary
    while (length > 0) {
        vlib_buffer_t *new_b0;
        u32 new_bi;
        
        if (vlib_buffer_alloc(vm, &new_bi, 1) != 1) {
            // Handle allocation failure
            clib_warning("Buffer allocation failed!");
            return;
        }

        new_b0 = vlib_get_buffer(vm, new_bi);
        new_b0->current_data = 0;
        new_b0->current_length = 0;
        new_b0->flags |= VLIB_BUFFER_TOTAL_LENGTH_VALID;

        last_b0->next_buffer = new_bi;
        last_b0->flags |= VLIB_BUFFER_NEXT_PRESENT;
        last_b0 = new_b0;

        bytes_to_copy = clib_min(length, buffer_size - sizeof(ip4_header_t) - sizeof(udp_header_t));
        clib_memcpy(vlib_buffer_get_current(new_b0), data, bytes_to_copy);
        vlib_buffer_chain_increase_length(b0, last_b0, bytes_to_copy);
        new_b0->current_length = bytes_to_copy;

        length -= bytes_to_copy;
        data += bytes_to_copy;
    }

    vnet_buffer(b0)->sw_if_index[VLIB_TX] = 0;
}

VLIB_NODE_FN(a2s_node)(vlib_main_t *vm,
                       vlib_node_runtime_t *node,
                       vlib_frame_t *frame) {
    u32 n_left_from, *from, *to_next;
    a2s_reply_next_t next_index;
    a2s_main_t *mp = &a2s_main;

    from = vlib_frame_vector_args(frame);
    n_left_from = frame->n_vectors;
    next_index = node->cached_next_index;

    while (n_left_from > 0) {
        u32 n_left_to_next;
        vlib_get_next_frame(vm, node, next_index, to_next, n_left_to_next);

        while (n_left_from > 0 && n_left_to_next > 0) {
            u32 bi0;
            vlib_buffer_t *b0;
            u32 next0 = A2S_NEXT_NODE; // Default to forward
            f64 now = vlib_time_now(vm);
            /* speculatively enqueue b0 to the current next frame */
            to_next[0] = bi0 = from[0];
            from++;
            to_next++;
            n_left_from--;
            n_left_to_next--;

            b0 = vlib_get_buffer(vm, bi0);
        
            // Extract IPv4 Header
            ip4_header_t *ip0 = vlib_buffer_get_current(b0);
            if (ip0->protocol == IP_PROTOCOL_UDP) {
                udp_header_t *udp0 = ip4_next_header(ip0);

                // Initialize bihash to search
                clib_bihash_kv_a2s_t search_kv, return_kv;
                search_kv.key.ip = ip0->dst_address.as_u32;
                search_kv.key.port = udp0->dst_port;

                // Search bihash
                if(clib_bihash_search_a2s(&mp->cache,&search_kv,&return_kv) < 0)
                    goto bypass;

                // Drop requests for expired servers
                if(return_kv.value.last - 10 > now) {
                    next0 = A2S_DROP;
                    goto bypass;
                }

                // Check for specific payload
                u8 *payload = (u8 *)(udp0 + 1);
                u16 host_payload_length = clib_net_to_host_u16(udp0->length) - 8;

                // Check for A2S_PLAYER or A2S_RULES challenge requests
                if (host_payload_length == 9 && (clib_memcmp(payload, a2s_player_request, sizeof(a2s_player_request)) == 0 || clib_memcmp(payload, a2s_rules_request, sizeof(a2s_rules_request)) == 0)) {
                   send_challenge(mp,b0);
                   goto end;
                }

                // Check for A2S_PLAYER request with challenge
                if (host_payload_length == 9 && clib_memcmp(payload, a2s_player_request, 5) == 0 ) {
                    u32 challenge;
                    clib_memcpy(&challenge, payload+5, 4);
                    if(!check_cookie(mp->clib_time,udp0,challenge)) {
                        next0 = A2S_DROP;
                        goto bypass;
                    }

                    send_data(mp,vm,b0,return_kv.value.PLAYER_DATA,return_kv.value.player_length);
                    goto end;
                } 

                // Check for A2S_RULES request with challenge
                if (host_payload_length == 9 && clib_memcmp(payload, a2s_rules_request, 5) == 0 ) {
                    u32 challenge;
                    clib_memcpy(&challenge, payload+5, 4);
                    if(!check_cookie(mp->clib_time,udp0,challenge)) {
                        next0 = A2S_DROP;
                        goto bypass;
                    }

                    send_data(mp,vm,b0,return_kv.value.RULES_DATA,return_kv.value.rules_length);
                    goto end;
                } 

                // Check for A2S_INFO request without challenge
                if (host_payload_length == 25 && clib_memcmp(payload, a2s_info_query, sizeof(a2s_info_query)) == 0) {
                   send_challenge(mp,b0);
                   goto end;
                }

                // Check for A2S_INFO query with challenge
                if (host_payload_length == 29 && clib_memcmp(payload, a2s_info_query, sizeof(a2s_info_query)) == 0) {
                    u32 challenge;
                    clib_memcpy(&challenge, payload+25, 4);
                    if(!check_cookie(mp->clib_time,udp0,challenge)) {
                        next0 = A2S_DROP;
                        goto bypass;
                    }

                    send_data(mp,vm,b0,return_kv.value.INFO_DATA,return_kv.value.info_length);
                    goto end;
                }

                if(return_kv.value.strict)
                    next0 = A2S_DROP;
            }

            goto bypass;
            
            end: 
            b0->flags |= VNET_BUFFER_F_LOCALLY_ORIGINATED;

            bypass:
            // Enqueue the buffer to the next node
            vlib_validate_buffer_enqueue_x1(vm, node, next_index, to_next, n_left_to_next, bi0, next0);
        }

        vlib_put_next_frame(vm, node, next_index, n_left_to_next);
    }

    return frame->n_vectors;
}
/* *INDENT-OFF* */
#ifndef CLIB_MARCH_VARIANT
VLIB_REGISTER_NODE(a2s_node) = {
    .name = "a2s",
    .vector_size = sizeof(u32),
    .format_trace = format_a2s_trace,
    .type = VLIB_NODE_TYPE_INTERNAL,
    .flags = VLIB_NODE_FLAG_TRACE,

    .n_next_nodes = A2S_N_NEXT,

    /* edit / add dispositions here */
    .next_nodes = {
        [A2S_DROP] = "error-drop",
        [A2S_NEXT_NODE] = "ip4-lookup",
    },
};
#endif /* CLIB_MARCH_VARIANT */
/* *INDENT-ON* */
/*
 * fd.io coding-style-patch-verification: ON
 *
 * Local Variables:
 * eval: (c-set-style "gnu")
 * End:
 */
