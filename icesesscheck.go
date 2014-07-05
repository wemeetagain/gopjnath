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

// IceSessCheck describes an ICE connectivity check. An ICE check
// contains a candidate pair, and will involve sending STUN Binding
// Request transaction for the purposes of verifying connectivity. A
// check is sent from the local candidate to the remote candidate of a
// candidate pair.
type IceSessCheck struct {
    c *C.pj_ice_sess_check
}

// Local candidate entry of this check. 
func (c *IceSessCheck) LCand() *IceSessCand {
    return &IceSessCand{c.c.lcand}
}

// Remote candidate entry of this check. 
func (c *IceSessCheck) RCand() *IceSessCand {
    return &IceSessCand{c.c.rcand}
}

// Check priority.
func (c *IceSessCheck) Priority() time.Time {
    return time.Unix(int64(binary.LittleEndian.Uint64(c.c.prio[:8])),0)
}

// Connectivity check state.
func (c *IceSessCheck) State() IceSessCheckState {
    return IceSessCheckState(c.c.state)
}

// STUN transmit data containing STUN Binding request that was sent as
// part of this check. The value will only be set when this check has a
// pending transaction, and is used to cancel the transaction when other
// check has succeeded. 
func (c *IceSessCheck) TxData() *StunTxData {
    return &StunTxData{c.c.tdata}
}

// Flag to indicate whether this check is nominated. A nominated check
// contains USE-CANDIDATE attribute in its STUN Binding request. 
func (c *IceSessCheck) Nominated() bool {
    return int(c.c.nominated) != 0
}

// When the check failed, this will contain the failure status of the
// STUN transaction. 
func (c *IceSessCheck) Error() error {
    return casterr(c.c.err_code)
}
