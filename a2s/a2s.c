/*
 * a2s.c - skeleton vpp engine plug-in
 *
 * Copyright (c) <current-year> <your-organization>
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <vnet/vnet.h>
#include <vnet/plugin/plugin.h>
#include <a2s/a2s.h>

#include <vlibapi/api.h>
#include <vlibmemory/api.h>
#include <vpp/app/version.h>
#include <stdbool.h>

#include <a2s/a2s.api_enum.h>
#include <a2s/a2s.api_types.h>

#define REPLY_MSG_ID_BASE amp->msg_id_base
#include <vlibapi/api_helper_macros.h>
#include <vppinfra/bihash_template.h>
#include <vppinfra/bihash_template.c>
a2s_main_t a2s_main;

/* Action function shared between message handler and debug CLI */

int a2s_enable_disable (a2s_main_t * amp, u32 sw_if_index,
                                   int enable_disable)
{
  vnet_sw_interface_t * sw;
  int rv = 0;

  /* Utterly wrong? */
  if (pool_is_free_index (amp->vnet_main->interface_main.sw_interfaces,
                          sw_if_index))
    return VNET_API_ERROR_INVALID_SW_IF_INDEX;

  /* Not a physical port? */
  sw = vnet_get_sw_interface (amp->vnet_main, sw_if_index);
  if (sw->type != VNET_SW_INTERFACE_TYPE_HARDWARE)
    return VNET_API_ERROR_INVALID_SW_IF_INDEX;

  vnet_feature_enable_disable ("ip4-unicast", "a2s",
                               sw_if_index, enable_disable, 0, 0);

  return rv;
}

static clib_error_t *
a2s_enable_disable_command_fn (vlib_main_t * vm,
                                   unformat_input_t * input,
                                   vlib_cli_command_t * cmd)
{
  a2s_main_t * amp = &a2s_main;
  u32 sw_if_index = ~0;
  int enable_disable = 1;

  int rv;

  while (unformat_check_input (input) != UNFORMAT_END_OF_INPUT)
    {
      if (unformat (input, "disable"))
        enable_disable = 0;
      else if (unformat (input, "%U", unformat_vnet_sw_interface,
                         amp->vnet_main, &sw_if_index))
        ;
      else
        break;
  }

  if (sw_if_index == ~0)
    return clib_error_return (0, "Please specify an interface...");

  rv = a2s_enable_disable (amp, sw_if_index, enable_disable);

  switch(rv)
    {
  case 0:
    break;

  case VNET_API_ERROR_INVALID_SW_IF_INDEX:
    return clib_error_return
      (0, "Invalid interface, only works on physical ports");
    break;

  case VNET_API_ERROR_UNIMPLEMENTED:
    return clib_error_return (0, "Device driver doesn't support redirection");
    break;

  default:
    return clib_error_return (0, "a2s_enable_disable returned %d",
                              rv);
    }
  return 0;
}

/* *INDENT-OFF* */
VLIB_CLI_COMMAND (a2s_enable_disable_command, static) =
{
  .path = "a2s enable-disable",
  .short_help =
  "a2s enable-disable <interface-name> [disable]",
  .function = a2s_enable_disable_command_fn,
};
/* *INDENT-ON* */

/* API message handler */
static void vl_api_a2s_enable_disable_t_handler
(vl_api_a2s_enable_disable_t * mp)
{
  clib_warning("what");
  vl_api_a2s_enable_disable_reply_t * rmp;
  a2s_main_t * amp = &a2s_main;
  int rv;

  rv = a2s_enable_disable (amp, ntohl(mp->sw_if_index),
                                      (int) (mp->enable_disable));

  REPLY_MACRO(VL_API_A2S_ENABLE_DISABLE_REPLY);
}

static void vl_api_a2s_set_data_t_handler
(vl_api_a2s_set_data_t *mp)
{
  vlib_main_t *vm = vlib_get_main();
  vl_api_a2s_set_data_reply_t * rmp;
  a2s_main_t * amp = &a2s_main;
  int rv;

  clib_bihash_kv_a2s_t kv;
  if(mp->is_add)
  {
    memcpy(&kv.key,&mp->kv.key,sizeof(a2s_key_t));
    kv.key.ip = clib_host_to_net_u32(kv.key.ip);
    kv.value.info_length = clib_net_to_host_u16(mp->kv.value.info_length);
    kv.value.player_length = clib_net_to_host_u16(mp->kv.value.player_length);
    kv.value.rules_length = clib_net_to_host_u16(mp->kv.value.rules_length);

    if (kv.value.info_length > 2000 || kv.value.player_length > 2000 | kv.value.rules_length > 10000)
      return;

    clib_memcpy(&kv.value.INFO_DATA,&mp->kv.value.INFO_DATA, kv.value.info_length);
    clib_memcpy(&kv.value.PLAYER_DATA,&mp->kv.value.PLAYER_DATA, kv.value.player_length);
    clib_memcpy(&kv.value.RULES_DATA,&mp->kv.value.RULES_DATA, kv.value.rules_length);

    kv.value.strict = mp->kv.value.strict;
    kv.value.last = vlib_time_now(vm);
  }

  rv = clib_bihash_add_del_a2s(&amp->cache,&kv,mp->is_add);


  REPLY_MACRO(VL_API_A2S_SET_DATA_REPLY);
}

/* API definitions */
#include <a2s/a2s.api.c>

static clib_error_t * a2s_init (vlib_main_t * vm)
{
  a2s_main_t * amp = &a2s_main;
  clib_error_t * error = 0;

  clib_warning("hi");


  amp->vlib_main = vm;
  amp->vnet_main = vnet_get_main();

  /* Add our API messages to the global name_crc hash table */
  amp->msg_id_base = setup_message_id_table ();
   
  clib_bihash_a2s_t *cache;
  cache = &amp->cache;
  clib_bihash_init_a2s(cache,"a2s cache",(u32)1000,(uword)(20000));

  clib_time_init(&amp->clib_time);

  return error;
}

VLIB_INIT_FUNCTION (a2s_init);

/* *INDENT-OFF* */
VNET_FEATURE_INIT (a2s, static) =
{
  .arc_name = "ip4-unicast",
  .node_name = "a2s",
  .runs_before = VNET_FEATURES ("ip4-flow-classify"),
};
/* *INDENT-ON */

/* *INDENT-OFF* */
VLIB_PLUGIN_REGISTER () =
{
  .version = VPP_BUILD_VER,
  .description = "A2S query cache",
};
/* *INDENT-ON* */

/*
 * fd.io coding-style-patch-verification: ON
 *
 * Local Variables:
 * eval: (c-set-style "gnu")
 * End:
 */
