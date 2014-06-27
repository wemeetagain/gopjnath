package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

type TurnTransportType int

const (
    TurnTransportUdp = TurnTransportType(C.PJ_TURN_TP_UDP)
    TurnTransportTcp = TurnTransportType(C.PJ_TURN_TP_TCP)
    TurnTransportTls = TurnTransportType(C.PJ_TURN_TP_TLS)
    )

type IceTransportConfig struct {
    t *C.pj_ice_strans_cfg
}

// void pj_ice_strans_cfg_default (pj_ice_strans_cfg *cfg)
func NewIceTransportConfig(c *Context) *IceTransportConfig {
    var cfg C.pj_ice_strans_cfg
    C.pj_ice_strans_cfg_default(&cfg)
    C.pj_stun_config_init(&cfg.stun_cfg,&c.cp.factory,0,c.io,c.tHeap)
    tc := IceTransportConfig{&cfg}
    return &tc
}

// void pj_ice_strans_cfg_copy (pj_pool_t *pool, pj_ice_strans_cfg *dst, const pj_ice_strans_cfg *src)
func (tc *IceTransportConfig) Copy() {
}

func (tc *IceTransportConfig) Destroy() {
}

// struct getters/setters

// int af
func (tc *IceTransportConfig) GetAf() int {
    return int(tc.t.af)
}

/* Currently only pj_AF_INET() (IPv4) is supported, and this is the default value. 
func (tc *IceTransportConfig) SetAf(i int) {
    tc.t.af = C.int(i)
}
*/

// pj_stun_config stun_cfg
func (tc *IceTransportConfig) StunConfig() StunConfig {
    return StunConfig{tc.t.stun_cfg}
}

// pj_dns_resolver * resolver

// pj_ice_sess_options opt
func (tc *IceTransportConfig) IceSessOptions() IceSessOptions {
    return IceSessOptions{tc.t.opt}
}

//// stun

// pj_stun_sock_cfg cfg
func (tc *IceTransportConfig) StunSockConfig() StunSockConfig {
    return StunSockConfig{tc.t.stun.cfg}
}

// unsigned max_host_cands
func (tc *IceTransportConfig) GetStunMaxHostCands() uint {
    return uint(tc.t.stun.max_host_cands)
}

func (tc *IceTransportConfig) SetStunMaxHostCands(u uint) {
    tc.t.stun.max_host_cands = C.uint(u)
}

// pj_bool_t loop_addr
func (tc *IceTransportConfig) GetStunLoopAddr() bool {
    return int(tc.t.stun.loop_addr) != 0
}

func (tc *IceTransportConfig) SetStunLoopAddr(b bool) {
    if b {
        tc.t.stun.loop_addr = C.pj_bool_t(C.int(1))
    } else {
        tc.t.stun.loop_addr = C.pj_bool_t(C.int(0))
    }
}

// pj_str_t server
func (tc *IceTransportConfig) GetStunServer() string {
    str := toString(tc.t.stun.server)
    return str
}

func (tc *IceTransportConfig) SetStunServer(s string) {
    destroyString(tc.t.stun.server)
    str := C.CString(s)
    tc.t.stun.server = C.pj_str(str)
}

// pj_uint16_t port
func (tc *IceTransportConfig) GetStunPort() uint16 {
    return uint16(tc.t.stun.port)
}

func (tc *IceTransportConfig) SetStunPort(u uint16) {
    tc.t.stun.port = C.pj_uint16_t(u)
}

// pj_bool_t ignore_stun_error
func (tc *IceTransportConfig) GetStunIgnoreStunError() bool {
    return int(tc.t.stun.ignore_stun_error) != 0
}

func (tc *IceTransportConfig) SetStunIgnoreStunError(b bool) {
    if b {
        tc.t.stun.ignore_stun_error = C.pj_bool_t(C.int(1))
    } else {
        tc.t.stun.ignore_stun_error = C.pj_bool_t(C.int(0))
    }
}

//// turn

//pj_turn_sock_cfg cfg
func (tc *IceTransportConfig) TurnSockConfig() TurnSockConfig {
    return TurnSockConfig{tc.t.turn.cfg}
}

// pj_str_t server
func (tc *IceTransportConfig) GetTurnServer() string {
    str := toString(tc.t.turn.server)
    return str}

func (tc *IceTransportConfig) SetTurnServer(s string) {
    destroyString(tc.t.turn.server)
    str := C.CString(s)
    tc.t.turn.server = C.pj_str(str)
}

// pj_uint16_t port
func (tc *IceTransportConfig) GetTurnPort() uint16 {
    return uint16(tc.t.turn.port)
}

func (tc *IceTransportConfig) SetTurnPort(u uint16) {
    tc.t.turn.port = C.pj_uint16_t(u)
}

// pj_turn_tp_type conn_type
func (tc *IceTransportConfig) GetTurnConnType() TurnTransportType {
    return TurnTransportType(tc.t.turn.conn_type)
}

func (tc *IceTransportConfig) SetTurnConnType(t TurnTransportType) {
    tc.t.turn.conn_type = C.pj_turn_tp_type(t)
}

// pj_stun_auth_cred auth_cred

// pj_turn_alloc_param alloc_param

//// comp[PJ_ICE_MAX_COMP]

// pj_qos_type qos_type

// pj_qos_params qos_params

// unsigned so_rcvbuf_size

// unsigned so_sndbuf_size
