/* SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause) */
/* Copyright Authors of Kmesh */

#ifndef _KMESH_CONFIG_H_
#define _KMESH_CONFIG_H_

// map size
#define MAP_SIZE_OF_FRONTEND      105000
#define MAP_SIZE_OF_SERVICE       5000
#define MAP_SIZE_OF_ENDPOINT      105000
#define MAP_SIZE_OF_BACKEND       100000
#define MAP_SIZE_OF_AUTH          8192
#define MAP_SIZE_OF_DSTINFO       8192
#define MAP_SIZE_OF_AUTH_TAILCALL 100000
#define MAP_SIZE_OF_AUTH_POLICY   512

// map name
#define map_of_frontend       km_frontend
#define map_of_service        km_service
#define map_of_endpoint       km_endpoint
#define map_of_backend        km_backend
#define map_of_auth           km_auth
#define map_of_tuple          km_tuple
#define map_of_tcp_info       km_tcpInfo
#define map_of_authz          km_authz
#define map_of_dst_info       km_dstInfo
#define map_of_tail_call_prog km_tailCallProg
#define xdp_tailcall_map      km_xdptailcall
#define map_of_kmesh_socket   km_socket
#define kmesh_tc_args         km_tcargs
#define map_of_wl_policy      km_wlpolicy

#endif // _CONFIG_H_
