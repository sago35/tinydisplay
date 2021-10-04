package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/sago35/tinydisplay/tinykb"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
)

func main() {
	display.FillScreen(black)

	for {
		str, err := run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q\n", str)

		time.Sleep(500 * time.Millisecond)
	}
}

func run() (string, error) {
	var kb tinykb.Keyboard
	kb = tinykb.New(display, 0, 150)
	kb.Display()

	str := ""
	needsRedraw := true
	for {
		key := display.GetPressedKey()
		needsWait := true
		switch key {
		case tinykb.KeyRight, tinykb.KeyLeft, tinykb.KeyUp, tinykb.KeyDown:
			kb.KeyEvent(tinykb.Key(key))
		case tinykb.KeyReturn:
			k := kb.GetKey()
			//fmt.Printf("%s\n", k)
			switch k {
			case tinykb.KeyShift:
				kb.Layer(1)
			case tinykb.KeyShiftRelease:
				kb.Layer(0)
			case tinykb.KeyReturn:
				str += "\n"
			case tinykb.KeyBackspace:
				str = str[:len(str)-1]
			case tinykb.KeyClose:
				return str, nil
			default:
				str += fmt.Sprintf("%c", k)
			}
			needsRedraw = true
		default:
			needsWait = false
		}

		if needsRedraw {
			display.FillRectangle(0, 0, 320, 150, black)
			tinyfont.WriteLine(display, &freemono.Regular9pt7b, 0, 15, str, white)
			needsRedraw = false
		}

		if needsWait {
			time.Sleep(200 * time.Millisecond)
		}
	}
}
