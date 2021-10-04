package tinykb

import (
	"image/color"

	"tinygo.org/x/drivers"
)

const (
	KeyEscape    = 0x100
	KeyReturn    = 0x101
	KeyTab       = 0x102
	KeyBackspace = 0x103
	KeyInsert    = 0x104
	KeyDelete    = 0x105
	KeyRight     = 0x106
	KeyLeft      = 0x107
	KeyDown      = 0x108
	KeyUp        = 0x109
	KeyPageUp    = 0x10A
	KeyPageDown  = 0x10B
	KeyHome      = 0x10C
	KeyEnd       = 0x10D
	KeyCloseMenu = 0x1FF
)

type Keyboard14x4 struct {
	Layout     [4][14]Key
	Disp       drivers.Displayer
	fontHeight int16
	bgcolor    uint16
	fgcolor    uint16
	Selected   [2]int
}

var keyboardFont = &Regular9pt7b

var K14x4 = [4][14]Key{
	{
		Key{Code: '`'},
		Key{Code: '1'},
		Key{Code: '2'},
		Key{Code: '3'},
		Key{Code: '4'},
		Key{Code: '5'},
		Key{Code: '6'},
		Key{Code: '7'},
		Key{Code: '8'},
		Key{Code: '9'},
		Key{Code: '0'},
		Key{Code: '-'},
		Key{Code: '='},
		Key{Code: KeyBackspace},
	},
	{
		Key{Code: 'q'},
		Key{Code: 'w'},
		Key{Code: 'e'},
		Key{Code: 'r'},
		Key{Code: 't'},
		Key{Code: 'y'},
		Key{Code: 'u'},
		Key{Code: 'i'},
		Key{Code: 'o'},
		Key{Code: 'p'},
		Key{Code: '['},
		Key{Code: ']'},
		Key{Code: '\\'},
		Key{Code: KeyTab},
	},
	{
		Key{Code: 'a'},
		Key{Code: 's'},
		Key{Code: 'd'},
		Key{Code: 'f'},
		Key{Code: 'g'},
		Key{Code: 'h'},
		Key{Code: 'j'},
		Key{Code: 'k'},
		Key{Code: 'l'},
		Key{Code: ';'},
		Key{Code: '\''},
		Key{Code: KeyReturn},
		Key{Code: KeyUp},
		Key{Code: KeyCloseMenu},
	},
	{
		Key{Code: 'z'},
		Key{Code: 'x'},
		Key{Code: 'c'},
		Key{Code: 'v'},
		Key{Code: 'b'},
		Key{Code: 'n'},
		Key{Code: 'm'},
		Key{Code: ','},
		Key{Code: '.'},
		Key{Code: '/'},
		Key{Code: ' '},
		Key{Code: KeyLeft},
		Key{Code: KeyDown},
		Key{Code: KeyRight},
	},
}

func New(display drivers.Displayer) *Keyboard14x4 {
	return &Keyboard14x4{
		Disp:     display,
		Layout:   K14x4,
		Selected: [2]int{-1, -1},
	}
}

func (k *Keyboard14x4) Display() {
	for y := 0; y < len(k.Layout); y++ {
		for x := 0; x < len(k.Layout[0]); x++ {
			k.Redraw(x, y, false)
		}
	}
}

func (k *Keyboard14x4) Redraw(x, y int, selected bool) {
	sz := int16(22)
	szi16 := int16(sz)
	ybase := int16(150)

	if selected {
		if k.Selected[0] == x && k.Selected[1] == y {
			return
		}
		if k.Selected[0] >= 0 && k.Selected[1] >= 0 {
			k.Redraw(k.Selected[0], k.Selected[1], false)
		}
		k.Selected[0] = x
		k.Selected[1] = y
	}

	btn := k.Layout[y][x]
	xxx := 1 + (szi16+1)*(int16(x)+0) + szi16/2*0
	yyy := ybase + ((szi16 + 1) * int16(y))
	if selected {
		btn.DisplaySelected()
	} else {
		btn.Display()
	}
	for yy := int16(0); yy < sz; yy++ {
		for xx := int16(0); xx < sz; xx++ {
			k.Disp.SetPixel(xxx+xx, yyy+yy, RGB565ToRGBA(buf[xx+yy*sz]))
		}
	}
}

func (k *Keyboard14x4) GetKey() rune {
	key := k.Layout[k.Selected[1]][k.Selected[0]]
	return key.Code
}

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
)

func RGB565ToRGBA(c uint16) color.RGBA {
	return color.RGBA{
		R: uint8((c & 0xF800) >> 8),
		G: uint8((c & 0x07E0) >> 3),
		B: uint8((c & 0x001F) << 3),
		A: 0xFF,
	}
}
