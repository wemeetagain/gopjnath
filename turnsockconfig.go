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

type TurnSockConfig struct {
    c C.struct_pj_turn_sock_cfg
}

func NewTurnSockConfig() *TurnSockConfig {
    c := &TurnSockConfig{}
    C.pj_turn_sock_cfg_default(c.c)
    return c
}

func (c *TurnSockConfig) SetMaxPacketSize(u uint) {
    c.c._max_pkg_size = C.uint(u)
}

func (c *TurnSockConfig) GetMaxPacketSize() uint {
    return uint(c.c._max_pkg_size)
}

//TODO  BoundAddr string

func (c *TurnSockConfig) SetPortRange(u uint16) {
    c.c._port_range = C.ushort(u)
}

func (c *TurnSockConfig) GetPortRange() uint16 {
    return uint16(c.c._port_range)
}

func (c *TurnSockConfig) SetQosType(u QosType) {
    c.c._qos_type = C.int(u)
}

func (c *TurnSockConfig) GetQosType() QosType {
    return QosType(c.c._qos_type)
}

func (c *TurnSockConfig) SetQosParams(u QosParams) {
    c.c._qos_params = u.p
}

func (c *TurnSockConfig) GetQosParams() QosParams {
    return QosParams{c.c._qos_params}
}

func (c *TurnSockConfig) SetQosIgnoreErr(u bool) {
    c.c._qos_ignore_error = C.int(u)
}

func (c *TurnSockConfig) GetQosIgnoreErr() bool {
    return bool(c.c._qos_ignore_error)
}

func (c *TurnSockConfig) SetRcvbufSize(u uint) {
    c.c._so_rcvbuf_size = C.uint(u)
}

func (c *TurnSockConfig) GetRcvbufSize() uint {
    return uint(c.c._so_rcvbuf_size)
}

func (c *TurnSockConfig) SetSndbufSize(u uint) {
    c.c._so_sndbuf_size = C.uint(u)
}

func (c *TurnSockConfig) GetSndbufSize() uint {
    return uint(c.c._so_sndbuf_size)
}
