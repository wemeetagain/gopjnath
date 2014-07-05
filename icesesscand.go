package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

type IceCandType int

const (
    IceCandTypeHost    = IceCandType(C.PJ_ICE_CAND_TYPE_HOST)
    IceCandTypeSrFlx   = IceCandType(C.PJ_ICE_CAND_TYPE_SRFLX)
    IceCandTypePrFlx   = IceCandType(C.PJ_ICE_CAND_TYPE_PRFLX)
    IceCandTypeRelayed = IceCandType(C.PJ_ICE_CAND_TYPE_RELAYED)
    )

// IceSessCand describes an ICE candidate. ICE candidate is a transport
// address that is to be tested by ICE procedures in order to determine
// its suitibility for usage for receipt of media. Candidates also have
// properties -- their type, priority, foundation, and base.
type IceSessCand struct {
    c *C.pj_ice_sess_cand
}

// The candidate type.
func (c *IceSessCand) Type() IceCandType {
    return IceCandType(c.c._type)
}

func (c *IceSessCand) SetType(t IceCandType) {
    c.c._type = C.pj_ice_cand_type(t)
}

// Status of this candidate. The value will be nil if candidate address
// has been resolved successfully, PJ_EPENDING when the address
// resolution process is in progress, or other value when the address
// resolution has completed with failure. 
func (c *IceSessCand) Status() error {
    return casterr(c.c.status)
}

// The component ID of this candidate. Note that component IDs starts
// with one for RTP and two for RTCP. In other words, it's not zero
// based. 
func (c *IceSessCand) ComponentId() uint8 {
    return uint8(c.c.comp_id)
}

func (c *IceSessCand) SetComponentId(i uint8) {
    c.c.comp_id = C.pj_uint8_t(i)
}

// Transport ID to be used to send packets for this candidate. 
func (c *IceSessCand) TransportId() uint8 {
    return uint8(c.c.transport_id)
}

func (c *IceSessCand) SetTransportId(i uint8) {
    c.c.transport_id = C.pj_uint8_t(i)
}

// Local preference value, which typically is 65535.
func (c *IceSessCand) LocalPref() uint16 {
    return uint16(c.c.local_pref)
}

func (c *IceSessCand) SetLocalPref(i uint16) {
    c.c.local_pref = C.pj_uint16_t(i)
}

// The foundation string, which is an identifier which value will be
// equivalent for two candidates that are of the same type, share the
// same base, and come from the same STUN server. The foundation is used
// to optimize ICE performance in the Frozen algorithm. 
func (c *IceSessCand) Foundation() string {
    str := toString(c.c.foundation)
    return str
}

func (c *IceSessCand) SetFoundation(s string) {
    destroyString(c.c.foundation)
    str := C.CString(s)
    c.c.foundation = C.pj_str(str)
}

// The candidate's priority, a 32-bit unsigned value which value will be
// calculated by the ICE session when a candidate is registered to the
// ICE session. 
func (c *IceSessCand) Priority() uint32 {
    return uint32(c.c.prio)
}

func (c *IceSessCand) SeetPriority(i uint32) {
    c.c.prio = C.pj_uint32_t(i)
}

// IP address of this candidate. For host candidates, this represents
// the local address of the socket. For reflexive candidates, the value
// will be the public address allocated in NAT router for the host
// candidate and as reported in MAPPED-ADDRESS or XOR-MAPPED-ADDRESS
// attribute of STUN Binding request. For relayed candidate, the value
// will be the address allocated in the TURN server by STUN Allocate
// request. 
func (c *IceSessCand) Addr() SockAddr {
    return SockAddr{c.c.addr}
}

func (c *IceSessCand) SetAddr(s SockAddr) {
    c.c.addr = s.s
}

// Base address of this candidate. "Base" refers to the address an agent
// sends from for a particular candidate. For host candidates, the base
// is the same as the host candidate itself. For reflexive candidates,
// the base is the local IP address of the socket. For relayed
// candidates, the base address is the transport address allocated in
// the TURN server for this candidate. 
func (c *IceSessCand) BaseAddr() SockAddr {
    return SockAddr{c.c.base_addr}
}

func (c *IceSessCand) SetBaseAddr(s SockAddr) {
    c.c.base_addr = s.s
}

// Related address, which is used for informational only and is not used
// in any way by the ICE session.
func (c *IceSessCand) RelAddr() SockAddr {
    return SockAddr{c.c.rel_addr}
}

func (c *IceSessCand) SetRelAddr(s SockAddr) {
    c.c.rel_addr = s.s
}
