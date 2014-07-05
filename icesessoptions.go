package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

type IceSessOptions struct {
    o C.pj_ice_sess_options
}

func (o *IceSessOptions) SetAggressive(b bool) {
    if b {
        o.o.aggressive = C.pj_bool_t(C.int(1))
    } else {
        o.o.aggressive = C.pj_bool_t(C.int(0))
    }
}

// Specify whether to use aggressive nomination. 
func (o *IceSessOptions) Aggressive() bool {
    return int(o.o.aggressive) != 0
}

func (o *IceSessOptions) SetNominatedCheckDelay(i uint) {
    o.o.nominated_check_delay = C.uint(i)
}

// For a controlled agent, specify how long it wants to wait (in
// milliseconds) for the controlling agent to complete sending
// connectivity check with nominated flag set to true for all components
// after the controlled agent has found that all connectivity checks in
// its checklist have been completed and there is at least one
// successful (but not nominated) check for every component. Default
// value for this option is
// ICE_CONTROLLED_AGENT_WAIT_NOMINATION_TIMEOUT. Specify -1 to disable
// this timer.
func (o *IceSessOptions) NominatedCheckDelay() uint {
    return uint(o.o.nominated_check_delay)
}

func (o *IceSessOptions) SetControlledAgentWantNomTimeout(i int) {
    o.o.controlled_agent_want_nom_timeout = C.int(i)
}

// For controlling agent if it uses regular nomination, specify the
// delay to perform nominated check (connectivity check with
// USE-CANDIDATE attribute) after all components have a valid pair.
// Default value is PJ_ICE_NOMINATED_CHECK_DELAY.
func (o *IceSessOptions) ControlledAgentWantNomTimeout() uint {
    return uint(o.o.controlled_agent_want_nom_timeout)
}
