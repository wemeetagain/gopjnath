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

type TurnSockConfig struct {
    MaxPacketSize uint
    BoundAddr string
    PortRange uint16
    QosType QosType
    QosParams QosParams
    QosIgnoreErr bool
    RcvbufSize uint
    SndbufSize uint
}
