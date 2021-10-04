package tinykb

import (
	"image/color"
)

type Keyboard interface {
	Display()
	KeyEvent(key Key)
	GetKey() Key
}

func RGB565ToRGBA(c uint16) color.RGBA {
	return color.RGBA{
		R: uint8((c & 0xF800) >> 8),
		G: uint8((c & 0x07E0) >> 3),
		B: uint8((c & 0x001F) << 3),
		A: 0xFF,
	}
}