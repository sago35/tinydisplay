package main

import (
	"fmt"
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
	display.FillScreen(black)

	tinyfont.WriteLine(display, &freesans.Bold18pt7b, 10, 35, "HELLO", white)

	index := 0
	for {
		drawNumKeys(index)
		index = (index + 1) % 10
		time.Sleep(time.Second)
	}
}

func drawNumKeys(index int) {
	var buttons [10]*Key
	for i := range buttons {
		buttons[i] = NewKey(32, 32, white)
	}

	for i, b := range buttons {
		x := 10 + (int16(i)%3)*35
		y := 80 + (int16(i)/3)*35
		if index == i {
			buttons[i].SetText(fmt.Sprintf("%d", (i+1)%10), red)
		} else {
			buttons[i].SetText(fmt.Sprintf("%d", (i+1)%10), black)
		}
		display.DrawRGBBitmap(x, y, b.buf, b.w, b.h)
	}
}
