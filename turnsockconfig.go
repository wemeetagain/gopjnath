package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

type TurnSockConfig struct {
    c *C.pj_turn_sock_cfg
}

func NewTurnSockConfig() *TurnSockConfig {
    c := &TurnSockConfig{}
    C.pj_turn_sock_cfg_default(c.c)
    return c
}

func (c *TurnSockConfig) SetMaxPacketSize(u uint) {
    c.c.max_pkt_size = C.uint(u)
}

func (c *TurnSockConfig) GetMaxPacketSize() uint {
    return uint(c.c.max_pkt_size)
}

func (c *TurnSockConfig) SetBoundAddr(s SockAddr) {
    c.c.bound_addr = s.s
}

func (c *TurnSockConfig) GetBoundAddr() *SockAddr {
    return &SockAddr{c.c.bound_addr}
}

func (c *TurnSockConfig) SetPortRange(u uint16) {
    c.c.port_range = C.pj_uint16_t(u)
}

func (c *TurnSockConfig) GetPortRange() uint16 {
    return uint16(c.c.port_range)
}

func (c *TurnSockConfig) SetQosType(u QosType) {
    c.c.qos_type = C.pj_qos_type(u)
}

func (c *TurnSockConfig) GetQosType() QosType {
    return QosType(c.c.qos_type)
}

/*
The pj_sock_set/get_qos_params() APIs are not portable, and it's probably only going to be implemented on Linux. Application should always try to use pj_sock_set_qos_type() instead. 
func (c *TurnSockConfig) SetQosParams(u QosParams) {
    c.c._qos_params = u.p
}

func (c *TurnSockConfig) GetQosParams() QosParams {
    return QosParams{c.c._qos_params}
}
*/

func (c *TurnSockConfig) SetQosIgnoreErr(b bool) {
    if b {
        c.c.qos_ignore_error = C.pj_bool_t(C.int(1))
    } else {
        c.c.qos_ignore_error = C.pj_bool_t(C.int(0))
    }
}

func (c *TurnSockConfig) GetQosIgnoreErr() bool {
    return int(c.c.qos_ignore_error) != 0
}

func (c *TurnSockConfig) SetRcvbufSize(u uint) {
    c.c.so_rcvbuf_size = C.uint(u)
}

func (c *TurnSockConfig) GetRcvbufSize() uint {
    return uint(c.c.so_rcvbuf_size)
}

func (c *TurnSockConfig) SetSndbufSize(u uint) {
    c.c.so_sndbuf_size = C.uint(u)
}

func (c *TurnSockConfig) GetSndbufSize() uint {
    return uint(c.c.so_sndbuf_size)
}
