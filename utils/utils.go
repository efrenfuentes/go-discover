package utils

import (
	"fmt"
	"log"
	"net"
	"os"
)

func IPS_Network(network string) {
	ip, ipnet, err := net.ParseCIDR(network)
	if err != nil {
		log.Fatal(err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc_ip(ip) {
		fmt.Printf("ip: %s\n", ip)
	}
}

func inc_ip(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
