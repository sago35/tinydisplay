package tinykb

import "tinygo.org/x/drivers"

type Keyboard14x4 struct {
	Layout  [4][14]Key
	Disp    drivers.Displayer
	bgcolor uint16
	fgcolor uint16
	Column  int
	Row     int
	X       int16
	Y       int16
}

func New(display drivers.Displayer, x, y int16) *Keyboard14x4 {
	return &Keyboard14x4{
		Disp: display,
		Layout: [4][14]Key{
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
				Key{Code: KeyClose},
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
		},
		X: x,
		Y: y,
	}
}

func (k *Keyboard14x4) Display() {
	for row := 0; row < len(k.Layout); row++ {
		for col := 0; col < len(k.Layout[0]); col++ {
			if col == k.Column && row == k.Row {
				k.Redraw(col, row, true)
			} else {
				k.Redraw(col, row, false)
			}
		}
	}
}

func (k *Keyboard14x4) Redraw(col, row int, selected bool) {
	btn := k.Layout[row][col]
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
	return k.Layout[k.Row][k.Column]
}

func (k *Keyboard14x4) KeyEvent(key Key) {
	col := k.Column
	row := k.Row
	switch key.Code {
	case KeyRight:
		col = (col + len(k.Layout[0]) + 1) % len(k.Layout[0])
	case KeyLeft:
		col = (col + len(k.Layout[0]) - 1) % len(k.Layout[0])
	case KeyUp:
		row = (row + len(k.Layout) - 1) % len(k.Layout)
	case KeyDown:
		row = (row + len(k.Layout) + 1) % len(k.Layout)
	default:
		return
	}
	k.Redraw(k.Column, k.Row, false)
	k.Column = col
	k.Row = row
	k.Redraw(col, row, true)
}
