package main

import (
	"image/color"

	"github.com/sago35/tinydisplay"
	"tinygo.org/x/tinyfont"
)

type Key struct {
	buf        []uint16
	w          int16
	h          int16
	fontHeight int16
	text       string
	bgcolor    uint16
	fgcolor    uint16
}

var keyboardFont = &Regular9pt7b

func NewKey(w, h int, bgcolor color.RGBA) *Key {
	return &Key{
		buf:        make([]uint16, w*h),
		w:          int16(w),
		h:          int16(h),
		fontHeight: int16(tinyfont.GetGlyph(keyboardFont, '0').Height),
		bgcolor:    tinydisplay.RGBATo565(bgcolor),
	}
}

func (l *Key) Size() (int16, int16) {
	return l.w, l.h
}

func (l *Key) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || y < 0 || l.w < x || l.h < y {
		return
	}
	l.buf[y*l.w+x] = tinydisplay.RGBATo565(c)
}

func (l *Key) Display() error {
	if l.text != "" {
		for i := range l.buf {
			l.buf[i] = l.bgcolor
		}

		tinyfont.WriteLine(l, keyboardFont, 4, l.fontHeight+4, l.text, tinydisplay.RGB565ToRGBA(l.fgcolor))
	}
	return nil
}

func (l *Key) SetText(str string, c color.RGBA) {
	l.text = str
	l.fgcolor = tinydisplay.RGBATo565(c)
	l.Display()
}

func (l *Key) SetBuf(buf []uint16) {
	l.buf = buf
	l.Display()
}
