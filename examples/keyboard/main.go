package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/sago35/tinydisplay/tinykb"
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
	display.FillScreen(black)

	tinyfont.WriteLine(display, &freesans.Bold18pt7b, 10, 35, "HELLO", white)

	x := 0
	y := 0
	for {
		key := display.GetPressedKey()
		needsWait := true
		switch key {
		case 0x106: // Right
			x = x + 1
		case 0x107: // Left
			if x > 0 {
				x = x - 1
			}
		case 0x108: // Down
			if y < 4 {
				y = y + 1
			}
		case 0x109: // Up
			if y > 0 {
				y = y - 1
			}
		case 0x101: // Return
			fmt.Printf("%c\n", k14x4.GetKey())
		default:
			needsWait = false
		}
		//if '1' <= key && key <= '0' || 'a' <= key && key <= 'z' || 'A' <= key && key <= 'Z' {
		drawKeyboard(x, y)
		if needsWait {
			time.Sleep(200 * time.Millisecond)
		}
		//time.Sleep(time.Second)
	}
}

func drawKeyboard(x, y int) {
	//drawNumKeys(index)
	//drawQwertyKeys(x, y)
	drawKeyboard14x4(x, y)
}

var k14x4 *tinykb.Keyboard14x4

func drawKeyboard14x4(xx, yy int) {
	if k14x4 == nil {
		k14x4 = tinykb.New(display)
		k14x4.Display()
	}
	k14x4.Redraw(xx, yy, true)
}

func drawNumKeys(index uint16) {
	var buttons [10]*Key
	for i := range buttons {
		buttons[i] = NewKey(32, 32, white)
	}

	for i, b := range buttons {
		x := 10 + (int16(i)%3)*35
		y := 80 + (int16(i)/3)*35
		if int(index) == (i+1)%10 {
			buttons[i].SetText(fmt.Sprintf("%d", (i+1)%10), red)
		} else {
			buttons[i].SetText(fmt.Sprintf("%d", (i+1)%10), black)
		}
		display.DrawRGBBitmap(x, y, b.buf, b.w, b.h)
	}
}

func drawQwertyKeys(xx, yy int) {
	var buttons [512]*Key
	sz := 22
	szi16 := int16(sz)

	//if 'A' <= index && index <= 'Z' {
	//	index = index - 'A' + 'a'
	//}
	//fmt.Printf("%d %d\n", xx, yy)

	ybase := int16(150)

	var x int16
	var y int16

	x = 0
	y = 0
	for _, b := range "`1234567890-=\u0103" {
		buttons[b] = NewKey(sz, sz, white)
		if x == int16(xx) && y == int16(yy) {
			buttons[b].SetText(string(b), red)
		} else {
			buttons[b].SetText(string(b), black)
		}
		display.DrawRGBBitmap(1+(szi16+1)*(x+0)+szi16/2*0, ybase+((szi16+1)*y), buttons[b].buf, buttons[b].w, buttons[b].h)
		x++
	}

	x = 0
	y = 1
	for _, b := range "qwertyuiop[]\\\u0102" {
		buttons[b] = NewKey(sz, sz, white)
		if x == int16(xx) && y == int16(yy) {
			buttons[b].SetText(string(b), red)
		} else {
			buttons[b].SetText(string(b), black)
		}
		display.DrawRGBBitmap(1+(szi16+1)*(x+0)+szi16/2*0, ybase+((szi16+1)*y), buttons[b].buf, buttons[b].w, buttons[b].h)
		x++
	}

	x = 0
	y = 2
	for _, b := range "asdfghjkl;'\u0101\u0109\u01FF" {
		buttons[b] = NewKey(sz, sz, white)
		if x == int16(xx) && y == int16(yy) {
			buttons[b].SetText(string(b), red)
		} else {
			buttons[b].SetText(string(b), black)
		}
		display.DrawRGBBitmap(1+(szi16+1)*(x+0)+szi16/2*0, ybase+((szi16+1)*y), buttons[b].buf, buttons[b].w, buttons[b].h)
		x++
	}

	x = 0
	y = 3
	for _, b := range "zxcvbnm,./ \u0107\u0108\u0106" {
		buttons[b] = NewKey(sz, sz, white)
		if x == int16(xx) && y == int16(yy) {
			buttons[b].SetText(string(b), red)
		} else {
			buttons[b].SetText(string(b), black)
		}
		display.DrawRGBBitmap(1+(szi16+1)*(x+0)+szi16/2*0, ybase+((szi16+1)*y), buttons[b].buf, buttons[b].w, buttons[b].h)
		x++
	}

	x = 1
	y = 4
	key3u := func(keyName string, keyCode uint16, x, y int16, pressed bool) {
		buttons[keyCode] = NewKey(sz*3+2, sz, white)
		if pressed {
			buttons[keyCode].SetText(keyName, red)
		} else {
			buttons[keyCode].SetText(keyName, black)
		}
		display.DrawRGBBitmap(1+(szi16+1)*x+szi16/2*0, ybase+((szi16+1)*y), buttons[keyCode].buf, buttons[keyCode].w, buttons[keyCode].h)
	}
	if false {
		//key3u("`", 0x100, 0, 0, false)
		key3u("Tab", 0x100, 0, 1, false)
		key3u("Esc", 0x100, 0, 2, false)
		key3u("Shift", 0x100, 0, 3, false)
		key3u("Ctrl", 0x100, 0, 4, false)
	} else if false {
		for _, bb := range []string{"Right"} {
			b := 0x100
			buttons[b] = NewKey(sz*3+2, sz, white)
			if x == int16(xx) && y == int16(yy) {
				buttons[b].SetText(bb, red)
			} else {
				buttons[b].SetText(bb, black)
			}
			display.DrawRGBBitmap(1+(szi16+1)*x+szi16/2*0, ybase+((szi16+1)*y), buttons[b].buf, buttons[b].w, buttons[b].h)
			x++
		}
	}
}
