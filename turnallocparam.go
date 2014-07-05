package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

// Allocation parameter, which can be given when application calls
// pj_turn_session_alloc() to allocate relay address in the TURN server.
// Application should call NewTurnAllocParam to initialize this
// structure with the default values. 
type TurnAllocParam struct {
    t *C.pj_turn_alloc_param
}

func NewTurnAllocParam() *TurnAllocParam {
    t := &TurnAllocParam{}
    C.pj_turn_alloc_param_default(t.t)
    return t
}

// The requested BANDWIDTH. Default is zero to not request any specific
//  bandwidth. Note that this attribute has been deprecated after
// TURN-08 draft, hence application should only use this attribute when
// talking to TURN-07 or older version. 
func (t *TurnAllocParam) Bandwidth() int {
    return int(t.t.bandwidth)
}

func (t *TurnAllocParam) SetBandwidth(i int) {
    t.t.bandwidth = C.int(i)
}

// The requested LIFETIME. Default is zero to not request any explicit
// allocation lifetime. 
func (t *TurnAllocParam) Lifetime() int {
    return int(t.t.lifetime)
}

func (t *TurnAllocParam) SetLifetime(i int) {
    t.t.lifetime = C.int(i)
}

// If set to non-zero, the TURN session will periodically send blank
// Send Indication every PJ_TURN_KEEP_ALIVE_SEC to refresh local NAT
// bindings. Default is zero. 
func (t *TurnAllocParam) KaInterval() int {
    return int(t.t.ka_interval)
}

func (t *TurnAllocParam) SeKaInterval(i int) {
    t.t.ka_interval = C.int(i)
}
