package gopjnath

import (
    "testing"
    )

func TestDnsResolver(t *testing.T) {
    context := NewContext("test_IceTransportConfig")
    d, err := context.NewDnsResolver()
    if err != nil {
        t.Fatalf("NewDnsResolver error: %s",err)
    }
    err = d.SetNs("208.67.222.222")
    if err != nil {
        t.Fatalf("SetNs error: %s",err)
    }
}
