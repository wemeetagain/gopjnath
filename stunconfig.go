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

type StunConfig struct {
}
