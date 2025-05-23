package tinydisplay

import (
	"encoding/gob"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"net/rpc"
	"sync"
	"time"

	"github.com/sago35/tinydisplay/defines"
	"tinygo.org/x/drivers/touch"
)

type Client struct {
	Client *rpc.Client
	draw.Image
	mu     sync.Mutex
	Width  int
	Height int
}

func init() {
	gob.Register(&image.RGBA{})
}

func NewClient(addr string, port, w, h int) (*Client, error) {
	c, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return nil, err
	}
	client := &Client{
		Client: c,
		Image:  image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{w, h}}),
		Width:  w,
		Height: h,
	}

	go client.Tick()
	return client, nil
}

func (c *Client) Tick() {
	for {
		c.mu.Lock()
		c.update()
		c.mu.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
}

func (c *Client) Size() (x, y int16) {
	args := defines.NotImpl{}
	ret := defines.SizeRetval{}
	err := c.Client.Call("Server.Size", args, &ret)
	if err != nil {
		panic(err)
	}
	return ret.X, ret.Y
}

func (c *Client) SetPixel(x, y int16, clr color.RGBA) {
	c.Set(int(x), int(y), clr)
}

func (c *Client) Display() error {
	return c.update()
}

func (c *Client) ClearBuffer() {
	c.FillScreen(color.RGBA{0, 0, 0, 0xFF})
}

func (c *Client) ClearDisplay() {
	c.ClearBuffer()
}

func (c *Client) WaitUntilIdle() {
}

func (c *Client) FillScreen(clr color.Color) {
	c.FillRectangle(0, 0, int16(c.Width), int16(c.Height), clr)
}

func (c *Client) FillRectangle(x, y, width, height int16, clr color.Color) error {
	for yy := y; yy < y+height; yy++ {
		for xx := x; xx < x+width; xx++ {
			c.Image.Set(int(xx), int(yy), clr)
		}
	}
	return nil
}

func (c *Client) DrawRectangle(x, y, width, height int16, clr color.Color) error {
	for yy := y; yy < y+height; yy++ {
		c.Image.Set(int(x), int(yy), clr)
		c.Image.Set(int(x+width-1), int(yy), clr)
	}

	for xx := x; xx < x+width; xx++ {
		c.Image.Set(int(xx), int(y), clr)
		c.Image.Set(int(xx), int(y+height-1), clr)
	}
	return nil
}

func (c *Client) DrawRGBBitmap(x, y int16, data []uint16, w, h int16) error {
	index := 0
	for yy := y; yy < y+h; yy++ {
		for xx := x; xx < x+w; xx++ {
			rgb565 := data[index]
			c.Image.Set(int(xx), int(yy), RGB565ToRGBA(rgb565))
			index += 1
		}
	}
	return nil
}

func (c *Client) DrawRGBBitmap8(x, y int16, data []uint8, w, h int16) error {
	index := 0
	for yy := y; yy < y+h; yy++ {
		for xx := x; xx < x+w; xx++ {
			rgb565 := uint16(data[index])<<8 + uint16(data[index+1])
			c.Image.Set(int(xx), int(yy), RGB565ToRGBA(rgb565))
			index += 2
		}
	}
	return nil
}

func (c *Client) SetImage(img draw.Image) {
	c.mu.Lock()
	c.Image = img
	c.mu.Unlock()
}

func (c *Client) update() error {
	args := defines.UpdateArgs{Image: c.Image}
	ret := defines.NotImpl{}
	err := c.Client.Call("Server.Update", args, &ret)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *Client) Set(x, y int, clr color.Color) {
	c.mu.Lock()
	c.Image.Set(x, y, clr)
	c.mu.Unlock()
}

func (c *Client) GetPressedKey() uint16 {
	args := defines.NotImpl{}
	ret := defines.GetPressedKeysRetval{}
	c.Client.Call("Server.GetPressedKeys", args, &ret)

	for _, key := range ret.Keys {
		if k, ok := keyMaps[key]; ok {
			return k
		}
	}

	return 0xFFFF
}

func (c *Client) ReadTouchPoint() touch.Point {
	var ret touch.Point
	c.Client.Call("Server.ReadTouchPoint", defines.NotImpl{}, &ret)
	return ret
}
