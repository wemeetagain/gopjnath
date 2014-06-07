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

type IceStreamTransport struct {
    i     unsafe.Pointer
    cb    unsafe.Pointer
    mutex sync.Mutex // ensure init is only called once
    init  func() // func that initializes xxx
    err   error // error returned from xxx
}

// pj_status_t pj_ice_strans_create (const char *name, const pj_ice_strans_cfg *cfg, unsigned comp_cnt, void *user_data, const pj_ice_strans_cb *cb, pj_ice_strans **p_ice_st)
func NewIceStreamTransport(name string, t IceTransportConfig, compCnt int) (*IceStreamTransport, error) {
    p := unsafe.Pointer{}
    stream := IceStreamTransport{}
    err := C.pj_ice_strans_create(name,t,compCnt,p,stream.cb,stream.i)
    if err != C.PJ_SUCCESS {
        return stream, casterr(err)
    }
    return stream, nil
}

// pj_ice_strans_state pj_ice_strans_get_state (pj_ice_strans *ice_st)
func (i *IceStreamTransport) GetState() TransportState {
    return TransportState(C.pj_ice_strans_get_state(i.i))
}

// const char * pj_ice_strans_state_name (pj_ice_strans_state state)
func GetTransportStateName(t TransportState) {
    return string(C.pj_ice_strans_state_name(C.pj_ice_strans_state(t)))
}

// pj_status_t pj_ice_strans_destroy (pj_ice_strans *ice_st)
func (i *IceStreamTransport) Destroy() error {
    return casterr(pj_ice_strans_destroy(i.i))
}

// void * pj_ice_strans_get_user_data (pj_ice_strans *ice_st)
// not implementing right now :)

// pj_status_t pj_ice_strans_get_options (pj_ice_strans *ice_st, pj_ice_sess_options *opt)
func (i *IceStreamTransport) GetOptions() (IceSessOptions,error) {
    o := IceSessionOptions{}
    status := C.pj_ice_strans_get_options(i.i,o.o)
    if status != C.PJ_SUCCESS {
        return o, casterr(status)
    }
    return o, nil
}

// pj_status_t pj_ice_strans_set_options (pj_ice_strans *ice_st, const pj_ice_sess_options *opt)
func (i *IceStreamTransport) SetOptions(o IceSessOptions) error {
    status := C.pj_ice_strans_set_options(i.i,o.o)
    if status != C.PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}

// pj_grp_lock_t * pj_ice_strans_get_grp_lock (pj_ice_strans *ice_st)


// pj_status_t pj_ice_strans_init_ice (pj_ice_strans *ice_st, pj_ice_sess_role role, const pj_str_t *local_ufrag, const pj_str_t *local_passwd)
func (i *IceStreamTransport) InitIce(r IceSessRole, locUfrag, ulocPwd string) error {
    status := C.pj_ice_strans_init_ice(i.i, C.pj_ice_sess_role(r), C.pj_str_t(locUfrag), C.pj_str_t(locPwd))
    if status != C.PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}

// pj_bool_t pj_ice_strans_has_sess (pj_ice_strans *ice_st)
func (i *IceStreamTransport) HasSess() bool {
    return bool(C.pj_ice_strans_has_sess(i.i))
}

// pj_bool_t pj_ice_strans_sess_is_running (pj_ice_strans *ice_st)
func (i *IceStreamTransport) SessIsRunning() bool {
    return bool(C.pj_ice_strans_sess_is_running(i.i))
}

// pj_bool_t pj_ice_strans_sess_is_complete (pj_ice_strans *ice_st)
func (i *IceStreamTransport) SessIsComplete() bool {
    return bool(C.pj_ice_strans_sess_is_complete(i.i))
}

// unsigned pj_ice_strans_get_running_comp_cnt (pj_ice_strans *ice_st)
func (i *IceStreamTransport) GetRunningCompCount() uint {
    return uint(C.pj_ice_strans_get_running_comp_cnt(i.i))
}

// pj_status_t pj_ice_strans_get_ufrag_pwd (pj_ice_strans *ice_st, pj_str_t *loc_ufrag, pj_str_t *loc_pwd, pj_str_t *rem_ufrag, pj_str_t *rem_pwd)
func (i *IceStreamTransport) GetUfragPwd() (string, string, string, string, error) {
    var locUfrag, locPwd, remUfrag, remPwd C.pj_str_t
    status := C.pj_ice_strans_get_ufrag_pwd(i.i,locUfrag, locPwd, remUfrag, remPwd)
    if status != PJ_SUCCESS {
        return locUfrag, locPwd, remUfrag, remPwd, casterr(status)
    }
    return locUfrag, locPwd, remUfrag, remPwd, nil
}

// unsigned pj_ice_strans_get_cands_count (pj_ice_strans *ice_st, unsigned comp_id)
func (i *IceStreamTransport) GetCandsCount() uint {
    return uint(C.pj_ice_strans_get_cands_count(i.i))
}

// pj_status_t pj_ice_strans_enum_cands (pj_ice_strans *ice_st, unsigned comp_id, unsigned *count, pj_ice_sess_cand cand[])
func (i *IceStreamTransport) GetCands() ([]IceSessionCand, error) {
    var count C.uint
    cands := make([]IceSessionCand)
    status := C.pj_ice_strans_enum_cands(i.i,count,cands)
    if status != PJ_SUCCESS {
        return cand, casterr(status)
    }
    return cands, nil
}

// pj_status_t pj_ice_strans_get_def_cand (pj_ice_strans *ice_st, unsigned comp_id, pj_ice_sess_cand *cand)
func (i *IceStreamTransport) GetCand(compId uint) (IceSessionCand, error) {
    id := C.uint(compId)
    cand := IceSessionCand{}
    status := C.pj_ice_strans_get_def_cand(i.i,id,cand)
    if status != PJ_SUCCESS {
        return cand, casterr(status)
    }
    return cand, nil
}

// pj_ice_sess_role pj_ice_strans_get_role (pj_ice_strans *ice_st)
func (i *IceStreamTransport) Role() IceSessRole {
    return IceSessRole(C.pj_ice_strans_get_role(i.i))
}

// pj_status_t pj_ice_strans_change_role (pj_ice_strans *ice_st, pj_ice_sess_role new_role)
func (i *IceStreamTransport) ChangeRole(r IceSessRole) error {
    status := C.pj_ice_strans_change_role(i.i,r)
    if status != PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}

// helper function to turn []IceSessCand -> pj_ice_sess_cand[]
func candSliceToCArray(cands []IceSessCand) {
    var array = unsafe.Pointer(C.calloc(C.size_t(len(cands)), 1))
    var arrayptr = uintptr(array)
    for i:=0; i < len(cands); i++ {
        *(*C.pj_ice_sess_cand) (unsafe.Pointer(arrayptr)) = C.pj_ice_sess_cand(cands[i])
        arrayptr++
    }
    return array
}

// pj_status_t pj_ice_strans_start_ice (pj_ice_strans *ice_st, const pj_str_t *rem_ufrag, const pj_str_t *rem_passwd, unsigned rcand_cnt, const pj_ice_sess_cand rcand[])
func (i *IceStreamTransport) StartIce(remUfrag,remPwd string, count uint, cands []IceSessCand) error {
    // must convert cands to C array
    candArray := candSliceToCArray(cands)
    defer C.free(candArray)
    status := C.pj_ice_strans_start_ice(i.i,C.pj_str_t(remUfrag),C.pj_str_t(remPwd),C.uint(count),candArray)
    if status != PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}

// const pj_ice_sess_check * pj_ice_strans_get_valid_pair (const pj_ice_strans *ice_st, unsigned comp_id)
func (i *IceStreamTransport) GetValidPair(compId uint) IceSessCheck {
    return IceSessCheck(C.pj_ice_strans_get_valid_pair(i.i,C.uint(compId)))
}


// pj_status_t pj_ice_strans_stop_ice (pj_ice_strans *ice_st)
func (i *IceStreamTransport) StopIce() error {
    status := C.pj_ice_strans_stop_ice(i.i)
    if status != PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}    

// pj_status_t pj_ice_strans_sendto (pj_ice_strans *ice_st, unsigned comp_id, const void *data, pj_size_t data_len, const pj_sockaddr_t *dst_addr, int dst_addr_len)

