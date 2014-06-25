package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

type StunSockConfig struct {
    c C.pj_stun_sock_cfg
}

func NewStunSockConfig() StunSockConfig {
    c := StunSockConfig{}
    C.pj_stun_sock_cfg_default(&c.c)
    return c
}

func (c *StunSockConfig) SetMaxPacketSize(u uint) {
    c.c.max_pkt_size = C.uint(u)
}

func (c *StunSockConfig) GetMaxPacketSize() uint {
    return uint(c.c.max_pkt_size)
}

func (c *StunSockConfig) SetAsyncCount(u uint) {
    c.c.async_cnt = C.uint(u)
}

func (c *StunSockConfig) GetAsyncCount() uint {
    return uint(c.c.async_cnt)
}

func (c *StunSockConfig) SetBoundAddr(s SockAddr) {
    c.c.bound_addr = s.s
}

func (c *StunSockConfig) GetBoundAddr() *SockAddr {
    return &SockAddr{c.c.bound_addr}
}

func (c *StunSockConfig) SetPortRange(u uint16) {
    c.c.port_range = C.pj_uint16_t(u)
}

func (c *StunSockConfig) GetPortRange() uint16 {
    return uint16(c.c.port_range)
}

func (c *StunSockConfig) SetKaInterval(u int) {
    c.c.ka_interval = C.int(u)
}

func (c *StunSockConfig) GetKaInterval() int {
    return int(c.c.ka_interval)
}

func (c *StunSockConfig) SetQosType(u QosType) {
    c.c.qos_type = C.pj_qos_type(u)
}

func (c *StunSockConfig) GetQosType() QosType {
    return QosType(c.c.qos_type)
}

/*
The pj_sock_set/get_qos_params() APIs are not portable, and it's probably only going to be implemented on Linux. Application should always try to use pj_sock_set_qos_type() instead. 
func (c *StunSockConfig) SetQosParams(u QosParams) {
    c.c._qos_params = u.p
}

func (c *StunSockConfig) GetQosParams() QosParams {
    return QosParams{c.c._qos_params}
}
*/

func (c *StunSockConfig) SetQosIgnoreErr(b bool) {
    if b {
        c.c.qos_ignore_error = C.pj_bool_t(C.int(1))
    } else {
        c.c.qos_ignore_error = C.pj_bool_t(C.int(0))
    }
}

func (c *StunSockConfig) GetQosIgnoreErr() bool {
    return int(c.c.qos_ignore_error) != 0
}

func (c *StunSockConfig) SetRcvbufSize(u uint) {
    c.c.so_rcvbuf_size = C.uint(u)
}

func (c *StunSockConfig) GetRcvbufSize() uint {
    return uint(c.c.so_rcvbuf_size)
}

func (c *StunSockConfig) SetSndbufSize(u uint) {
    c.c.so_sndbuf_size = C.uint(u)
}

func (c *StunSockConfig) GetSndbufSize() uint {
    return uint(c.c.so_sndbuf_size)
}
