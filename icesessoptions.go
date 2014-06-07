package gopjnath

/*
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

type IceSessOptions struct {
    o C.struct_pj_ice_sess_options
}

func (o *IceSessOptions) SetAggressive(b bool) {
    o.o._aggressive = C.int(b)
}

func (o *IceSessOptions) GetAggressive() bool {
    return bool(o.o._aggressive)
}

func (o *IceSessOptions) SetNominatedCheckDelay(i int) {
    o.o._nominated_check_delay = C.int(i)
}

func (o *IceSessOptions) GetNominatedCheckDelay() int {
    return int(o.o._nominated_check_delay)
}

func (o *IceSessOptions) SetControlledAgentWantNomTimeout(i uint) {
    o.o._controlled_agent_want_nom_timeout = C.uint(i)
}

func (o *IceSessOptions) GetControlledAgentWantNomTimeout() uint {
    return uint(o.o._controlled_agent_want_nom_timeout)
}
