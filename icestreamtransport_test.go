package gopjnath

import (
    //"fmt"
    "testing"
    )

func TestIceTransport(t *testing.T) {
    context := NewContext("test")
    c := NewIceTransportConfig(context)
    sf := c.StunConfig()
    sf.SetContext(*context)
    trans,err := NewIceStreamTransport("test",*c,1)
    if err != nil {
        t.Fatalf("NewIceStreamTransport error: %s",err)
    }
    _ = trans
    t.Log(trans.GetState())
}
