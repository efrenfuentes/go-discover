package main

import (
	"fmt"
	"github.com/efrenfuentes/go-discover/arp"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"time"
)

func scanning(table *arp.ArpTable) {
	fmt.Println("Scanning Arp Table..")
	table.Read()
}

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

func checking(table arp.ArpTable) {
	fmt.Println("Checking device response for ping..")
	for _, device := range table.Devices {
		ra, err := net.ResolveIPAddr("ip4:icmp", device.IPAddr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		device.Up = ping_exist(ra)
		fmt.Println(device.String())
	}
}

func main() {
	table := arp.ArpTable{}

	scanning(&table)
	checking(table)
}
