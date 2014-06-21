package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    "encoding/binary"
    "time"
    )

type IceSessCheckState int

const (
    IceSessCheckStateFrozen     = IceSessCheckState(C.PJ_ICE_SESS_CHECK_STATE_FROZEN)
    IceSessCheckStateWaiting    = IceSessCheckState(C.PJ_ICE_SESS_CHECK_STATE_WAITING)
    IceSessCheckStateInProgress = IceSessCheckState(C.PJ_ICE_SESS_CHECK_STATE_IN_PROGRESS)
    IceSessCheckStateSucceeded  = IceSessCheckState(C.PJ_ICE_SESS_CHECK_STATE_SUCCEEDED)
    IceSessCheckStateFailed     = IceSessCheckState(C.PJ_ICE_SESS_CHECK_STATE_FAILED)
    )

type IceSessCheck struct {
    c *C.pj_ice_sess_check
}

func (c *IceSessCheck) LCand() *IceSessCand {
    return &IceSessCand{c.c.lcand}
}

func (c *IceSessCheck) RCand() *IceSessCand {
    return &IceSessCand{c.c.rcand}
}

func (c *IceSessCheck) Priority() time.Time {
    return time.Unix(int64(binary.LittleEndian.Uint64(c.c.prio[:8])),0)
}

func (c *IceSessCheck) State() IceSessCheckState {
    return IceSessCheckState(c.c.state)
}

func (c *IceSessCheck) TxData() *StunTxData {
    return &StunTxData{c.c.tdata}
}

func (c *IceSessCheck) Nominated() bool {
    return int(c.c.nominated) != 0
}

func (c *IceSessCheck) Error() error {
    return casterr(c.c.err_code)
}
