package server

import (
	"image/color"
	"sort"
	"strings"

	"github.com/sago35/tinydisplay/defines"
	"tinygo.org/x/drivers/touch"
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

func (s *Server) Size(args *defines.NotImpl, ret *defines.SizeRetval) error {
	ret.X, ret.Y = s.Device.Size()
	return nil
}

func (s *Server) SetPixel(args *defines.SetPixelArgs, ret *defines.NotImpl) error {
	s.Device.SetPixel(args.X, args.Y, args.C)
	return nil
}

func (s *Server) Display(args, ret *defines.NotImpl) error {
	return s.Device.Display()
}

func (s *Server) FillScreen(args *defines.FillScreenArgs, ret *defines.NotImpl) error {
	s.Device.FillScreen(args.C)
	return nil
}

func (s *Server) FillRectangle(args *defines.FillRectangleArgs, ret *defines.NotImpl) error {
	return s.Device.FillRectangle(args.X, args.Y, args.Width, args.Height, args.C)
}

func (s *Server) DrawRGBBitmap8(args *defines.DrawRGBBitmap8Args, ret *defines.NotImpl) error {
	return s.Device.DrawRGBBitmap8(args.X, args.Y, args.Data, args.W, args.H)
}

func (s *Server) ShowAndRun(args, ret *defines.NotImpl) error {
	s.Device.ShowAndRun()
	return nil
}

func (s *Server) Update(args *defines.UpdateArgs, ret *defines.NotImpl) error {
	for y := 0; y < args.Image.Bounds().Dy(); y++ {
		for x := 0; x < args.Image.Bounds().Dx(); x++ {
			s.Device.image.Set(x, y, args.Image.At(x, y))
		}
	}
	s.Device.Update()
	return nil
}

func (s *Server) GetPressedKeys(args *defines.NotImpl, ret *defines.GetPressedKeysRetval) error {
	s.Device.mu.Lock()
	for key := range s.Device.KeysPressed {
		ret.Keys = append(ret.Keys, key)
	}
	sort.Slice(ret.Keys, func(i, j int) bool {
		return strings.Compare(string(ret.Keys[i]), string(ret.Keys[j])) < 0
	})
	s.Device.mu.Unlock()
	return nil
}

func (s *Server) ReadTouchPoint(args *defines.NotImpl, ret *touch.Point) error {

	s.Device.mu.Lock()

	*ret = s.Device.TouchPoint
	if !s.Device.DragInProgress {
		s.Device.TouchPoint.X = 0
		s.Device.TouchPoint.Y = 0
		s.Device.TouchPoint.Z = 0
	}

	s.Device.mu.Unlock()

	return nil
}
