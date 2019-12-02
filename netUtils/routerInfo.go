package netUtils

import (
	"golang.org/x/net/route"
)

var defaultGateway = [4]byte{0, 0, 0, 0}

func GetGateway() ([4]byte, error) {
	rib, _ := route.FetchRIB(0, route.RIBTypeRoute, 0)
	messages, err := route.ParseRIB(route.RIBTypeRoute, rib)

	if err != nil {
		return defaultGateway, err
	}
	var outGateway [4]byte
	for _, message := range messages {
		routeMessage := message.(*route.RouteMessage)
		addresses := routeMessage.Addrs

		var destination, gateway *route.Inet4Addr
		ok := false

		if destination, ok = addresses[0].(*route.Inet4Addr); !ok {
			continue
		}

		if gateway, ok = addresses[1].(*route.Inet4Addr); !ok {
			continue
		}

		if destination == nil || gateway == nil {
			continue
		}

		if destination.IP == defaultGateway {
			continue
		} else {
			outGateway = gateway.IP
			return destination.IP, nil
		}
	}
	return outGateway, nil
}
