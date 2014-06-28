#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
#include "_cgo_export.h"

void poll(int *quit,long *delay,pj_ioqueue_t *io, pj_timer_heap_t *theap)
{
  while (!*quit)
  {
    const pj_time_val d = {0, *delay};
    pj_ioqueue_poll(io,&d);
    pj_timer_heap_poll(theap,NULL);
  }
}
