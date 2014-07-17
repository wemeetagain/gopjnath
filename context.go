package gopjnath

/*
#cgo pkg-config: libpjproject
#include <pjlib.h>
#include <pjlib-util.h>
#include <pjnath.h>

typedef struct poll_struct {
  int *quit;
  int *delay;
  pj_ioqueue_t *io;
  pj_timer_heap_t *tHeap;
} poll_struct;

void poll(poll_struct *args);
*/
import "C"

import (
    "unsafe"
    )

type Context struct {
	name *C.char
	cp C.pj_caching_pool
    pool *C.pj_pool_t
    tHeap *C.pj_timer_heap_t
    io *C.pj_ioqueue_t
    poll *C.pj_thread_t
    quit *C.int
}

func NewContext(name string) *Context {
	c := Context{}
	c.name = C.CString(name)
    q := C.int(0)
    c.quit = (*C.int) (C.malloc(C.size_t(unsafe.Sizeof(q))))
    *c.quit = q
    
	_ = C.pj_caching_pool_init(&c.cp, &C.pj_pool_factory_default_policy, C.pj_size_t(100))
    c.pool = C.pj_pool_create(&c.cp.factory,c.name,C.pj_size_t(1000),C.pj_size_t(1000),nil)
    C.pj_timer_heap_create(c.pool,C.pj_size_t(100),&c.tHeap)
    C.pj_ioqueue_create(c.pool,C.pj_size_t(16),&c.io)
    
    // set up polling
    var pollArgs C.poll_struct
    pollArgs.quit = c.quit
    delay := C.int(50)
    pollArgs.delay = &delay
    pollArgs.io = c.io
    pollArgs.tHeap = c.tHeap

    C.pj_thread_create(c.pool,c.name,(*C.pj_thread_proc) (C.poll),unsafe.Pointer(&pollArgs),0,0,&c.poll)
    
    return &c
}

func (c *Context) Destroy() {
	C.free(unsafe.Pointer(c.name))
	*c.quit = C.int(1)
}
