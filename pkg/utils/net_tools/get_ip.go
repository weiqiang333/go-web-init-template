package net_tools

import (
	"log"
	"net"
)

// GetLocalhostIps 获取本地网络接口 ip list
func GetLocalhostIps() ([]string, error) {
	var localhostIps []string
	ifaces, err := net.Interfaces()
	if err != nil {
		return localhostIps, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Println("Failed GetLocalhostIps error: ", err.Error())
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPAddr:
				ip = v.IP
			case *net.IPNet:
				ip = v.IP
			}
			if ip == nil || ip.IsUnspecified() || ip.IsLoopback() {
				continue
			}
			if ipv4 := ip.To4(); ipv4 != nil {
				localhostIps = append(localhostIps, ipv4.String())
				continue
			}
			if ipv6 := ip.To16(); ipv6 != nil {
				localhostIps = append(localhostIps, ipv6.String())
				continue
			}
		}
	}
	return localhostIps, nil
}
