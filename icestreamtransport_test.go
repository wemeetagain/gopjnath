package gopjnath

import (
    //"fmt"
    "testing"
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
    trans,err := NewIceStreamTransport("test",*c,1,nil,testIceCallback)
    if err != nil {
        t.Fatalf("NewIceStreamTransport error: %s",err)
    }
    _ = trans
    stt := trans.GetState()
    t.Logf("GetState: %d",stt)
    t.Logf("State Name: %s",GetTransportStateName(stt))
    
}
