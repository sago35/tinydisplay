package defines

import (
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2"
)

type SizeRetval struct {
	X, Y int16
}

type SetPixelArgs struct {
	X, Y int16
	C    color.RGBA
}

type FillScreenArgs struct {
	C color.RGBA
}

type FillRectangleArgs struct {
	X, Y, Width, Height int16
	C                   color.RGBA
}

type DrawRGBBitmap8Args struct {
	X, Y, W, H int16
	Data       []uint8
}

type UpdateArgs struct {
	Image draw.Image
}

type GetPressedKeysRetval struct {
	Keys []fyne.KeyName
}

type NotImpl struct {
}
