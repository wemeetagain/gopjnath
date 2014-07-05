package gopjnath

import (
    //"fmt"
    "testing"
    //"time"
    )

var tester *testing.T

func testIceCallback(op IceTransportOp,err error) {
    tester.Logf("IceTransportOp: %d",op)
    if err != nil {
        tester.Fatalf("Ice Callback error: %s",err)
    }
}

func TestIceTransport(t *testing.T) {
    tester = t
    context := NewContext("test")
    c := NewIceTransportConfig(context)
    c.SetStunServer("203.183.172.196")
    c.SetStunPort(3478)
    trans,err := NewIceStreamTransport("test",*c,1,nil,testIceCallback)
    if err != nil {
        t.Fatalf("NewIceStreamTransport error: %s",err)
    }
    _ = trans
    stt := trans.State()
    if stt != 1 {
        t.Fatalf("State should return 1, returned: %d",stt)
    }
    t.Logf("State: %d",stt)
    t.Logf("State Name: %s",TransportStateName(stt))
    //time.Sleep(10*time.Second)
    hassess := trans.HasSess()
    if hassess != false {
        t.Fatalf("HassSess should return false, returned: %t",hassess)
    }
    t.Logf("HasSess: %t",hassess)
    sisrun := trans.SessIsRunning()
    if sisrun != false {
        t.Fatalf("SessIsRunning should return false, returned: %t",sisrun)
    }
    t.Logf("SessIsRunning: %t",sisrun)
    siscom := trans.SessIsComplete()
    if siscom != false {
        t.Fatalf("SessIsComplete should return false, returned: %t",siscom)
    }
    t.Logf("SessIsComplete: %t",siscom)
    testufrag := "testufrag"
    testpwd := "testpwd"
    err = trans.InitIce(IceSessRoleControlling,testufrag,testpwd)
    if err != nil {
        t.Fatalf("StartIce error: %s",err)
    }
    // run after ice starts otherwise, boom
    lu,lp, ru, rp, err := trans.UfragPwd()
    if err != nil {
        t.Fatalf("UfragPwd error: %s",err)
    }
    t.Logf("UfragPwd: %s %s %s %s",lu,lp,ru,rp)
    numCands := trans.CandsCount(1)
    t.Fatalf("CandCount: %d",numCands)
}
