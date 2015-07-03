package arp

import (
	"bufio"
	"os"
	"strings"
	"device"
)

const (
	f_IPAddr int = iota
	f_HWType
	f_Flags
	f_HWAddr
	f_Mask
	f_Device
)

type ArpTable struct {
	Devices []device.Device
}

func (table *ArpTable) Add(info device.Device) {
	table.Devices = append(table.Devices, info)
}

func (table *ArpTable) Read() {
	f, err := os.Open("/proc/net/arp")

	if err != nil {
		table.Devices = nil
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan() // skip the field descriptions

	for s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)
		info := device.Device{
			IPAddr: fields[f_IPAddr],
			HWAddr: fields[f_HWAddr],
		}
		table.Add(info)
	}
}
