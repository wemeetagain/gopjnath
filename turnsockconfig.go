package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

// TurnSockConfig describes options that can be specified when creating
// the TURN socket. Application should call NewTurnSockConfig to
// initialize this structure with its default values before using it. 
type TurnSockConfig struct {
    c C.pj_turn_sock_cfg
}

func NewTurnSockConfig() TurnSockConfig {
    c := TurnSockConfig{}
    C.pj_turn_sock_cfg_default(&c.c)
    return c
}

// Packet buffer size. Default value is PJ_TURN_MAX_PKT_LEN.
func (c *TurnSockConfig) MaxPacketSize() uint {
    return uint(c.c.max_pkt_size)
}

func (c *TurnSockConfig) SetMaxPacketSize(u uint) {
    c.c.max_pkt_size = C.uint(u)
}

// Specify the interface where the socket should be bound to. If the
// address is zero, socket will be bound to INADDR_ANY. If the address
// is non-zero, socket will be bound to this address only. If the port
// is set to zero, the socket will bind at any port (chosen by the OS). 
func (c *TurnSockConfig) BoundAddr() *SockAddr {
    return &SockAddr{&c.c.bound_addr}
}

func (c *TurnSockConfig) SetBoundAddr(s SockAddr) {
    c.c.bound_addr = *s.s
}

// Specify the port range for TURN socket binding, relative to the start
// port number specified in bound_addr. Note that this setting is only
// applicable when the start port number is non zero. Default value is
// zero.
func (c *TurnSockConfig) PortRange() uint16 {
    return uint16(c.c.port_range)
}

func (c *TurnSockConfig) SetPortRange(u uint16) {
    c.c.port_range = C.pj_uint16_t(u)
}

// QoS traffic type to be set on this transport. When application wants
// to apply QoS tagging to the transport, it's preferable to set this
// field rather than qos_param fields since this is more portable.
// Default value is PJ_QOS_TYPE_BEST_EFFORT.
func (c *TurnSockConfig) QosType() QosType {
    return QosType(c.c.qos_type)
}

func (c *TurnSockConfig) SetQosType(u QosType) {
    c.c.qos_type = C.pj_qos_type(u)
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

// Specify if STUN socket should ignore any errors when setting the QoS
// traffic type/parameters. Default: true
func (c *TurnSockConfig) QosIgnoreErr() bool {
    return int(c.c.qos_ignore_error) != 0
}

func (c *TurnSockConfig) SetQosIgnoreErr(b bool) {
    if b {
        c.c.qos_ignore_error = C.pj_bool_t(C.int(1))
    } else {
        c.c.qos_ignore_error = C.pj_bool_t(C.int(0))
    }
}

// Specify target value for socket receive buffer size. It will be
// applied using setsockopt(). When it fails to set the specified size,
// it will try with lower value until the highest possible has been
// successfully set. Default: 0 (OS default)
func (c *TurnSockConfig) RcvbufSize() uint {
    return uint(c.c.so_rcvbuf_size)
}

func (c *TurnSockConfig) SetRcvbufSize(u uint) {
    c.c.so_rcvbuf_size = C.uint(u)
}

// Specify target value for socket send buffer size. It will be applied
// using setsockopt(). When it fails to set the specified size, it will
// try with lower value until the highest possible has been successfully
// set. Default: 0 (OS default)
func (c *TurnSockConfig) SndbufSize() uint {
    return uint(c.c.so_sndbuf_size)
}

func (c *TurnSockConfig) SetSndbufSize(u uint) {
    c.c.so_sndbuf_size = C.uint(u)
}

