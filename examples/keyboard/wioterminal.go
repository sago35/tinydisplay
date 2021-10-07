//go:build wioterminal
// +build wioterminal

package main

import (
	"machine"
	"time"

	"github.com/sago35/tinydisplay/tinykb"
	"tinygo.org/x/drivers/ili9341"
)

var (
	display *ili9341.Device
)

func init() {
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       machine.LCD_SCK_PIN,
		SDO:       machine.LCD_SDO_PIN,
		SDI:       machine.LCD_SDI_PIN,
		Frequency: 48e6,
	})

	// configure backlight
	backlight := machine.LCD_BACKLIGHT
	backlight.Configure(machine.PinConfig{machine.PinOutput})

	display = ili9341.NewSPI(
		machine.SPI3,
		machine.LCD_DC,
		machine.LCD_SS_PIN,
		machine.LCD_RESET,
	)

	// configure display
	display.Configure(ili9341.Config{})

	backlight.High()

	display.SetRotation(ili9341.Rotation270)

	display.FillScreen(white)
	time.Sleep(100 * time.Millisecond)

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
		return tinykb.KeyBackspace
	} else if !machine.WIO_KEY_B.Get() {
		return ' '
	} else if !machine.WIO_KEY_C.Get() {
		return tinykb.KeyReturn
	} else if !machine.WIO_5S_UP.Get() {
		return tinykb.KeyUp
	} else if !machine.WIO_5S_LEFT.Get() {
		return tinykb.KeyLeft
	} else if !machine.WIO_5S_RIGHT.Get() {
		return tinykb.KeyRight
	} else if !machine.WIO_5S_DOWN.Get() {
		return tinykb.KeyDown
	} else if !machine.WIO_5S_PRESS.Get() {
		return tinykb.KeyReturn
	}
	return 0xFFFF
}
