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

func (o *IceSessOptions) GetAggressive() bool {
    return int(o.o.aggressive) != 0
}

func (o *IceSessOptions) SetNominatedCheckDelay(i uint) {
    o.o.nominated_check_delay = C.uint(i)
}

func (o *IceSessOptions) GetNominatedCheckDelay() uint {
    return uint(o.o.nominated_check_delay)
}

func (o *IceSessOptions) SetControlledAgentWantNomTimeout(i int) {
    o.o.controlled_agent_want_nom_timeout = C.int(i)
}

func (o *IceSessOptions) GetControlledAgentWantNomTimeout() uint {
    return uint(o.o.controlled_agent_want_nom_timeout)
}
