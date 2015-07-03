package device

import (
    "strconv"
)

type Device struct {
    IPAddr string
    HWAddr string
    Up     bool
}

func (device *Device) String() string {
    return "IP: " + device.IPAddr +
        " MAC: " + device.HWAddr +
        " Up:" + strconv.FormatBool(device.Up)
}
