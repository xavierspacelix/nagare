//go:build !windows && !darwin
// +build !windows,!darwin

package camera

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"time"
)

type stubDevice struct {
	id    int
	frame int
}

func newDevice() CamDevice {
	return &stubDevice{}
}

func discoverCameras() []Info {
	return []Info{
		{ID: 0, Name: "Camera 0 (stub)"},
		{ID: 1, Name: "Camera 1 (stub)"},
	}
}

func (d *stubDevice) Open(id int) error {
	d.id = id
	return nil
}

func (d *stubDevice) Read() (*Frame, error) {
	d.frame++
	time.Sleep(30 * time.Millisecond)
	img := image.NewRGBA(image.Rect(0, 0, 640, 480))
	c := color.RGBA{R: uint8(d.frame % 256), G: 100, B: 200, A: 255}
	for x := 0; x < 640; x++ {
		for y := 0; y < 480; y++ {
			img.Set(x, y, c)
		}
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, nil); err != nil {
		return nil, fmt.Errorf("encode frame: %w", err)
	}
	return &Frame{Data: buf.Bytes(), Width: 640, Height: 480}, nil
}

func (d *stubDevice) Close() error {
	return nil
}
