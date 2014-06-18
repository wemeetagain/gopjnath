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
    o C.struct_pj_ice_sess_options
}

func (o *IceSessOptions) SetAggressive(b bool) {
    o.o.aggressive = C.int(b)
}

func (o *IceSessOptions) GetAggressive() bool {
    return bool(o.o.aggressive)
}

func (o *IceSessOptions) SetNominatedCheckDelay(i int) {
    o.o.nominated_check_delay = C.int(i)
}

func (o *IceSessOptions) GetNominatedCheckDelay() int {
    return int(o.o.nominated_check_delay)
}

func (o *IceSessOptions) SetControlledAgentWantNomTimeout(i uint) {
    o.o.controlled_agent_want_nom_timeout = C.uint(i)
}

func (o *IceSessOptions) GetControlledAgentWantNomTimeout() uint {
    return uint(o.o.controlled_agent_want_nom_timeout)
}
