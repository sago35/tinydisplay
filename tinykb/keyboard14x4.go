package tinykb

import "tinygo.org/x/drivers"

type Keyboard14x4 struct {
	Layout  [2][4][14]Key
	Disp    drivers.Displayer
	bgcolor uint16
	fgcolor uint16
	Column  int
	Row     int
	X       int16
	Y       int16
	layer   int
}

var Keyboard14x4Layout = [2][4][14]Key{
	{
		{
			'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '=',
			KeyBackspace,
		},
		{
			'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']', '\\',
			KeyTab,
		},
		{
			'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\'',
			KeyReturn,
			KeyUp,
			KeyShift,
		},
		{
			'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/', ' ',
			KeyLeft,
			KeyDown,
			KeyRight,
		},
	},
	{
		{
			'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+',
			KeyBackspace,
		},
		{
			'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '{', '}', '|',
			KeyClose,
		},
		{
			'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ':', '"',
			KeyReturn,
			KeyUp,
			KeyShiftRelease,
		},
		{
			'Z', 'X', 'C', 'V', 'B', 'N', 'M', '<', '>', '?', ' ',
			KeyLeft,
			KeyDown,
			KeyRight,
		},
	},
}

func New(display drivers.Displayer, x, y int16) *Keyboard14x4 {
	return &Keyboard14x4{
		Disp:   display,
		Layout: Keyboard14x4Layout,
		layer:  0,
		X:      x,
		Y:      y,
	}
}

func (k *Keyboard14x4) Display() {
	for row := 0; row < len(k.Layout[k.layer]); row++ {
		for col := 0; col < len(k.Layout[k.layer][0]); col++ {
			if col == k.Column && row == k.Row {
				k.Redraw(col, row, true)
			} else {
				k.Redraw(col, row, false)
			}
		}
	}
}

func (k *Keyboard14x4) Redraw(col, row int, selected bool) {
	btn := k.Layout[k.layer][row][col]
	xxx := k.X + (sz+1)*(int16(col)+0) + sz/2*0
	yyy := k.Y + ((sz + 1) * int16(row))
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

func (k *Keyboard14x4) GetKey() Key {
	return k.Layout[k.layer][k.Row][k.Column]
}

func (k *Keyboard14x4) Layer(index int) {
	k.layer = index
	k.Display()
}

func (k *Keyboard14x4) KeyEvent(key Key) {
	col := k.Column
	row := k.Row
	switch key {
	case KeyRight:
		col = (col + len(k.Layout[k.layer][0]) + 1) % len(k.Layout[k.layer][0])
	case KeyLeft:
		col = (col + len(k.Layout[k.layer][0]) - 1) % len(k.Layout[k.layer][0])
	case KeyUp:
		row = (row + len(k.Layout[k.layer]) - 1) % len(k.Layout[k.layer])
	case KeyDown:
		row = (row + len(k.Layout[k.layer]) + 1) % len(k.Layout[k.layer])
	default:
		return
	}
	k.Redraw(k.Column, k.Row, false)
	k.Column = col
	k.Row = row
	k.Redraw(col, row, true)
}
