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

type StunSockConfig struct {
    c C.struct_pj_stun_sock_cfg
}

func NewStunSockConfig() *StunSockConfig {
    c := &StunSockConfig{}
    C.pj_stun_sock_cfg_default(c.c)
    return c
}

func (c *StunSockConfig) SetMaxPacketSize(u uint) {
    c.c._max_pkg_size = C.uint(u)
}

func (c *StunSockConfig) GetMaxPacketSize() uint {
    return uint(c.c._max_pkg_size)
}

func (c *StunSockConfig) SetAsyncCount(u uint) {
    c.c._async_cnt = C.uint(u)
}

func (c *StunSockConfig) GetAsyncCount() uint {
    return uint(c.c._async_cnt)
}

//TODO  BoundAddr string

func (c *StunSockConfig) SetPortRange(u uint16) {
    c.c._port_range = C.ushort(u)
}

func (c *StunSockConfig) GetPortRange() uint16 {
    return uint16(c.c._port_range)
}

func (c *StunSockConfig) SetKaInterval(u int) {
    c.c._ka_interval = C.int(u)
}

func (c *StunSockConfig) GetKaInterval() int {
    return int(c.c._ka_interval)
}

func (c *StunSockConfig) SetQosType(u QosType) {
    c.c._qos_type = C.int(u)
}

func (c *StunSockConfig) GetQosType() QosType {
    return QosType(c.c._qos_type)
}

func (c *StunSockConfig) SetQosParams(u QosParams) {
    c.c._qos_params = u.p
}

func (c *StunSockConfig) GetQosParams() QosParams {
    return QosParams{c.c._qos_params}
}

func (c *StunSockConfig) SetQosIgnoreErr(u bool) {
    c.c._qos_ignore_error = C.int(u)
}

func (c *StunSockConfig) GetQosIgnoreErr() bool {
    return bool(c.c._qos_ignore_error)
}

func (c *StunSockConfig) SetRcvbufSize(u uint) {
    c.c._so_rcvbuf_size = C.uint(u)
}

func (c *StunSockConfig) GetRcvbufSize() uint {
    return uint(c.c._so_rcvbuf_size)
}

func (c *StunSockConfig) SetSndbufSize(u uint) {
    c.c._so_sndbuf_size = C.uint(u)
}

func (c *StunSockConfig) GetSndbufSize() uint {
    return uint(c.c._so_sndbuf_size)
}
