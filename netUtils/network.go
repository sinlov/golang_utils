package netUtils

import (
	"fmt"
	"net"
)

func NetworkLocalIPv4() (ipv4 string, err error) {
	var (
		netAdds []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP addr
		isIpNet bool
	)
	// all net interface
	if netAdds, err = net.InterfaceAddrs(); err != nil {
		return "", err
	}
	// get first not lo address
	for _, addr = range netAdds {
		// must ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// skip IPv6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}

	err = fmt.Errorf("ERR NO LOCAL IP FOUND")
	return
}

func NetworkLocalIPv6() (ipv6 string, err error) {
	var (
		netAdds []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP addr
		isIpNet bool
	)
	// all net interface
	if netAdds, err = net.InterfaceAddrs(); err != nil {
		return "", err
	}
	// get first not lo address
	for _, addr = range netAdds {
		// must ipv6, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// skip IPv4
			if ipNet.IP.To16() != nil {
				ipv6 = ipNet.IP.String()
				return
			}
		}
	}

	err = fmt.Errorf("ERR NO LOCAL IP FOUND")
	return
}

func NetworkMacAddr() (macAddr string, err error) {
	var (
		netAdds []net.Addr
		ipNet   *net.IPNet // IP addr
		isIpNet bool
	)

	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	// all net interface
	if netAdds, err = net.InterfaceAddrs(); err != nil {
		return "", err
	}
	for id, addr := range netAdds {
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			inter := interfaces[id]
			if inter.HardwareAddr == nil {
				continue
			}
			macAddr = inter.HardwareAddr.String()
			return
		}
	}
	return
}
