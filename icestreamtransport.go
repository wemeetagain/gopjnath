package gopjnath

/*
#include <pjnath.h>
#include <pjlib-util.h>
#include <pjlib.h>

void * ice_cb(pj_ice_strans *ice_strans, pj_ice_strans_op op, pj_status_t status);
*/
import "C"

import (
    "unsafe"
    )

type IceStreamTransport struct {
    i             *C.pj_ice_strans
    cb            *C.pj_ice_strans_cb
    OnRxData      func(uint,[]byte,SockAddr)
    OnIceComplete func(IceTransportOp,error)
}

// pj_status_t pj_ice_strans_create (const char *name, const pj_ice_strans_cfg *cfg, unsigned comp_cnt, void *user_data, const pj_ice_strans_cb *cb, pj_ice_strans **p_ice_st)
func NewIceStreamTransport(name string, t IceTransportConfig, compCnt int) (*IceStreamTransport, error) {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    cnt := C.uint(compCnt)
    stream := IceStreamTransport{}
    err := C.pj_ice_strans_create(n,t.t,cnt,unsafe.Pointer(&stream),stream.cb,&stream.i)
    if err != C.PJ_SUCCESS {
        return &stream, casterr(err)
    }
    return &stream, nil
}

// pj_ice_strans_state pj_ice_strans_get_state (pj_ice_strans *ice_st)
func (i *IceStreamTransport) GetState() TransportState {
    return TransportState(C.pj_ice_strans_get_state(i.i))
}

// const char * pj_ice_strans_state_name (pj_ice_strans_state state)
func GetTransportStateName(t TransportState) string {
    return C.GoString(C.pj_ice_strans_state_name(C.pj_ice_strans_state(t)))
}

// pj_status_t pj_ice_strans_destroy (pj_ice_strans *ice_st)
func (i *IceStreamTransport) Destroy() error {
    return casterr(C.pj_ice_strans_destroy(i.i))
}

// void * pj_ice_strans_get_user_data (pj_ice_strans *ice_st)
// not implementing right now :)

// pj_status_t pj_ice_strans_get_options (pj_ice_strans *ice_st, pj_ice_sess_options *opt)
func (i *IceStreamTransport) GetOptions() (IceSessOptions,error) {
    o := IceSessOptions{}
    status := C.pj_ice_strans_get_options(i.i,&o.o)
    if status != C.PJ_SUCCESS {
        return o, casterr(status)
    }
    return o, nil
}

// pj_status_t pj_ice_strans_set_options (pj_ice_strans *ice_st, const pj_ice_sess_options *opt)
func (i *IceStreamTransport) SetOptions(o IceSessOptions) error {
    status := C.pj_ice_strans_set_options(i.i,&o.o)
    if status != C.PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}

// pj_grp_lock_t * pj_ice_strans_get_grp_lock (pj_ice_strans *ice_st)


// pj_status_t pj_ice_strans_init_ice (pj_ice_strans *ice_st, pj_ice_sess_role role, const pj_str_t *local_ufrag, const pj_str_t *local_passwd)
func (i *IceStreamTransport) InitIce(r IceSessRole, locUfrag, locPwd string) error {
    uf := C.CString(locUfrag)
    defer C.free(unsafe.Pointer(uf))
    ufrag := C.pj_str(uf)
    p := C.CString(locPwd)
    defer C.free(unsafe.Pointer(p))
    pwd := C.pj_str(p)
    status := C.pj_ice_strans_init_ice(i.i, C.pj_ice_sess_role(r), &ufrag, &pwd)
    if status != C.PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}

// pj_bool_t pj_ice_strans_has_sess (pj_ice_strans *ice_st)
func (i *IceStreamTransport) HasSess() bool {
    return int(C.pj_ice_strans_has_sess(i.i)) != 0
}

// pj_bool_t pj_ice_strans_sess_is_running (pj_ice_strans *ice_st)
func (i *IceStreamTransport) SessIsRunning() bool {
    return int(C.pj_ice_strans_sess_is_running(i.i)) != 0
}

// pj_bool_t pj_ice_strans_sess_is_complete (pj_ice_strans *ice_st)
func (i *IceStreamTransport) SessIsComplete() bool {
    return int(C.pj_ice_strans_sess_is_complete(i.i)) != 0
}

// unsigned pj_ice_strans_get_running_comp_cnt (pj_ice_strans *ice_st)
func (i *IceStreamTransport) GetRunningCompCount() uint {
    return uint(C.pj_ice_strans_get_running_comp_cnt(i.i))
}

// pj_status_t pj_ice_strans_get_ufrag_pwd (pj_ice_strans *ice_st, pj_str_t *loc_ufrag, pj_str_t *loc_pwd, pj_str_t *rem_ufrag, pj_str_t *rem_pwd)
func (i *IceStreamTransport) GetUfragPwd() (string, string, string, string, error) {
    var locUfrag, locPwd, remUfrag, remPwd *C.pj_str_t
    status := C.pj_ice_strans_get_ufrag_pwd(i.i,locUfrag, locPwd, remUfrag, remPwd)
    lu := C.pj_strbuf(locUfrag)
    lp := C.pj_strbuf(locPwd)
    ru := C.pj_strbuf(remUfrag)
    rp := C.pj_strbuf(remPwd)
    defer C.free(unsafe.Pointer(lu))
    defer C.free(unsafe.Pointer(lp))
    defer C.free(unsafe.Pointer(ru))
    defer C.free(unsafe.Pointer(rp))
    if status != C.PJ_SUCCESS {
        return C.GoString(lu), C.GoString(lp), C.GoString(ru), C.GoString(rp), casterr(status)
    }
    return C.GoString(lu), C.GoString(lp), C.GoString(ru), C.GoString(rp), nil
}

// unsigned pj_ice_strans_get_cands_count (pj_ice_strans *ice_st, unsigned comp_id)
func (i *IceStreamTransport) GetCandsCount(compId uint) uint {
    return uint(C.pj_ice_strans_get_cands_count(i.i,C.uint(compId)))
}

// helper function to turn pj_ice_sess_cand[] -> []IceSessCand
func cArrayToCandSlice(c unsafe.Pointer, length uint) []IceSessCand {
    cands := make([]IceSessCand,length)
    var arrayptr = uintptr(c)
    for i:=0; i < int(length); i++ {
        cands = append(cands,IceSessCand{(*C.pj_ice_sess_cand) (unsafe.Pointer(arrayptr))})
        arrayptr++
    }
    return cands
}

// pj_status_t pj_ice_strans_enum_cands (pj_ice_strans *ice_st, unsigned comp_id, unsigned *count, pj_ice_sess_cand cand[])
func (i *IceStreamTransport) GetCands(compId uint) ([]IceSessCand, error) {
    var c *C.pj_ice_sess_cand
    max := C.uint(100)
    status := C.pj_ice_strans_enum_cands(i.i,C.uint(compId),&max,c)
    if status != C.PJ_SUCCESS {
        return cArrayToCandSlice(unsafe.Pointer(c),uint(max)), casterr(status)
    }
    return cArrayToCandSlice(unsafe.Pointer(c),uint(max)), nil
}

// pj_status_t pj_ice_strans_get_def_cand (pj_ice_strans *ice_st, unsigned comp_id, pj_ice_sess_cand *cand)
func (i *IceStreamTransport) GetCand(compId uint) (IceSessCand, error) {
    id := C.uint(compId)
    cand := IceSessCand{}
    status := C.pj_ice_strans_get_def_cand(i.i,id,cand.c)
    if status != C.PJ_SUCCESS {
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
    status := C.pj_ice_strans_change_role(i.i,C.pj_ice_sess_role(C.int(r)))
    if status != C.PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}

// helper function to turn []IceSessCand -> pj_ice_sess_cand[]
func candSliceToCArray(cands []IceSessCand) unsafe.Pointer {
    var array = unsafe.Pointer(C.calloc(C.size_t(len(cands)), 1))
    var arrayptr = uintptr(array)
    for i:=0; i < len(cands); i++ {
        *(*C.pj_ice_sess_cand) (unsafe.Pointer(arrayptr)) = C.pj_ice_sess_cand(*cands[i].c)
        arrayptr++
    }
    return array
}

// pj_status_t pj_ice_strans_start_ice (pj_ice_strans *ice_st, const pj_str_t *rem_ufrag, const pj_str_t *rem_passwd, unsigned rcand_cnt, const pj_ice_sess_cand rcand[])
func (i *IceStreamTransport) StartIce(remUfrag,remPwd string, count uint, cands []IceSessCand) error {
    // must convert cands to C array
    candArray := candSliceToCArray(cands)
    defer C.free(candArray)
    rUfrag := C.CString(remUfrag)
    rPwd := C.CString(remPwd)
    defer C.free(unsafe.Pointer(rUfrag))
    defer C.free(unsafe.Pointer(rPwd))
    ru := C.pj_str(rUfrag)
    rp := C.pj_str(rPwd)
    status := C.pj_ice_strans_start_ice(i.i,&ru,&rp,C.uint(count),(*C.pj_ice_sess_cand) (candArray))
    if status != C.PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}

// const pj_ice_sess_check * pj_ice_strans_get_valid_pair (const pj_ice_strans *ice_st, unsigned comp_id)
func (i *IceStreamTransport) GetValidPair(compId uint) IceSessCheck {
    return IceSessCheck{C.pj_ice_strans_get_valid_pair(i.i,C.uint(compId))}
}


// pj_status_t pj_ice_strans_stop_ice (pj_ice_strans *ice_st)
func (i *IceStreamTransport) StopIce() error {
    status := C.pj_ice_strans_stop_ice(i.i)
    if status != C.PJ_SUCCESS {
        return casterr(status)
    }
    return nil
}    

// pj_status_t pj_ice_strans_sendto (pj_ice_strans *ice_st, unsigned comp_id, const void *data, pj_size_t data_len, const pj_sockaddr_t *dst_addr, int dst_addr_len)
func (i *IceStreamTransport) Send(compId uint, data []byte, s SockAddr) {
    
}

//export go_ice_callback
func go_ice_callback(i *C.pj_ice_strans,o C.pj_ice_strans_op,s C.pj_status_t) {
    ((*IceStreamTransport) (C.pj_ice_strans_get_user_data (i))).OnIceComplete(IceTransportOp(o),casterr(s))
}

//export go_data_callback
func go_data_callback(i *C.pj_ice_strans, comp_id C.unsigned, pkt unsafe.Pointer, size C.pj_size_t, src_addr *C.pj_sockaddr_t, src_addr_len C.unsigned) {
    sz := int(size)
    data := make([]byte,sz)
    data_ptr := uintptr(pkt)
    for index := 0; index < sz; index++ {
        data[index] = *((*byte) (unsafe.Pointer(data_ptr)))
        data_ptr++
    }
    //TODO make ral SockAddr
    s := SockAddr{}
    ((*IceStreamTransport) (C.pj_ice_strans_get_user_data (i))).OnRxData(uint(comp_id),data,s)
}
