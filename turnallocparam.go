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

type TurnAllocParam struct {
    t unsafe.Pointer
}

func NewTurnAllocParam() TurnAllocParam {
    t := &TurnAllocParam{}
    C.pj_turn_alloc_param(&t.t)
    return t
}


