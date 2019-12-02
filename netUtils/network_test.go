package netUtils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkLocalIPv4(t *testing.T) {
	// mock NetworkLocalIPv4
	//t.Logf("mock NetworkLocalIPv4")
	notWant := "0.0.0.0"
	// do NetworkLocalIPv4
	//t.Logf("do NetworkLocalIPv4")
	pv4, err := NetworkLocalIPv4()
	if err != nil {
		t.Errorf("TestNetworkLocalIPv4 err: %v", err)
	}
	t.Logf("got IPv4: %v", pv4)
	// verify NetworkLocalIPv4
	assert.NotEqual(t, notWant, pv4)
}

func TestNetworkLocalIPv6(t *testing.T) {
	// mock NetworkLocalIPv6
	notWant := "0:0:0:0:0:0:0:0"
	//t.Logf("mock NetworkLocalIPv6")
	// do NetworkLocalIPv6
	//t.Logf("do NetworkLocalIPv6")
	iPv6, err := NetworkLocalIPv6()
	if err != nil {
		t.Errorf("NetworkLocalIPv6 err: %v", err)
	}
	t.Logf("got IPv6: %v", iPv6)
	// verify NetworkLocalIPv6
	assert.NotEqual(t, notWant, iPv6)
}

func TestNetworkMacAddr(t *testing.T) {
	// mock NetworkMacAddr
	notWant := "00:00:00:00"
	//t.Logf("mock NetworkMacAddr")
	// do NetworkMacAddr
	//t.Logf("do NetworkMacAddr")
	macAddr, err := NetworkMacAddr()
	if err != nil {
		t.Errorf("NetworkMacAddr err: %v", err)
	}
	t.Logf("got NetworkMacAddr: %v", macAddr)
	// verify NetworkMacAddr
	assert.NotEqual(t, notWant, macAddr)
}
