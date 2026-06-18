//go:build !stub
// +build !stub

package vision

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"

	"nagare/internal/camera"
)

func (p *Pipeline) Process(raw *camera.Frame) (*ProcessedFrame, error) {
	if raw == nil || len(raw.Data) == 0 {
		return nil, fmt.Errorf("empty frame")
	}

	mat, err := gocv.NewMatFromBytes(raw.Height, raw.Width, gocv.MatTypeCV8UC3, raw.Data)
	if err != nil {
		return nil, fmt.Errorf("create mat from bytes: %w", err)
	}
	defer mat.Close()

	processed := gocv.NewMat()
	defer processed.Close()

	gocv.Resize(mat, &processed, image.Point{
		X: p.config.TargetWidth,
		Y: p.config.TargetHeight,
	}, 0, 0, gocv.InterpolationLinear)

	if p.config.Mirror {
		gocv.Flip(processed, &processed, 1)
	}

	if processed.Channels() == 3 {
		gocv.CvtColor(processed, &processed, gocv.ColorBGRToRGB)
	}

	return &ProcessedFrame{
		Data:   processed.ToBytes(),
		Width:  processed.Cols(),
		Height: processed.Rows(),
	}, nil
}
