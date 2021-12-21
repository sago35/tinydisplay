package main

import (
	"image/color"

	"tinygo.org/x/drivers/touch"
)

var (
	resistiveTouch touch.Pointer

	white   = color.RGBA{255, 255, 255, 255}
	black   = color.RGBA{0, 0, 0, 255}
	red     = color.RGBA{255, 0, 0, 255}
	green   = color.RGBA{0, 255, 0, 255}
	blue    = color.RGBA{0, 0, 255, 255}
	magenta = color.RGBA{255, 0, 255, 255}
	yellow  = color.RGBA{255, 255, 0, 255}
	cyan    = color.RGBA{0, 255, 255, 255}

	oldColor     color.RGBA
	currentColor color.RGBA
)

const (
	penRadius = 3
	boxSize   = 30

	Xmin = 0
	Xmax = 0xFFFF
	Ymin = 0
	Ymax = 0xFFFF
)

func main() {
	display, resistiveTouch = InitDisplayAndTouch()

	// fill the background and activate the backlight
	width, height := display.Size()
	display.FillRectangle(0, 0, width, height, black)

	// make color selection boxes
	display.FillRectangle(0, 0, boxSize, boxSize, red)
	display.FillRectangle(boxSize, 0, boxSize, boxSize, yellow)
	display.FillRectangle(boxSize*2, 0, boxSize, boxSize, green)
	display.FillRectangle(boxSize*3, 0, boxSize, boxSize, cyan)
	display.FillRectangle(boxSize*4, 0, boxSize, boxSize, blue)
	display.FillRectangle(boxSize*5, 0, boxSize, boxSize, magenta)
	display.FillRectangle(boxSize*6, 0, boxSize, boxSize, black)
	display.FillRectangle(boxSize*7, 0, boxSize, boxSize, white)

	// set the initial color to red and draw a box to highlight it
	oldColor = red
	currentColor = red
	display.DrawRectangle(0, 0, boxSize, boxSize, white)

	for {

		point := resistiveTouch.ReadTouchPoint()
		touch := touch.Point{}
		if point.Z>>6 > 100 {
			rawX := mapval(point.X, Xmin, Xmax, 0, 320)
			rawY := mapval(point.Y, Ymin, Ymax, 0, 240)
			touch.X = rawX
			touch.Y = rawY
			touch.Z = 1
		} else {
			touch.X = 0
			touch.Y = 0
			touch.Z = 0
		}

		if touch.Z > 0 {
			HandleTouch(touch)
		}
	}
}

// based on Arduino's "map" function
func mapval(x int, inMin int, inMax int, outMin int, outMax int) int {
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}

func HandleTouch(touch touch.Point) {

	if int16(touch.Y) < boxSize {
		oldColor = currentColor
		x := int16(touch.X)
		switch {
		case x < boxSize:
			currentColor = red
		case x < boxSize*2:
			currentColor = yellow
		case x < boxSize*3:
			currentColor = green
		case x < boxSize*4:
			currentColor = cyan
		case x < boxSize*5:
			currentColor = blue
		case x < boxSize*6:
			currentColor = magenta
		case x < boxSize*7:
			currentColor = black
		case x < boxSize*8:
			currentColor = white
		}

		if oldColor == currentColor {
			return
		}

		display.DrawRectangle((x/boxSize)*boxSize, 0, boxSize, boxSize, white)
		switch oldColor {
		case red:
			x = 0
		case yellow:
			x = boxSize
		case green:
			x = boxSize * 2
		case cyan:
			x = boxSize * 3
		case blue:
			x = boxSize * 4
		case magenta:
			x = boxSize * 5
		case black:
			x = boxSize * 6
		case white:
			x = boxSize * 7
		}
		display.FillRectangle(int16(x), 0, boxSize, boxSize, oldColor)

	}

	if (int16(touch.Y) - penRadius) > boxSize {
		display.FillRectangle(
			int16(touch.X), int16(touch.Y), penRadius*2, penRadius*2, currentColor)
	}
}
