package gopjnath

/*
#cgo pkg-config: libpjproject
#include <pjlib.h>
#include <pjlib-util.h>
#include <pjnath.h>
*/
import "C"

import (
    "errors"
    "unsafe"
    )

func init() {
    C.pj_init()
    C.pjlib_util_init()
    C.pjnath_init()
}

type Context struct {
	name *C.char
	cp C.pj_caching_pool
    pool *C.pj_pool_t
    tHeap *C.pj_timer_heap_t
    io *C.pj_ioqueue_t
    poll *C.pj_thread_t
    quit *C.int
}

/*
func (c *Context) poll() {
    var delay C.pj_time_val
    delay.msec = C.long(10)
    for !c.quit {
        C.pj_ioqueue_poll(c.io,&delay)
        C.pj_timer_heap_poll(c.tHeap,nil)
    }
}
*/

type IceTransportOp int

const (
    IceTransportOpStateInit        = IceTransportOp(C.PJ_ICE_STRANS_OP_INIT)
    IceTransportOpStateNegotiation = IceTransportOp(C.PJ_ICE_STRANS_OP_NEGOTIATION)
    IceTransportOpStateKeepAlive   = IceTransportOp(C.PJ_ICE_STRANS_OP_KEEP_ALIVE)
    )

type TransportState int

const (
    TransportStateNull      = TransportState(C.PJ_ICE_STRANS_STATE_NULL)
    TransportStateInit      = TransportState(C.PJ_ICE_STRANS_STATE_INIT)
    TransportStateReady     = TransportState(C.PJ_ICE_STRANS_STATE_READY)
    TransportStateSessReady = TransportState(C.PJ_ICE_STRANS_STATE_SESS_READY)
    TransportStateNego      = TransportState(C.PJ_ICE_STRANS_STATE_NEGO)
    TransportStateRunning   = TransportState(C.PJ_ICE_STRANS_STATE_RUNNING)
    TransportStateFailed    = TransportState(C.PJ_ICE_STRANS_STATE_FAILED)
    )

type IceSessRole int

const (
    IceSessRoleUnknown     = IceSessRole(C.PJ_ICE_SESS_ROLE_UNKNOWN)
    IceSessRoleControlled  = IceSessRole(C.PJ_ICE_SESS_ROLE_CONTROLLED)
    IceSessRoleControlling = IceSessRole(C.PJ_ICE_SESS_ROLE_CONTROLLING)
    )

type QosType int 

const (
    QosTypeBestEffort = QosType(C.PJ_QOS_TYPE_BEST_EFFORT)
    QosTypeBackground = QosType(C.PJ_QOS_TYPE_BACKGROUND)
    QosTypeVideo      = QosType(C.PJ_QOS_TYPE_VIDEO)
    QosTypeVoice      = QosType(C.PJ_QOS_TYPE_VOICE)
    QosTypeControl    = QosType(C.PJ_QOS_TYPE_CONTROL)
    )

func casterr(err C.pj_status_t) error {
    buf := unsafe.Pointer(C.calloc(80,1))
    s := C.pj_strerror(err,(*C.char) (buf),80)
    str := C.pj_strbuf(&s)
    defer C.free(unsafe.Pointer(str))
    return errors.New(C.GoString(str))
}

func toString(s C.pj_str_t) string {
    str := C.pj_strbuf(&s)
    ptr := uintptr(unsafe.Pointer(str))
    var str2 = unsafe.Pointer( C.calloc( C.size_t(C.pj_strlen(&s) + 1), 1 ) )
    ptr2 := uintptr(str2)
    for i := 0; i < int(C.pj_strlen(&s)); i++ {
        *(*C.char) (unsafe.Pointer(ptr2)) = *(*C.char) (unsafe.Pointer(ptr))
        ptr++
        ptr2++
    }
    defer C.free(str2)
    return C.GoString((*C.char)(str2))
}

func destroyString(s C.pj_str_t) {
    str := C.pj_strbuf(&s)
    C.free(unsafe.Pointer(str))
}
