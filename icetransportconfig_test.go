package gopjnath

import (
    //"fmt"
    "testing"
    )
  
func TestIceTransportConfig(t *testing.T) {
    c := NewIceTransportConfig()
    c.SetStunMaxHostCands(uint(10))
    // stun max host cands (uint)
    if a := c.GetStunMaxHostCands(); a != uint(10) {
        t.Errorf("error at GetStunMaxHostCands: %d" , a)
    }
    // stun loop addr (bool)
    if a := c.GetStunLoopAddr(); a {
        t.Errorf("error at GetStunLoopAddr: %t" , a)
    }
    c.SetStunLoopAddr(true)
    if a := c.GetStunLoopAddr(); !a {
        t.Errorf("error at GetStunLoopAddr: %t" , a)
    }
    // stun server (string)
    c.SetStunServer("test")
    if a := c.GetStunServer(); a != "test" {
        t.Errorf("error at GetStunServer: %q" , a)
    }
    // stun port (uint16)
    c.SetStunPort(uint16(9988))
    if a := c.GetStunPort(); a != uint16(9988) {
        t.Errorf("error at GetStunPort: %d" , a)
    }
    // turn conn type (TurnTransportType)
    c.SetTurnConnType(TurnTransportUdp)
    if a := c.GetTurnConnType(); a != TurnTransportUdp {
        t.Errorf("error at GetTurnConnType: %d" , a)
    }
}
