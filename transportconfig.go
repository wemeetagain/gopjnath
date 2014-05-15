package gopjnath

/*
#cgo pkg-config: libpjnath
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/

import "C"

import (
    "sync"
    "syscall"
    "time"
    "unsafe"
    )

type TurnTransportType int

const (
    TurnTransportUdp TurnTransportType(C.PJ_TURN_TP_UDP)
    TurnTransportTcp TurnTransportType(C.PJ_TURN_TP_TCP)
    TurnTransportTls TurnTransportType(C.PJ_TURN_TP_TLS)
    )

type TransportConfig struct {
    t C.struct_pj_ice_strans_cfg
}

// void pj_ice_strans_cfg_default (pj_ice_strans_cfg *cfg)
func NewTransportConfig() *TransportConfig {
    tc := &TransportConfig{}
    C.pj_ice_strans_cfg_default(tc.t)
    return tc
}

// void pj_ice_strans_cfg_copy (pj_pool_t *pool, pj_ice_strans_cfg *dst, const pj_ice_strans_cfg *src)

// struct getters/setters

// int af
func (tc *TransportConfig) GetAf() int {
    return int(tc.t.af)
}

func (tc *TransportConfig) SetAf(i int) {
    tc.t.af = C.int(i)
}

// pj_stun_config stun_cfg

// pj_dns_resolver * resolver

// pj_ice_sess_options opt

//// stun

// pj_stun_sock_cfg cfg
 
// unsigned max_host_cands
func (tc *TransportConfig) GetStunMaxHostCands() uint {
    return uint(tc.t.stun.max_host_cands)
}

func (tc *TransportConfig) SetStunMaxHostCands(u uint) {
    tc.t.stun.max_host_cands = C.uint(u)
}

// pj_bool_t loop_addr
func (tc *TransportConfig) GetStunLoopAddr() bool {
    return bool(tc.t.stun.loop_addr)
}

func (tc *TransportConfig) SetStunLoopAddr(b bool) {
    tc.t.stun.loop_addr = C.pj_bool_t(b)
}

// pj_str_t server
func (tc *TransportConfig) GetStunServer() string {
    return string(tc.t.stun.server)
}

func (tc *TransportConfig) SetStunServer(s string) {
    tc.t.stun.server = C.pj_str_t(s)
}

// pj_uint16_t port
func (tc *TransportConfig) GetStunPort() uint16 {
    return uint16(tc.t.stun.port)
}

func (tc *TransportConfig) SetStunPort(u uint16) {
    tc.t.stun.port = C.pj_uint16_t(u)
}

// pj_bool_t ignore_stun_error
func (tc *TransportConfig) GetStunIgnoreStunError() bool {
    return bool(tc.t.stun.ignore_stun_error)
}

func (tc *TransportConfig) SetStunIgnoreStunError(b bool) {
    tc.t.stun.ignore_stun_error = C.pj_bool_t(b)
}

//// turn

//pj_turn_sock_cfg cfg

// pj_str_t server
func (tc *TransportConfig) GetTurnServer() string {
    return string(tc.t.turn.server)
}

func (tc *TransportConfig) SetTurnServer(s string) {
    tc.t.turn.server = C.pj_str_t(s)
}

// pj_uint16_t port
func (tc *TransportConfig) GetTurnPort() uint16 {
    return uint16(tc.t.turn.port)
}

func (tc *TransportConfig) SetTurnPort(u uint16) {
    tc.t.turn.port = C.pj_uint16_t(u)
}

// pj_turn_tp_type conn_type
func (tc *TransportConfig) GetTurnConnType() TurnTransportType {
    return TurnTransportType(tc.t.turn.conn_type)
}

func (tc *TransportConfig) SetTurnConnType(t TurnTransportType) {
    tc.t.turn.conn_type = C.pj_turn_tp_type(t)
}

// pj_stun_auth_cred auth_cred

// pj_turn_alloc_param alloc_param

//// comp[PJ_ICE_MAX_COMP]

// pj_qos_type qos_type

// pj_qos_params qos_params

// unsigned so_rcvbuf_size

// unsigned so_sndbuf_size



