#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
#include "_cgo_export.h"

typedef struct poll_struct
{
  int *quit;
  int *delay;
  pj_ioqueue_t *io;
  pj_timer_heap_t *tHeap;
} poll_struct;

void poll(poll_struct *args)
{
  int *quit = args->quit;
  int *delay = args->delay;
  pj_ioqueue_t *io = args->io;
  pj_timer_heap_t *tHeap = args->tHeap;
  while(*quit==0)
  {
    pj_time_val d = {0, 0};
    pj_time_val d2 = {0, 0};
    d.msec = *delay;
    int e = pj_timer_heap_poll(tHeap,&d2);
    int c = pj_ioqueue_poll(io,&d);
  }
}
