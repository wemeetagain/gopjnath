package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

import (
    "unsafe"
    )

type DnsResolver struct {
    d *C.pj_dns_resolver
}

func (c *Context) NewDnsResolver() (*DnsResolver,error) {
    d := DnsResolver{}
    name := C.CString("resolver")
    status := C.pj_dns_resolver_create(&c.cp.factory, name, 0, c.tHeap, c.io, &d.d)
    return &d, casterr(status)
}

func (d *DnsResolver) SetNs(name string) error {
    char := C.CString(name)
    defer C.free(unsafe.Pointer(char))
    ns := C.pj_str(char)
    status := C.pj_dns_resolver_set_ns(d.d, 1, &ns, nil)
    return casterr(status)
}
