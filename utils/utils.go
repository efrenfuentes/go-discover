package utils

import (
	"fmt"
	"log"
	"net"
)

func LocalAddresses() []net.Addr {
	ip_table := []net.Addr{}
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return nil
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip_table = append(ip_table, a)
				}
			}
		}
	}
	return ip_table
}

func IPS_Network(network string) ([]string, error) {
	ip_table := []string{}
	ip, ipnet, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc_ip(ip) {
		ip_table = append(ip_table, ip.String())
	}
	return ip_table, nil
}

func inc_ip(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		(ip)[j]++
		if (ip)[j] > 0 {
			break
		}
	}
}
