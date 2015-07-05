package device

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"strconv"
	"time"
)

type Device struct {
	IPAddr net.IPAddr
	HWAddr string
}

func (device *Device) String() string {
	isUp, _ := device.IsUp()
	return "IP: " + device.IPAddr.String() +
		" MAC: " + device.HWAddr +
		" Up:" + strconv.FormatBool(isUp)
}

func (device *Device) SetIP(ip string) error {
	ipAddr, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		return err
	}

	device.IPAddr = *ipAddr

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
