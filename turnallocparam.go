package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    )

type TurnAllocParam struct {
    t *C.pj_turn_alloc_param
}

func NewTurnAllocParam() *TurnAllocParam {
    t := &TurnAllocParam{}
    C.pj_turn_alloc_param_default(t.t)
    return t
}

func (t *TurnAllocParam) SetBandwidth(i int) {
    t.t.bandwidth = C.int(i)
}

func (t *TurnAllocParam) GetBandwidth() int {
    return int(t.t.bandwidth)
}

func (t *TurnAllocParam) SetLifetime(i int) {
    t.t.lifetime = C.int(i)
}

func (t *TurnAllocParam) GetLifetime() int {
    return int(t.t.lifetime)
}

func (t *TurnAllocParam) SeKaInterval(i int) {
    t.t.ka_interval = C.int(i)
}

func (t *TurnAllocParam) GetKaInterval() int {
    return int(t.t.ka_interval)
}
