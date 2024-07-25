/*
 * a2s.c - a2s vpp-api-test plug-in
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
#include <vat/vat.h>
#include <vlibapi/api.h>
#include <vlibmemory/api.h>
#include <vppinfra/error.h>
#include <stdbool.h>

#define __plugin_msg_base a2s_test_main.msg_id_base
#include <vlibapi/vat_helper_macros.h>

uword unformat_sw_if_index (unformat_input_t * input, va_list * args);

/* Declare message IDs */
#include <a2s/a2s.api_enum.h>
#include <a2s/a2s.api_types.h>

typedef struct
{
  /* API message ID base */
  u16 msg_id_base;
  vat_main_t *vat_main;
} a2s_test_main_t;

a2s_test_main_t a2s_test_main;

static int api_a2s_enable_disable (vat_main_t * vam)
{
  unformat_input_t * i = vam->input;
  int enable_disable = 1;
  u32 sw_if_index = ~0;
  vl_api_a2s_enable_disable_t * mp;
  int ret;

  /* Parse args required to build the message */
  while (unformat_check_input (i) != UNFORMAT_END_OF_INPUT)
    {
      if (unformat (i, "%U", unformat_sw_if_index, vam, &sw_if_index))
          ;
        else if (unformat (i, "sw_if_index %d", &sw_if_index))
          ;
      else if (unformat (i, "disable"))
          enable_disable = 0;
      else
          break;
    }

  if (sw_if_index == ~0)
    {
      errmsg ("missing interface name / explicit sw_if_index number \n");
      return -99;
    }

  /* Construct the API message */
  M(A2S_ENABLE_DISABLE, mp);
  mp->sw_if_index = ntohl (sw_if_index);
  mp->enable_disable = enable_disable;

  /* send it... */
  S(mp);

  /* Wait for a reply... */
  W (ret);
  return ret;
}

static int api_a2s_set_data(vat_main_t *vam)
{
    clib_warning("hi");

  unformat_input_t * i = vam->input;
  vl_api_a2s_set_data_t * mp;
  int ret;
  u32 ip = ~0;
  u16 port = ~0;
  u8 *info_data = 0, *player_data = 0, *rules_data = 0;
  u16 info_length = 0, player_length = 0, rules_length = 0;
  bool strict = false;

  /* Parse args required to build the message */
  while (unformat_check_input (i) != UNFORMAT_END_OF_INPUT)
    {
      if (unformat (i, "ip %U", unformat_ip4_address, &ip))
          ;
      else if (unformat (i, "port %d", &port))
          ;
      else if (unformat (i, "info_data %v", &info_data))
          info_length = vec_len(info_data);
      else if (unformat (i, "player_data %v", &player_data))
          player_length = vec_len(player_data);
      else if (unformat (i, "rules_data %v", &rules_data))
          rules_length = vec_len(rules_data);
      else if (unformat (i, "strict"))
          strict = true;
      else
          break;
    }

  if (ip == ~0 || port == (u16)~0)
    {
      errmsg ("missing ip address / port number \n");
      return -99;
    }

  /* Construct the API message */
  M(A2S_SET_DATA, mp);
  mp->kv.key.ip = ntohl(ip);
  mp->kv.key.port = ntohs(port);
  clib_memcpy(mp->kv.value.INFO_DATA, info_data, info_length);
  clib_memcpy(mp->kv.value.PLAYER_DATA, player_data, player_length);
  clib_memcpy(mp->kv.value.RULES_DATA, rules_data, rules_length);
  mp->kv.value.info_length = ntohs(info_length);
  mp->kv.value.player_length = ntohs(player_length);
  mp->kv.value.rules_length = ntohs(rules_length);
  mp->kv.value.strict = strict;

  /* send it... */
  S(mp);

  /* Wait for a reply... */
  W (ret);
  return ret;
}


/*
 * List of messages that the a2s test plugin sends,
 * and that the data plane plugin processes
 */
#include <a2s/a2s.api_test.c>

/*
 * fd.io coding-style-patch-verification: ON
 *
 * Local Variables:
 * eval: (c-set-style "gnu")
 * End:
 */
