package vision

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"testing"

	"nagare/internal/camera"
)

func makeTestFrame(w, h int) *camera.Frame {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 128, 255})
		}
	}
	var buf bytes.Buffer
	jpeg.Encode(&buf, img, nil)
	return &camera.Frame{Data: buf.Bytes(), Width: w, Height: h}
}

func TestProcessFrame(t *testing.T) {
	p := NewPipeline(DefaultConfig())

	frame, err := p.Process(makeTestFrame(20, 20))
	if err != nil {
		t.Fatal("process frame:", err)
	}
	if frame.Width == 0 || frame.Height == 0 {
		t.Fatal("invalid processed frame dimensions")
	}
	if len(frame.Data) == 0 {
		t.Fatal("empty processed frame data")
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.TargetWidth != 640 {
		t.Fatalf("expected width 640, got %d", cfg.TargetWidth)
	}
	if cfg.TargetHeight != 480 {
		t.Fatalf("expected height 480, got %d", cfg.TargetHeight)
	}
	if !cfg.Mirror {
		t.Fatal("expected mirror to be true")
	}
}

func TestCustomResolution(t *testing.T) {
	p := NewPipeline(Config{
		TargetWidth:  32,
		TargetHeight: 24,
		Mirror:       false,
	})

	frame, err := p.Process(makeTestFrame(64, 48))
	if err != nil {
		t.Fatal("process:", err)
	}
	if frame.Width != 32 || frame.Height != 24 {
		t.Fatalf("expected 32x24, got %dx%d", frame.Width, frame.Height)
	}
}

func TestEmptyFrame(t *testing.T) {
	p := NewPipeline(DefaultConfig())

	_, err := p.Process(nil)
	if err == nil {
		t.Fatal("expected error for nil frame")
	}

	_, err = p.Process(&camera.Frame{Data: nil})
	if err == nil {
		t.Fatal("expected error for empty data frame")
	}
}
