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

type IceSessCheckState int

const (
    IceSessCheckStateFrozen     IceSessCheckState(C.PJ_SESS_CHECK_STATE_FROZEN)
    IceSessCheckStateWaiting    IceSessCheckState(C.PJ_SESS_CHECK_STATE_WAITING)
    IceSessCheckStateInProgress IceSessCheckState(C.PJ_SESS_CHECK_STATE_IN_PROGRESS)
    IceSessCheckStateSucceeded  IceSessCheckState(C.PJ_SESS_CHECK_STATE_SUCCEEDED)
    IceSessCheckStateFailed     IceSessCheckState(C.PJ_SESS_CHECK_STATE_FAILED)

type IceSessCheck struct {
    c C.struct_pj_ice_sess_check
}

func (c *IceSessCheck) GetLCand() *IceSessCand {
    return &IceSessCand{c.c._lcand}
}

func (c *IceSessCheck) GetRCand() *IceSessCand {
    return &IceSessCand{c.c._rcand}
}

func (c *IceSessCheck) GetPriority() time.Time {
    return time.Unix(int(c.c._prio),0)
}

func (c *IceSessCheck) GetState() IceSessCheckState {
    return IceSessCheckState(c.c._state)
}

func (c *IceSessCheck) GetTxData() *StunTxData {
    return &StunTxData{c.c._tdata}
}

func (c *IceSessCheck) GetNominated() bool {
    return bool(c.c._nominated)
}

func (c *IceSessCheck) GetErrCode() error {
    return casterr(c.c._err_code)
}
