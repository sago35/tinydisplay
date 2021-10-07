package tinykb

import (
	"fmt"
	"image/color"

	"tinygo.org/x/tinyfont"
)

const (
	sz            = 22
	fgcolor       = uint16(0x0000)
	bgcolor       = uint16(0xFFFF)
	selectedColor = uint16(0xC618)
)

var (
	buf        [sz * sz]uint16
	fontHeight int16
)

type Key rune

var keyboardFont = &Regular9pt7b

func init() {
	fontHeight = int16(tinyfont.GetGlyph(keyboardFont, '0').Height)
}

func (k *Key) Size() (int16, int16) {
	return sz, sz
}

func (k *Key) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || y < 0 || sz < x || sz < y {
		return
	}
	buf[y*sz+x] = RGBATo565(c)
}

func (k *Key) Display() error {
	for i := range buf {
		buf[i] = bgcolor
	}

	tinyfont.WriteLine(k, keyboardFont, 4, int16(fontHeight)+4, string(*k), RGB565ToRGBA(fgcolor))
	return nil
}

func (k *Key) DisplaySelected() error {
	for i := range buf {
		buf[i] = selectedColor
	}

	tinyfont.WriteLine(k, keyboardFont, 4, int16(fontHeight)+4, string(*k), color.RGBA{0xFF, 0x00, 0x00, 0xFF})
	return nil
}

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

func (k Key) String() string {
	switch k {
	case KeyEscape:
		return "KeyEscape"
	case KeyReturn:
		return "KeyReturn"
	case KeyTab:
		return "KeyTab"
	case KeyBackspace:
		return "KeyBackspace"
	case KeyInsert:
		return "KeyInsert"
	case KeyDelete:
		return "KeyDelete"
	case KeyRight:
		return "KeyRight"
	case KeyLeft:
		return "KeyLeft"
	case KeyDown:
		return "KeyDown"
	case KeyUp:
		return "KeyUp"
	case KeyPageUp:
		return "KeyPageUp"
	case KeyPageDown:
		return "KeyPageDown"
	case KeyHome:
		return "KeyHome"
	case KeyEnd:
		return "KeyEnd"
	case KeyShift:
		return "KeyShift"
	case KeyShiftRelease:
		return "KeyShiftRelease"
	case KeyClose:
		return "KeyClose"
	default:
		return fmt.Sprintf("%c", k)
	}
}
