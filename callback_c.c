#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
#include "_cgo_export.h"

void * ice_cb(pj_ice_strans *ice_strans, pj_ice_strans_op op, pj_status_t status)
{
  go_ice_callback(ice_strans,op,status);
}
