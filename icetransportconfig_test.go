package gopjnath

import (
    //"fmt"
    "testing"
    )
  
func TestIceTransportConfig(t *testing.T) {
	context := NewContext("test_IceTransportConfig")
    c := NewIceTransportConfig(context)
        dns,err := context.NewDnsResolver()
    if err != nil {
        t.Fatalf("NewDnsResolver error: %s",err)
    }
    err = dns.SetNs("8.8.8.8")
    if err != nil {
        t.Fatalf("SetNs error: %s",err)
    }
    c.SetDnsResolver(dns)
    c.SetStunMaxHostCands(uint(10))
    // stun max host cands (uint)
    if a := c.StunMaxHostCands(); a != uint(10) {
        t.Errorf("error at StunMaxHostCands: %d" , a)
    }
    // stun loop addr (bool)
    if a := c.StunLoopAddr(); a {
        t.Errorf("error at StunLoopAddr: %t" , a)
    }
    c.SetStunLoopAddr(true)
    if a := c.StunLoopAddr(); !a {
        t.Errorf("error at StunLoopAddr: %t" , a)
    }
    // stun server (string)
    c.SetStunServer("test")
    if a := c.StunServer(); a != "test" {
        t.Errorf("error at StunServer: %q" , a)
    }
    // stun port (uint16)
    c.SetStunPort(uint16(9988))
    if a := c.StunPort(); a != uint16(9988) {
        t.Errorf("error at StunPort: %d" , a)
    }
    // turn conn type (TurnTransportType)
    c.SetTurnConnType(TurnTransportUdp)
    if a := c.TurnConnType(); a != TurnTransportUdp {
        t.Errorf("error at TurnConnType: %d" , a)
    }
}
