package main

import (
	"image/color"
	"time"

	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
)

func main() {
	display := initdisplay.InitDisplay()
	width, height := display.Size()

	display.FillScreen(black)

	display.FillRectangle(0, 0, width/2, height/2, white)
	display.FillRectangle(width/2, 0, width/2, height/2, red)
	display.FillRectangle(0, height/2, width/2, height/2, green)
	display.FillRectangle(width/2, height/2, width/2, height/2, blue)
	display.FillRectangle(width/4, height/4, width/2, height/2, black)
	for {
		time.Sleep(time.Hour)
	}

}
