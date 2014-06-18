package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    "unsafe"
    )

type IceCandType int

const (
    IceCandTypeHost    = IceCandType(C.PJ_ICE_CAND_TYPE_HOST)
    IceCandTypeSrFlx   = IceCandType(C.PJ_ICE_CAND_TYPE_SRFLX)
    IceCandTypePrFlx   = IceCandType(C.PJ_ICE_CAND_TYPE_PRFLX)
    IceCandTypeRelayed = IceCandType(C.PJ_ICE_CAND_TYPE_RELAYED)
    )

type IceSessCand struct {
    c *C.struct_pj_ice_sess_cand
}

func (c *IceSessCand) Type() IceCandType {
    return IceCandType(c.c._type)
}

func (c *IceSessCand) Status() error {
    return casterr(c.c.status)
}

func (c *IceSessCand) ComponentId() uint8 {
    return uint8(c.c.comp_id)
}

func (c *IceSessCand) TransportId() uint8 {
    return uint8(c.c.transport_id)
}

func (c *IceSessCand) LocalPref() uint16 {
    return uint16(c.c.local_pref)
}

func (c *IceSessCand) Foundation() string {
    f := C.pj_strbuf(&c.c.foundation)
    defer C.free(unsafe.Pointer(f))
    return C.GoString(f)
}

func (c *IceSessCand) Priority() uint32 {
    return uint32(c.c.prio)
}

func (c *IceSessCand) Addr() SockAddr {
    return SockAddr{c.c.addr}
}

func (c *IceSessCand) BaseAddr() SockAddr {
    return SockAddr{c.c.base_addr}
}

func (c *IceSessCand) RelAddr() SockAddr {
    return SockAddr{c.c.rel_addr}
}
