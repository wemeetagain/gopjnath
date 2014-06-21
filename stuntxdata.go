package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>
*/
import "C"

type StunTxData struct {
    d *C.pj_stun_tx_data
}
