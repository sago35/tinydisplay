package tinykb

import (
	"image/color"

	"github.com/sago35/tinydisplay"
	"tinygo.org/x/tinyfont"
)

var (
	buf        [22 * 22]uint16
	fontHeight int16
)

const (
	fgcolor = uint16(0x0000)
	bgcolor = uint16(0xFFFF)
)

type Key struct {
	Code rune
}

func init() {
	fontHeight = int16(tinyfont.GetGlyph(keyboardFont, '0').Height)
}

func (l *Key) Size() (int16, int16) {
	return 22, 22
}

func (l *Key) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || y < 0 || 22 < x || 22 < y {
		return
	}
	buf[y*22+x] = tinydisplay.RGBATo565(c)
}

func (l *Key) Display() error {
	for i := range buf {
		buf[i] = bgcolor
	}

	tinyfont.WriteLine(l, keyboardFont, 4, int16(fontHeight)+4, string(l.Code), tinydisplay.RGB565ToRGBA(fgcolor))
	return nil
}

func (l *Key) DisplaySelected() error {
	for i := range buf {
		buf[i] = bgcolor
	}

	tinyfont.WriteLine(l, keyboardFont, 4, int16(fontHeight)+4, string(l.Code), color.RGBA{0xFF, 0x00, 0x00, 0xFF})
	return nil
}
