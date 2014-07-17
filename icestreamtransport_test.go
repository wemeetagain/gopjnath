package gopjnath

import (
    //"fmt"
    //"net"
    "testing"
    "time"
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
    c := context.NewIceTransportConfig()
    dns,err := context.NewDnsResolver()
    if err != nil {
        t.Fatalf("NewDnsResolver error: %s",err)
    }
    err = dns.SetNs("8.8.8.8")
    if err != nil {
        t.Fatalf("SetNs error: %s",err)
    }
    c.SetDnsResolver(dns)
    ssc := NewStunSockConfig()
    sa, _ := NewSockAddr(AfIP,"0.0.0.0",0)
    ssc.SetBoundAddr(sa)
    c.SetStunSockConfig(ssc)
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
    t.Logf("State Name: %s",TransportStateName(stt))
    
    for i:=0;i<5;i++ {
        time.Sleep(time.Second)
    }
    
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
    err = trans.InitIce(IceSessRoleControlled,testufrag,testpwd)
    if err != nil {
        t.Fatalf("StartIce error: %s",err)
    }
    for i:=0;i<5;i++ {
        time.Sleep(time.Second)
    }
    
    // run after ice starts otherwise, boom
    //lu,lp, ru, rp, err := trans.UfragPwd()
    //if err != nil {
    //    t.Fatalf("UfragPwd error: %s",err)
    //}
    //t.Logf("UfragPwd: %s %s %s %s",lu,lp,ru,rp)
    numCands := trans.CandsCount(1)
    t.Logf("CandCount: %d",numCands)
    cands,err := trans.Cands(1)
    if err != nil {
        t.Fatalf("Cands error: %s",err)
    }
    adr := cands[0].Addr()
    t.Logf("cand[0] IP: %d",(&adr).IP())
    t.Logf("cand[0] Port: %d",(&adr).Port())
    t.Logf("cand[0] LocalPref: %d",cands[0].LocalPref())
    t.Logf("cand[0] TransportId: %d",cands[0].TransportId())
    t.Logf("cand[0] ComponentId: %d",cands[0].ComponentId())
    t.Logf("cand[0] Priority: %d",cands[0].Priority())
    
}
