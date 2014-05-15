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
    

func casterr(fromcgo error) error {
    errno := fromcgo.(syscall.Errno)
    if !ok {
        return fromcgo
    }
    pjstatus := pjStatus(errno)
    return errno
}
