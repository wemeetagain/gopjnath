package gopjnath

/*
#cgo pkg-config: libpjproject
#include <pjlib.h>
#include <pjlib-util.h>
#include <pjnath.h>

void poll(int *quit,long *delay,pj_ioqueue_t *io, pj_timer_heap_t *theap);
*/
import "C"

import (
    "unsafe"
    )

func NewContext(name string) *Context {
	c := Context{}
	c.name = C.CString(name)
	
	C.pj_caching_pool_init(&c.cp, &C.pj_pool_factory_default_policy, C.pj_size_t(100))
    c.pool = C.pj_pool_create(&c.cp.factory,c.name,C.pj_size_t(2000),C.pj_size_t(2000),nil)
    C.pj_timer_heap_create(c.pool,C.pj_size_t(2000),&c.tHeap)
    C.pj_ioqueue_create(c.pool,C.pj_size_t(32),&c.io)
    
    // set up polling
    pollArgs := C.malloc(C.size_t(4))
    pollPtr := uintptr(unsafe.Pointer(pollArgs))
    *(**C.int) (unsafe.Pointer(pollPtr)) = c.quit
    pollPtr++
    //delay := C.malloc(C.sizeof(C.long))
    //(C.long) (unsafe.Pointer(delay)) = C.long(10)
    //*(**C.long) (unsafe.Pointer(pollPtr)) = delay
    pollPtr++
    *(**C.pj_ioqueue_t) (unsafe.Pointer(pollPtr)) = c.io
    pollPtr++
    *(**C.pj_timer_heap_t) (unsafe.Pointer(pollPtr)) = c.tHeap
    
    C.pj_thread_create(c.pool,c.name,(*C.pj_thread_proc) (C.poll),pollArgs,0,0,&c.poll)
    return &c
}

func (c *Context) Destroy() {
	C.free(unsafe.Pointer(c.name))
	//c.quit = &C.int(1)
}
