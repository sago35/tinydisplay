package main

import (
	"image/color"
	"log"

	"github.com/sago35/tinydisplay"
	"tinygo.org/x/tinydraw"
)

func main() {
	d, err := tinydisplay.NewClient("127.0.0.1", 9812, 320, 240)
	if err != nil {
		log.Fatal(err)
	}
	display := *d

	//white := color.RGBA{0, 0, 0, 255}
	yellow := color.RGBA{255, 0, 0, 255}
	black := color.RGBA{1, 1, 1, 255}

	display.ClearBuffer()
	display.ClearDisplay()
	display.FillScreen(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})

	tinydraw.Line(&display, 10, 10, 94, 10, black)
	tinydraw.Line(&display, 94, 16, 10, 16, yellow)
	tinydraw.Line(&display, 10, 20, 10, 202, yellow)
	tinydraw.Line(&display, 16, 202, 16, 20, black)

	tinydraw.Line(&display, 40, 40, 80, 80, black)
	tinydraw.Line(&display, 40, 40, 80, 70, black)
	tinydraw.Line(&display, 40, 40, 80, 60, black)
	tinydraw.Line(&display, 40, 40, 80, 50, black)
	tinydraw.Line(&display, 40, 40, 80, 40, black)

	tinydraw.Line(&display, 100, 100, 40, 100, yellow)
	tinydraw.Line(&display, 100, 100, 40, 90, yellow)
	tinydraw.Line(&display, 100, 100, 40, 80, yellow)
	tinydraw.Line(&display, 100, 100, 40, 70, yellow)
	tinydraw.Line(&display, 100, 100, 40, 60, yellow)
	tinydraw.Line(&display, 100, 100, 40, 50, yellow)

	tinydraw.Rectangle(&display, 30, 120, 20, 20, black)
	tinydraw.FilledRectangle(&display, 34, 124, 12, 12, yellow)

	tinydraw.Circle(&display, 52, 180, 20, black)
	tinydraw.FilledCircle(&display, 52, 180, 16, yellow)

	tinydraw.Triangle(&display, 60, 110, 100, 130, 84, 150, black)
	tinydraw.FilledTriangle(&display, 65, 114, 96, 130, 84, 145, yellow)

	display.Display()
	display.WaitUntilIdle()
	println("You could remove power now")
}
