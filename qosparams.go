package gopjnath

/*
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

type QosWmmPriority int

var (
    QosWmmPrioBulkEffort = QosWmmPriority(PJ_QOS_PRIO_BULK_EFFORT)
    QosWmmPrioBulk       = QosWmmPriority(PJ_QOS_PRIO_BULK)
    QosWmmPrioVideo      = QosWmmPriority(PJ_QOS_PRIO_VIDEO)
    QosWmmPrioVoice      = QosWmmPriority(PJ_QOS_PRIO_VOICE)
    )

type QosParams struct {
    p C.struct_pj_qos_params
}

func NewQosParams() *QosParams {
    q := &QosParams
    return q
}

func (q *QosParams) SetFlags(b byte) {
    q.p._flags = C.uchar(b)
}

func (q *QosParams) GetFlags() byte {
    return byte(q.p._flags)
}

func (q *QosParams) SetDscpVal(b byte) {
    q.p._dscp_val = C.uchar(b)
}

func (q *QosParams) GetDscpVal() byte {
    return byte(q.p._dscp_val)
}

func (q *QosParams) SetSoPriority(b byte) {
    q.p._so_prio = C.uchar(b)
}

func (q *QosParams) GetSoPriority() byte {
    return byte(q.p._so_prio)
}

func (q *QosParams) SetWmmPriority(b QosWmmPriority) {
    q.p._wmm_prio = C.int(b)
}

func (q *QosParams) GetWmmPriority() QosWmmPriority {
    return QosWmmPriority(q.p._wmm_prio)
}
