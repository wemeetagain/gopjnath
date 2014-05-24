package gopjnath

/*
#cgo pkg-config: libpjnath
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/

/*
#define CHECK(expr)     status=expr; \
                        if (status!=PJ_SUCCESS) { \
                            my_perror(#expr, status); \
                            return status; \
                        }
*/

import "C"

import (
    "sync"
    "syscall"
    "time"
    "unsafe"
    )

type cachingPool unsafe.Pointer

var (
    cp C.struct_pj_caching_pool
    pool C.struct_pj_pool_t
    stunConfig *stunConfig
    )

func init() {
    C.CHECK( C.pj_init() )
    C.CHECK( C.pjlib_init() )
    C.CHECK( C.pjnath_init() )
    C.pj_caching_pool_init(cp, &C.pj_pool_factory_default_policy, 0)
    pool = C.pj_pool_create(&cp._factory,"",512,512,nil)
    
}

type TransportOp int

const (
    TransportStateInit        TransportOp(C.PJ_ICE_STRANS_OP_INIT)
    TransportStateNegotiation TransportOp(C.PJ_ICE_STRANS_OP_NEGOTIATION)
    TransportStateKeepAlive   TransportOp(C.PJ_ICE_STRANS_OP_KEEP_ALIVE)
    )

type TransportState int

const (
    TransportStateNull      TransportState(C.PJ_ICE_STRANS_STATE_NULL)
    TransportStateInit      TransportState(C.PJ_ICE_STRANS_STATE_INIT)
    TransportStateReady     TransportState(C.PJ_ICE_STRANS_STATE_READY)
    TransportStateSessReady TransportState(C.PJ_ICE_STRANS_STATE_SESS_READY)
    TransportStateNego      TransportState(C.PJ_ICE_STRANS_STATE_NEGO)
    TransportStateRunning   TransportState(C.PJ_ICE_STRANS_STATE_RUNNING)
    TransportStateFailed    TransportState(C.PJ_ICE_STRANS_STATE_FAILED)
    )

type TransportType int

const (
    TpNone = TransportType(C.TP_NONE)
    TpStun = TransportType(C.TP_STUN)
    TpTurn = TransportType(C.TP_TURN)
    )

type IceSessRole int

const (
    IceSessRoleUnknown     IceSessRole(C.PJ_ICE_SESS_ROLE_UNKNOWN)
    IceSessRoleControlled  IceSessRole(C.PJ_ICE_SESS_ROLE_CONTROLLED)
    IceSessRoleControlling IceSessRole(C.PJ_ICE_SESS_ROLE_CONTROLLING)
    
type QosType int 

const (
    QosTypeBestEffort QosType(C.PJ_QOS_TYPE_BEST_EFFORT)
    QosTypeBackground QosType(C.PJ_QOS_TYPE_BACKGROUND)
    QosTypeVideo      QosType(C.PJ_QOS_TYPE_VIDEO)
    QosTypeVoice      QosType(C.PJ_QOS_TYPE_VOICE)
    QosTypeControl    QosType(C.PJ_QOS_TYPE_CONTROL)
    )

func casterr(fromcgo error) error {
    errno := fromcgo.(syscall.Errno)
    if !ok {
        return fromcgo
    }
    pjstatus := pjStatus(errno)
    return errno
}
