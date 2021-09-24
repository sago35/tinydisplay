package tinydisplay

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Device struct {
	Width  int
	Height int
	canvas *canvas.Image
	window fyne.Window
	image  draw.Image
}

func New(w, h int) *Device {
	a := app.New()
	wi := a.NewWindow("tinydisplay")

	cimage := &canvas.Image{}
	cimage.SetMinSize(fyne.Size{Width: float32(w), Height: float32(h)})
	rgba := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{w, h}})
	cimage.Image = rgba

	wi.SetContent(container.NewVBox(
		cimage,
	))
	canvas := wi.Canvas()
	canvas.SetOnTypedKey(func(ev *fyne.KeyEvent) {
		fmt.Printf("%#v\n", ev)
	})

	return &Device{
		Width:  w,
		Height: h,
		window: wi,
		canvas: cimage,
		image:  rgba,
	}
}

func (d *Device) Size() (x, y int16) {
	return int16(d.Width), int16(d.Height)
}

func (d *Device) SetPixel(x, y int16, c color.Color) {
	d.image.Set(int(x), int(y), c)
	d.canvas.Refresh()
}

func (d *Device) Display() error {
	panic("not impl")
}

func (d *Device) FillScreen(c color.Color) {
	d.FillRectangle(0, 0, int16(d.Width), int16(d.Height), c)
}

func (d *Device) FillRectangle(x, y, width, height int16, c color.Color) error {
	for yy := y; yy < y+height; yy++ {
		for xx := x; xx < x+width; xx++ {
			d.image.Set(int(xx), int(yy), c)
		}
	}
	return nil
}

func (d *Device) DrawRGBBitmap8(x, y int16, data []uint8, w, h int16) error {
	index := 0
	for yy := y; yy < y+h; yy++ {
		for xx := x; xx < x+w; xx++ {
			rgb565 := uint16(data[index])<<8 + uint16(data[index+1])
			d.image.Set(int(xx), int(yy), RGB565ToRGBA(rgb565))
			index += 2
		}
	}
	return nil
}

func (d *Device) Update() error {
	d.canvas.Refresh()
	return nil
}

func (d *Device) ShowAndRun() {
	d.window.ShowAndRun()
}

func RGB565ToRGBA(c uint16) color.Color {
	return color.RGBA{
		R: uint8((c & 0xF800) >> 8),
		G: uint8((c & 0x07E0) >> 3),
		B: uint8((c & 0x001F) << 3),
		A: 0xFF,
	}
}

// RGBATo565 converts a color.RGBA to uint16 used in the display
func RGBATo565(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) +
		((g & 0xFC00) >> 5) +
		((b & 0xF800) >> 11))
}
