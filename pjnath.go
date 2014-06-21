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

type cachingPool unsafe.Pointer

// TODO make init ... better
var (
    cp *C.pj_caching_pool
    pool *C.pj_pool_t
    tHeap *C.pj_timer_heap_t
    io *C.pj_ioqueue_t
    quit bool
    )

func init() {
    C.pj_init()
    C.pjlib_util_init()
    C.pjnath_init()
    C.pj_caching_pool_init(cp, &C.pj_pool_factory_default_policy, 0)
    str := C.CString("main")
    defer C.free(unsafe.Pointer(str))
    pool = C.pj_pool_create(&cp.factory,str,512,512,nil)
    C.pj_timer_heap_create(pool,C.pj_size_t(1000),&tHeap)
    C.pj_ioqueue_create(pool,C.pj_size_t(16),&io)
}

func poll() {
    var delay C.pj_time_val
    delay.msec = C.long(10)
    for !quit {
        C.pj_ioqueue_poll(io,&delay)
        C.pj_timer_heap_poll(tHeap,nil)
    }
}

type TransportOp int

const (
    TransportOpStateInit        = TransportOp(C.PJ_ICE_STRANS_OP_INIT)
    TransportOpStateNegotiation = TransportOp(C.PJ_ICE_STRANS_OP_NEGOTIATION)
    TransportOpStateKeepAlive   = TransportOp(C.PJ_ICE_STRANS_OP_KEEP_ALIVE)
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
    var buf *C.char
    s := C.pj_strerror(err,buf,80)
    str := C.pj_strbuf(&s)
    defer C.free(unsafe.Pointer(str))
    return errors.New(C.GoString(str))
}
