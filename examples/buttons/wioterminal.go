//go:build wioterminal
// +build wioterminal

package main

import (
	"machine"

	"github.com/sago35/tinydisplay/examples/initdisplay"
	"tinygo.org/x/drivers/ili9341"
)

var (
	display *ili9341.Device
)

func init() {
	display = initdisplay.InitDisplay()
	display.FillScreen(white)

	machine.WIO_KEY_A.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.WIO_KEY_B.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.WIO_KEY_C.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	machine.WIO_5S_UP.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.WIO_5S_LEFT.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.WIO_5S_RIGHT.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.WIO_5S_DOWN.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.WIO_5S_PRESS.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
}

func GetPressedKey() uint16 {
	if !machine.WIO_KEY_A.Get() {
		return KeyBackspace
	} else if !machine.WIO_KEY_B.Get() {
		return ' '
	} else if !machine.WIO_KEY_C.Get() {
		return KeyReturn
	} else if !machine.WIO_5S_UP.Get() {
		return KeyUp
	} else if !machine.WIO_5S_LEFT.Get() {
		return KeyLeft
	} else if !machine.WIO_5S_RIGHT.Get() {
		return KeyRight
	} else if !machine.WIO_5S_DOWN.Get() {
		return KeyDown
	} else if !machine.WIO_5S_PRESS.Get() {
		return KeyReturn
	}
	return 0xFFFF
}
