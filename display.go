package tinydisplay

import (
	"image"
	"image/color"
	"image/draw"
	"sort"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"tinygo.org/x/drivers/touch"
)

type touchWidget struct {
	widget.BaseWidget
	obj fyne.CanvasObject
	d   *Device
}

func (s *touchWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s.obj)
}

func (s *touchWidget) setTouchPoint(p fyne.Position) {

	x := int(p.X)
	y := int(p.Y)

	//Clamp X & Y values to width and height. If the pointer is moved outside the window these values may be negative or larger than the window
	if x < 0 {
		x = 0
	} else if x >= s.d.Width {
		x = s.d.Width - 1
	}

	if y < 0 {
		y = 0
	} else if y >= s.d.Height {
		y = s.d.Height - 1
	}

	s.d.TouchPoint.X = x * ((1 << 16) / s.d.Width)
	s.d.TouchPoint.Y = y * ((1 << 16) / s.d.Height)
	s.d.TouchPoint.Z = 0xFFFF
}

func (s *touchWidget) Tapped(p *fyne.PointEvent) {
	s.d.mu.Lock()

	s.setTouchPoint(p.Position)

	s.d.mu.Unlock()
}

func (s *touchWidget) Dragged(d *fyne.DragEvent) {
	s.d.mu.Lock()

	s.setTouchPoint(d.Position)
	s.d.DragInProgress = true

	s.d.mu.Unlock()
}

func (s *touchWidget) DragEnd() {
	s.d.mu.Lock()

	//Leave previous touch point set as it may not have been read yet
	s.d.DragInProgress = false

	s.d.mu.Unlock()
}

type Device struct {
	Width          int
	Height         int
	canvas         *canvas.Image
	window         fyne.Window
	image          draw.Image
	KeysPressed    map[fyne.KeyName]bool
	TouchPoint     touch.Point
	DragInProgress bool
	mu             sync.Mutex
}

func New(w, h int) *Device {
	a := app.New()
	wi := a.NewWindow("tinydisplay")

	cimage := &canvas.Image{}
	cimage.SetMinSize(fyne.Size{Width: float32(w), Height: float32(h)})
	rgba := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{w, h}})
	cimage.Image = rgba

	o := &touchWidget{
		obj: cimage,
	}
	o.ExtendBaseWidget(o)

	wi.SetContent(container.NewVBox(
		o,
	))

	d := &Device{
		Width:       w,
		Height:      h,
		window:      wi,
		canvas:      cimage,
		image:       rgba,
		KeysPressed: map[fyne.KeyName]bool{},
	}

	o.d = d

	if wc, ok := wi.Canvas().(desktop.Canvas); ok {
		wc.SetOnKeyDown(func(ev *fyne.KeyEvent) {
			d.mu.Lock()
			d.KeysPressed[ev.Name] = true
			d.mu.Unlock()
			d.DumpPressedKeys()
		})
		wc.SetOnKeyUp(func(ev *fyne.KeyEvent) {
			d.mu.Lock()
			delete(d.KeysPressed, ev.Name)
			d.mu.Unlock()
			d.DumpPressedKeys()
		})

	}

	return d
}

func (d *Device) DumpPressedKeys() {
	keys := []string{}
	for k, _ := range d.KeysPressed {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
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

func RGB565ToRGBA(c uint16) color.RGBA {
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
