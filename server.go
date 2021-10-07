package tinydisplay

import (
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2"
)

type Server struct {
	Device *Device
}

func NewServer(w, h int) *Server {
	s := &Server{
		Device: New(w, h),
	}

	s.Device.FillScreen(color.RGBA{0, 0, 0, 255})

	return s
}

type SizeRetval struct {
	X, Y int16
}

func (s *Server) Size(args *NotImpl, ret *SizeRetval) error {
	ret.X, ret.Y = s.Device.Size()
	return nil
}

type SetPixelArgs struct {
	X, Y int16
	C    color.RGBA
}

func (s *Server) SetPixel(args *SetPixelArgs, ret *NotImpl) error {
	s.Device.SetPixel(args.X, args.Y, args.C)
	return nil
}

func (s *Server) Display(args, ret *NotImpl) error {
	return s.Device.Display()
}

type FillScreenArgs struct {
	C color.RGBA
}

func (s *Server) FillScreen(args *FillScreenArgs, ret *NotImpl) error {
	s.Device.FillScreen(args.C)
	return nil
}

type FillRectangleArgs struct {
	X, Y, Width, Height int16
	C                   color.RGBA
}

func (s *Server) FillRectangle(args *FillRectangleArgs, ret *NotImpl) error {
	return s.Device.FillRectangle(args.X, args.Y, args.Width, args.Height, args.C)
}

type DrawRGBBitmap8Args struct {
	X, Y, W, H int16
	Data       []uint8
}

func (s *Server) DrawRGBBitmap8(args *DrawRGBBitmap8Args, ret *NotImpl) error {
	return s.Device.DrawRGBBitmap8(args.X, args.Y, args.Data, args.W, args.H)
}

func (s *Server) ShowAndRun(args, ret *NotImpl) error {
	s.Device.ShowAndRun()
	return nil
}

type UpdateArgs struct {
	Image draw.Image
}

func (s *Server) Update(args *UpdateArgs, ret *NotImpl) error {
	for y := 0; y < args.Image.Bounds().Dy(); y++ {
		for x := 0; x < args.Image.Bounds().Dx(); x++ {
			s.Device.image.Set(x, y, args.Image.At(x, y))
		}
	}
	s.Device.Update()
	return nil
}

type GetPressedKeysRetval struct {
	Keys []fyne.KeyName
}

func (s *Server) GetPressedKeys(args *NotImpl, ret *GetPressedKeysRetval) error {
	s.Device.mu.Lock()
	for key := range s.Device.KeysPressed {
		ret.Keys = append(ret.Keys, key)
	}
	s.Device.mu.Unlock()
	return nil
}

type NotImpl struct {
}
