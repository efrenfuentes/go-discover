package arp

import (
	"bufio"
	"os"
	"strings"
  "strconv"
)

const (
	f_IPAddr int = iota
	f_HWType
	f_Flags
	f_HWAddr
	f_Mask
	f_Device
)

type ArpInfo struct {
	IPAddr string
	HWAddr string
	Up bool
}

func (info *ArpInfo) String() string {
	return "IP: " + info.IPAddr +
		     " MAC: " + info.HWAddr +
				 " Up:" + strconv.FormatBool(info.Up)
}

type ArpTable struct {
	Device []ArpInfo
}

func (table *ArpTable) Add(info ArpInfo) {
	table.Device = append(table.Device, info)
}

func (table *ArpTable) Read() {
	f, err := os.Open("/proc/net/arp")

	if err != nil {
		table.Device = nil
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan() // skip the field descriptions

	for s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)
		info := ArpInfo{
      IPAddr: fields[f_IPAddr],
      HWAddr: fields[f_HWAddr],
    }
		table.Add(info)
	}
}
