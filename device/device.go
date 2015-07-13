package device

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"strconv"
	"time"
)

type Device struct {
	IPAddr   net.IPAddr
	Names    []string
	HWAddr   string
	Discover bool
}

func (device *Device) String() string {
	return "IP: " + device.IPAddr.String() + " MAC: " + device.HWAddr +
		" Up:" + strconv.FormatBool(device.Discover)
}

func (device *Device) SetIP(ip string) error {
	ipAddr, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		return err
	}

	device.IPAddr = *ipAddr

	device.Discover, _ = device.IsUp()

	device.Names, _ = net.LookupAddr(ip)

	return nil
}

func (device *Device) IsUp() (bool, error) {
	isUp := false

	p := fastping.NewPinger()
	p.AddIPAddr(&device.IPAddr)

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		isUp = true
	}

	err := p.Run()
	if err != nil {
		return isUp, err
	}

	return isUp, nil
}

func DeviceInSlice(device Device, list []Device) bool {
	for _, d := range list {
		if d.IPAddr.IP.String() == device.IPAddr.IP.String() {
			return true
		}
	}
	return false
}

func DeviceJoinSlice(list_base []Device, list_new []Device) []Device {
	list_result := list_base
	for _, d := range list_new {
		if !DeviceInSlice(d, list_result) {
			list_result = append(list_result, d)
		}
	}
	return list_result
}
