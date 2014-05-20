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

type StunSockConfig struct {
    MaxPacketSize uint
    AsyncCount uint
    BoundAddr string
    PortRange uint16
    KaInterval int
    QosType QosType
    QosIgnoreErr bool
    RcvbufSize uint
    SndbufSize uint
}
