package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    "bytes"
    "encoding/binary"
    "net"
    )

type SockAddr struct {
    s C.union_pj_sockaddr
}

func (s *SockAddr) IP() net.IP {
    switch {
    case bytes.Equal([]byte{0,1},s.s[:2]): //IPv4
        return net.IPv4(s.s[4],s.s[5],s.s[6],s.s[7])
    case bytes.Equal([]byte{0,2},s.s[:2]): //IPv6
        p := make(net.IP, net.IPv6len)
        copy(p, s.s[8:24])
        return p
    }
    return nil
}

func (s *SockAddr) Port() uint16 {
    return binary.LittleEndian.Uint16(s.s[2:4])
}
