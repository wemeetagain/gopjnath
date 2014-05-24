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

func (c *IceSessCheck) LCand() *IceSessCand {
    return &IceSessCand{c.c._lcand}
}

func (c *IceSessCheck) RCand() *IceSessCand {
    return &IceSessCand{c.c._rcand}
}

func (c *IceSessCheck) Priority() time.Time {
    return time.Unix(int(c.c._prio),0)
}

func (c *IceSessCheck) State() IceSessCheckState {
    return IceSessCheckState(c.c._state)
}

func (c *IceSessCheck) TxData() *StunTxData {
    return &StunTxData{c.c._tdata}
}

func (c *IceSessCheck) Nominated() bool {
    return bool(c.c._nominated)
}

func (c *IceSessCheck) Error() error {
    return casterr(c.c._err_code)
}
