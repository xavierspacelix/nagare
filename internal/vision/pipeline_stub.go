//go:build !windows && !darwin
// +build !windows,!darwin

package vision

import (
	"bytes"
	"fmt"
	"image"

	"nagare/internal/camera"
)

func (p *Pipeline) Process(raw *camera.Frame) (*ProcessedFrame, error) {
	if raw == nil || len(raw.Data) == 0 {
		return nil, fmt.Errorf("empty frame")
	}

	img, _, err := image.Decode(bytes.NewReader(raw.Data))
	if err != nil {
		return nil, fmt.Errorf("decode frame: %w", err)
	}

	processed := resizeImage(img, p.config.TargetWidth, p.config.TargetHeight)

	bounds := processed.Bounds()
	rgb := make([]byte, bounds.Dx()*bounds.Dy()*3)
	idx := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := processed.At(x, y).RGBA()
			rgb[idx] = byte(r >> 8)
			rgb[idx+1] = byte(g >> 8)
			rgb[idx+2] = byte(b >> 8)
			idx += 3
		}
	}

	return &ProcessedFrame{
		Data:   rgb,
		Width:  bounds.Dx(),
		Height: bounds.Dy(),
	}, nil
}

func resizeImage(img image.Image, targetW, targetH int) image.Image {
	bounds := img.Bounds()
	srcW := bounds.Dx()
	srcH := bounds.Dy()

	if srcW == targetW && srcH == targetH {
		return img
	}

	dst := image.NewRGBA(image.Rect(0, 0, targetW, targetH))
	for dy := 0; dy < targetH; dy++ {
		for dx := 0; dx < targetW; dx++ {
			sx := dx * srcW / targetW
			sy := dy * srcH / targetH
			dst.Set(dx, dy, img.At(sx, sy))
		}
	}
	return dst
}
