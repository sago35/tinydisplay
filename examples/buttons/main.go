package main

import (
	"image/color"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
)

func main() {
	needsRedraw := true
	x := int16(160)
	y := int16(120)
	sz := int16(10)
	mv := int16(5)
	prevKey := uint16(0xFFFF)
	guard := time.Now()
	tinyfont.WriteLine(display, &freesans.Regular9pt7b, 10, 20, "Up / Down / Right / Left to move", black)
	for {
		key := GetPressedKey()

		if prevKey != key || time.Now().Sub(guard) > 500*time.Millisecond {
			if prevKey == key {
				guard = time.Now().Add(-1 * 400 * time.Millisecond)
			} else {
				guard = time.Now()
			}
			switch key {
			case KeyRight:
				display.FillRectangle(x-sz/2, y-sz/2, sz, sz, white)
				if x+mv < 320 {
					x += mv
				}
				needsRedraw = true

			case KeyLeft:
				display.FillRectangle(x-sz/2, y-sz/2, sz, sz, white)
				if x-mv > 0 {
					x -= mv
				}
				needsRedraw = true

			case KeyUp:
				display.FillRectangle(x-sz/2, y-sz/2, sz, sz, white)
				if y-mv > 30 {
					y -= mv
				}
				needsRedraw = true

			case KeyDown:
				display.FillRectangle(x-sz/2, y-sz/2, sz, sz, white)
				if y+mv < 240 {
					y += mv
				}
				needsRedraw = true
			}

			if needsRedraw {
				display.FillRectangle(x-sz/2, y-sz/2, sz, sz, black)
				needsRedraw = false
			}
		}
		prevKey = key
	}
}

// The following is a definition of a special key that goes beyond the ASCII
// range.
const (
	KeyEscape       = 0x100
	KeyReturn       = 0x101
	KeyTab          = 0x102
	KeyBackspace    = 0x103
	KeyInsert       = 0x104
	KeyDelete       = 0x105
	KeyRight        = 0x106
	KeyLeft         = 0x107
	KeyDown         = 0x108
	KeyUp           = 0x109
	KeyPageUp       = 0x10A
	KeyPageDown     = 0x10B
	KeyHome         = 0x10C
	KeyEnd          = 0x10D
	KeyShift        = 0x1FD
	KeyShiftRelease = 0x1FE
	KeyClose        = 0x1FF
)
