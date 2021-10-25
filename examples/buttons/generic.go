//go:build !baremetal
// +build !baremetal

package main

import (
	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var (
	display *initdisplay.TinyDisplay
)

func init() {
	display = initdisplay.InitDisplay()
	display.FillScreen(white)
}

func GetPressedKey() uint16 {
	return display.GetPressedKey()
}
