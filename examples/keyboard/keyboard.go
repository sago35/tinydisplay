package main

import (
	"image/color"

	"github.com/sago35/tinydisplay"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

type Key struct {
	buf        []uint16
	w          int16
	h          int16
	fontHeight int16
	bgcolor    uint16
}

func NewKey(w, h int, bgcolor color.RGBA) *Key {
	return &Key{
		buf:        make([]uint16, w*h),
		w:          int16(w),
		h:          int16(h),
		fontHeight: int16(tinyfont.GetGlyph(&freemono.Bold18pt7b, '0').Height),
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
	return nil
}

func (l *Key) SetText(str string, c color.RGBA) {
	for i := range l.buf {
		l.buf[i] = l.bgcolor
	}

	tinyfont.WriteLine(l, &freemono.Bold18pt7b, 3, l.fontHeight+2, str, c)
}
