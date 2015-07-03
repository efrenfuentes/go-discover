package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"os"
	"time"
)

func ping_exist(ip *net.IPAddr) bool {
	p := fastping.NewPinger()
	p.AddIPAddr(ip)
	exist := false

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		/*fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)*/
		exist = true
	}

	err := p.Run()
	if err != nil {
		fmt.Println(err)
	}

	return exist
}

func main() {
	ip, ipnet, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ra, err := net.ResolveIPAddr("ip4:icmp", ip.String())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("ip: %s up: %t\n", ip, ping_exist(ra))
	}
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
