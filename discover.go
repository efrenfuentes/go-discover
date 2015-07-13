package main

import (
	"fmt"
	"github.com/efrenfuentes/go-discover/arp"
	"github.com/efrenfuentes/go-discover/device"
	"github.com/efrenfuentes/go-discover/utils"
)

func scanning(table *arp.ArpTable) {
	fmt.Println("Scanning Arp Table..")
	table.Read()
}

func print_devices(devices []device.Device) {
	fmt.Println("Discover result")
	for _, device := range devices {
		if device.Discover {
			fmt.Println(device.String())
			fmt.Print("Names: ")
			fmt.Println(device.Names)
		}
	}
}

func main() {
	fmt.Println("Get local ips from interfaces...")
	addr_interfaces := utils.LocalAddresses()

	devices_interfaces := []device.Device{}

	for _, a := range addr_interfaces {
		fmt.Printf("Get ips for %v network...\n", a)
		ips_on_network, _ := utils.IPS_Network(a.String())
		for _, ip := range ips_on_network {
			info := device.Device{}
			info.SetIP(ip)
			fmt.Printf("scanning %s...\n", ip)
			if !device.DeviceInSlice(info, devices_interfaces) {
				devices_interfaces = append(devices_interfaces, info)
			}
		}
	}

	arp_table := arp.ArpTable{}
	scanning(&arp_table)

	fmt.Println("Search duplicates...")
	devices := device.DeviceJoinSlice(arp_table.Devices, devices_interfaces)

	print_devices(devices)
}
