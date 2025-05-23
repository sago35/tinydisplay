//go:build wioterminal
// +build wioterminal

package initdisplay

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ili9341"
)

// InitDisplay initializes the display of each board.
func InitDisplay() *ili9341.Device {
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       machine.LCD_SCK_PIN,
		SDO:       machine.LCD_SDO_PIN,
		SDI:       machine.LCD_SDI_PIN,
		Frequency: 60000000,
	})

	d := ili9341.NewSPI(
		*machine.SPI3,
		machine.LCD_DC,
		machine.LCD_SS_PIN,
		machine.LCD_RESET,
	)
	d.Configure(ili9341.Config{
		Rotation: ili9341.Rotation270,
	})
	d.FillScreen(color.RGBA{255, 255, 255, 255})

	machine.LCD_BACKLIGHT.Configure(machine.PinConfig{machine.PinOutput})
	machine.LCD_BACKLIGHT.High()

	return d
}
