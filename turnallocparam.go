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

type TurnAllocParam struct {
    t C.struct_pj_turn_alloc_param
}

func NewTurnAllocParam() *TurnAllocParam {
    t := &TurnAllocParam{}
    C.pj_turn_alloc_param_default(t.t)
    return t
}

func (t *TurnAllocParam) SetBandwidth(i int) {
    t.t._bandwidth = C.int(i)
}

func (t *TurnAllocParam) GetBandwidth() int {
    return int(t.t._bandwidth)
}

func (t *TurnAllocParam) SetLifetime(i int) {
    t.t._lifetime = C.int(i)
}

func (t *TurnAllocParam) GetLifetime() int {
    return int(t.t._lifetime)
}

func (t *TurnAllocParam) SeKaInterval(i int) {
    t.t._ka_interval = C.int(i)
}

func (t *TurnAllocParam) GetKaInterval() int {
    return int(t.t._ka_interval)
}
