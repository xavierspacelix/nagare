package tray

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

func generateIcon() []byte {
	return generatePNG(color.RGBA{R: 91, G: 95, B: 248, A: 255})
}

func generateGrayIcon() []byte {
	return generatePNG(color.RGBA{R: 156, G: 163, B: 175, A: 255})
}

func generatePNG(c color.RGBA) []byte {
	const size = 32
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			img.Set(x, y, c)
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil
	}
	return buf.Bytes()
}
