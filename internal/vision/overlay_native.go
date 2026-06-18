//go:build !stub
// +build !stub

package vision

import (
	"fmt"
	"image"
	"image/color"
	"log/slog"

	"gocv.io/x/gocv"

	"nagare/models"
)

type nativeOverlayImpl struct {
	win    *gocv.Window
	logger *slog.Logger
	config OverlayConfig
	fps    *fpsCounter
}

var (
	landmarkColor      = color.RGBA{R: 91, G: 95, B: 248, A: 255}
	connectionColor    = color.RGBA{R: 67, G: 71, B: 231, A: 255}
	labelColor         = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	activeColor        = color.RGBA{R: 16, G: 185, B: 129, A: 255}
	inactiveColor      = color.RGBA{R: 156, G: 163, B: 175, A: 255}
)

func newOverlayImpl(cfg OverlayConfig, logger *slog.Logger) (overlayImpl, error) {
	if !cfg.Enabled {
		return &noopOverlayImpl{}, nil
	}

	win := gocv.NewWindow(cfg.WindowName)
	logger.Info("debug overlay window created", "name", cfg.WindowName)

	return &nativeOverlayImpl{
		win:    win,
		logger: logger,
		config: cfg,
		fps:    newFPSCounter(),
	}, nil
}

func (o *nativeOverlayImpl) annotate(frame *ProcessedFrame, data *models.HandData, tracking bool) (*ProcessedFrame, error) {
	mat, err := gocv.NewMatFromBytes(frame.Height, frame.Width, gocv.MatTypeCV8UC3, frame.Data)
	if err != nil {
		return nil, fmt.Errorf("create mat: %w", err)
	}
	defer mat.Close()

	if data != nil && data.Confidence >= 0.5 {
		o.drawLandmarks(&mat, data)
	}

	o.drawStatus(&mat, data, tracking)

	return &ProcessedFrame{
		Data:   mat.ToBytes(),
		Width:  frame.Width,
		Height: frame.Height,
	}, nil
}

func (o *nativeOverlayImpl) drawLandmarks(mat *gocv.Mat, data *models.HandData) {
	for _, conn := range handConnections {
		a := data.Landmarks[conn[0]]
		b := data.Landmarks[conn[1]]
		pt1 := image.Pt(int(a.X), int(a.Y))
		pt2 := image.Pt(int(b.X), int(b.Y))
		gocv.Line(mat, pt1, pt2, connectionColor, 2)
	}

	for i, lm := range data.Landmarks {
		pt := image.Pt(int(lm.X), int(lm.Y))
		gocv.Circle(mat, pt, 4, landmarkColor, -1)

		labelPt := image.Pt(int(lm.X)+6, int(lm.Y)-4)
		gocv.PutText(mat, models.LandmarkNames[i], labelPt, gocv.FontHersheySimplex, 0.3, labelColor, 1)
	}
}

func (o *nativeOverlayImpl) drawStatus(mat *gocv.Mat, data *models.HandData, tracking bool) {
	fpsText := fmt.Sprintf("FPS: %d", o.fps.value())
	gocv.PutText(mat, fpsText, image.Pt(12, 28), gocv.FontHersheySimplex, 0.6, labelColor, 2)

	statusColor := inactiveColor
	statusText := "Tracking: OFF"
	if tracking {
		statusText = "Tracking: ON"
		statusColor = activeColor
	}
	gocv.PutText(mat, statusText, image.Pt(12, 54), gocv.FontHersheySimplex, 0.6, statusColor, 2)

	if data != nil {
		confText := fmt.Sprintf("Confidence: %.2f", data.Confidence)
		gocv.PutText(mat, confText, image.Pt(12, 80), gocv.FontHersheySimplex, 0.5, labelColor, 1)

		orientStr := "Right"
		if data.Orientation == models.HandOrientationLeft {
			orientStr = "Left"
		}
		gocv.PutText(mat, "Hand: "+orientStr, image.Pt(12, 100), gocv.FontHersheySimplex, 0.5, labelColor, 1)
	}
}

func (o *nativeOverlayImpl) show(frame *ProcessedFrame) error {
	if o.win == nil || frame == nil || len(frame.Data) == 0 {
		return nil
	}

	mat, err := gocv.NewMatFromBytes(frame.Height, frame.Width, gocv.MatTypeCV8UC3, frame.Data)
	if err != nil {
		return fmt.Errorf("create mat for display: %w", err)
	}
	defer mat.Close()

	o.win.IMShow(mat)
	o.win.WaitKey(1)
	return nil
}

func (o *nativeOverlayImpl) close() error {
	if o.win != nil {
		o.win.Close()
		o.win = nil
		o.logger.Info("debug overlay closed")
	}
	return nil
}
