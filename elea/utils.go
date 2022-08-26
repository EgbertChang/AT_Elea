package elea

import "net"

func LocalIP() string {
	addresses, _ := net.InterfaceAddrs()
	for _, v := range addresses {
		if ipNet, ok := v.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
