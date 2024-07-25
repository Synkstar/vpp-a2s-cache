#undef BIHASH_TYPE
#undef BIHASH_KVP_PER_PAGE
#undef BIHASH_32_64_SVM
#undef BIHASH_ENABLE_STATS
#undef BIHASH_KVP_AT_BUCKET_LEVEL
#undef BIHASH_LAZY_INSTANTIATE
#undef BIHASH_BUCKET_PREFETCH_CACHE_LINES
#undef BIHASH_USE_HEAP

// Define your custom bihash type and parameters
#define BIHASH_TYPE _a2s
#define BIHASH_KVP_PER_PAGE 7
#define BIHASH_KVP_AT_BUCKET_LEVEL 1
#define BIHASH_LAZY_INSTANTIATE 0
#define BIHASH_BUCKET_PREFETCH_CACHE_LINES 2
#define BIHASH_USE_HEAP 1

#ifndef __included_bihash_a2s_h__
#define __included_bihash_a2s_h__

#include <vppinfra/heap.h>
#include <vppinfra/format.h>
#include <vppinfra/pool.h>
#include <vppinfra/xxhash.h>
#include <vppinfra/crc32.h>
#define A2S_SIZE 2000
#define A2S_RULES_SIZE 10000

// Define the key and value structures
typedef struct {
    u32 ip;
    u16 port;
} a2s_key_t;

typedef struct {
    f64 last;
    unsigned char INFO_DATA[A2S_SIZE];
    unsigned char PLAYER_DATA[A2S_SIZE];
    unsigned char RULES_DATA[A2S_RULES_SIZE];
    u16 info_length;
    u16 player_length;
    u16 rules_length;
    bool strict;
} a2s_data_t;

typedef struct {
    a2s_key_t key;
    a2s_data_t value;
} clib_bihash_kv_a2s_t;

// Define necessary utility functions
static inline void
clib_bihash_mark_free_a2s(clib_bihash_kv_a2s_t *v)
{
    v->value.last = 0;
}

static inline int
clib_bihash_is_free_a2s(clib_bihash_kv_a2s_t *v)
{
    return v->value.last == 0;
}

static inline u64
clib_bihash_hash_a2s(clib_bihash_kv_a2s_t *v)
{
    u64 tmp = ((u64)v->key.ip << 16) | v->key.port;
    return clib_xxhash(tmp);
}

static inline u8 *
format_bihash_kvp_a2s(u8 *s, va_list *args)
{
    clib_bihash_kv_a2s_t *v = va_arg(*args, clib_bihash_kv_a2s_t *);

    s = format(s, "key ip: %u port: %u value last: %llu data: %s",
               v->key.ip, v->key.port, v->value.last, v->value.INFO_DATA);
    return s;
}

static inline int
clib_bihash_key_compare_a2s(a2s_key_t a, a2s_key_t b)
{
    return (a.ip == b.ip && a.port == b.port);
}

// Include the bihash template
#undef __included_bihash_template_h__
#include <vppinfra/bihash_template.h>

#endif /* __included_bihash_a2s_h__ */
