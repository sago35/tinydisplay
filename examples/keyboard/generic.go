package main

import (
	"log"
	"time"

	"github.com/sago35/tinydisplay"
)

var (
	display *tinydisplay.Client
)

func init() {
	var err error
	display, err = tinydisplay.NewClient("", 9812, 320, 240)
	if err != nil {
		log.Fatal(err)
	}

	display.FillScreen(white)
	time.Sleep(100 * time.Millisecond)
}
