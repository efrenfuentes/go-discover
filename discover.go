package main

import (
	"fmt"
	"github.com/efrenfuentes/go-discover/arp"
	"github.com/efrenfuentes/go-discover/utils"
)

func scanning(table *arp.ArpTable) {
	fmt.Println("Scanning Arp Table..")
	table.Read()
}

func print_table(table arp.ArpTable) {
	fmt.Println("Checking device response for ping..")
	for _, device := range table.Devices {
		fmt.Println(device.String())
	}
}

func main() {
	table := arp.ArpTable{}

	scanning(&table)
	print_table(table)

	utils.IPS_Network("192.168.1.0/24")
}
