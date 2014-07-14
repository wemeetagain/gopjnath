package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

// StunSockConfig describes the settings to be given to the STUN
// transport during its creation. Application should initialize this
// structure by calling NewStunSockConfig
type StunSockConfig struct {
    c C.pj_stun_sock_cfg
}

func NewStunSockConfig() StunSockConfig {
    c := StunSockConfig{}
    C.pj_stun_sock_cfg_default(&c.c)
    return c
}

// Packet buffer size. Default value is PJ_STUN_SOCK_PKT_LEN.
func (c *StunSockConfig) MaxPacketSize() uint {
    return uint(c.c.max_pkt_size)
}

func (c *StunSockConfig) SetMaxPacketSize(u uint) {
    c.c.max_pkt_size = C.uint(u)
}

// Specify the number of simultaneous asynchronous read operations to be
// invoked to the ioqueue. Having more than one read operations will
// increase performance on multiprocessor systems since the application
// will be able to process more than one incoming packets
// simultaneously. Default value is 1. 
func (c *StunSockConfig) AsyncCount() uint {
    return uint(c.c.async_cnt)
}

func (c *StunSockConfig) SetAsyncCount(u uint) {
    c.c.async_cnt = C.uint(u)
}

// Specify the interface where the socket should be bound to. If the
// address is zero, socket will be bound to INADDR_ANY. If the address
// is non-zero, socket will be bound to this address only, and the
// transport will have only one address alias (the alias_cnt field in
// pj_stun_sock_info structure. If the port is set to zero, the socket
// will bind at any port (chosen by the OS).
func (c *StunSockConfig) BoundAddr() *SockAddr {
    return &SockAddr{&c.c.bound_addr}
}

func (c *StunSockConfig) SetBoundAddr(s SockAddr) {
    c.c.bound_addr = *s.s
}

// Specify the port range for STUN socket binding, relative to the start
// port number specified in bound_addr. Note that this setting is only
// applicable when the start port number is non zero. Default value is zero.
func (c *StunSockConfig) PortRange() uint16 {
    return uint16(c.c.port_range)
}

func (c *StunSockConfig) SetPortRange(u uint16) {
    c.c.port_range = C.pj_uint16_t(u)
}

// Specify the STUN keep-alive duration, in seconds. The STUN transport
// does keep-alive by sending STUN Binding request to the STUN server.
// If this value is zero, the PJ_STUN_KEEP_ALIVE_SEC value will be used.
// If the value is negative, it will disable STUN keep-alive. 
func (c *StunSockConfig) KaInterval() int {
    return int(c.c.ka_interval)
}

func (c *StunSockConfig) SetKaInterval(u int) {
    c.c.ka_interval = C.int(u)
}

// QoS traffic type to be set on this transport. When application wants
// to apply QoS tagging to the transport, it's preferable to set this
// field rather than qos_param fields since this is more portable. 
// Default value is PJ_QOS_TYPE_BEST_EFFORT.
func (c *StunSockConfig) QosType() QosType {
    return QosType(c.c.qos_type)
}

func (c *StunSockConfig) SetQosType(u QosType) {
    c.c.qos_type = C.pj_qos_type(u)
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

// Specify if STUN socket should ignore any errors when setting the QoS
// traffic type/parameters. Default: true
func (c *StunSockConfig) QosIgnoreErr() bool {
    return int(c.c.qos_ignore_error) != 0
}

func (c *StunSockConfig) SetQosIgnoreErr(b bool) {
    if b {
        c.c.qos_ignore_error = C.pj_bool_t(C.int(1))
    } else {
        c.c.qos_ignore_error = C.pj_bool_t(C.int(0))
    }
}

// Specify target value for socket receive buffer size. It will be
// applied using setsockopt(). When it fails to set the specified size,
// it will try with lower value until the highest possible is
// successfully set. Default: 0 (OS default)
func (c *StunSockConfig) RcvbufSize() uint {
    return uint(c.c.so_rcvbuf_size)
}

func (c *StunSockConfig) SetRcvbufSize(u uint) {
    c.c.so_rcvbuf_size = C.uint(u)
}

// Specify target value for socket send buffer size. It will be applied
// using setsockopt(). When it fails to set the specified size, it will
// try with lower value until the highest possible is successfully set.
// Default: 0 (OS default)
func (c *StunSockConfig) SndbufSize() uint {
    return uint(c.c.so_sndbuf_size)
}

func (c *StunSockConfig) SetSndbufSize(u uint) {
    c.c.so_sndbuf_size = C.uint(u)
}
