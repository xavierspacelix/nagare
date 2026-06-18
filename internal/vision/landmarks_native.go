//go:build !stub
// +build !stub

package vision

import (
	"fmt"
	"image"
	"log/slog"
	"os"

	"gocv.io/x/gocv"

	"nagare/models"
)

type dnnExtractor struct {
	net    *gocv.Net
	config LandmarkConfig
	logger *slog.Logger
}

func newNativeExtractor(cfg LandmarkConfig, logger *slog.Logger) (LandmarkExtractor, error) {
	if _, err := os.Stat(cfg.ModelPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("model not found at %s: %w", cfg.ModelPath, err)
	}

	net := gocv.ReadNetFromONNX(cfg.ModelPath)
	if net.Empty() {
		return nil, fmt.Errorf("load model from %s: empty network", cfg.ModelPath)
	}

	logger.Info("hand landmark model loaded", "path", cfg.ModelPath)
	return &dnnExtractor{
		net:    &net,
		config: cfg,
		logger: logger,
	}, nil
}

func (e *dnnExtractor) Extract(frame *ProcessedFrame) (*models.HandData, error) {
	if frame == nil || len(frame.Data) == 0 {
		return nil, fmt.Errorf("empty frame")
	}

	mat, err := gocv.NewMatFromBytes(frame.Height, frame.Width, gocv.MatTypeCV8UC3, frame.Data)
	if err != nil {
		return nil, fmt.Errorf("create mat: %w", err)
	}
	defer mat.Close()

	blob := gocv.BlobFromImage(
		mat,
		1.0/255.0,
		image.Pt(e.config.InputWidth, e.config.InputHeight),
		gocv.NewScalar(0, 0, 0, 0),
		false,
		false,
	)
	defer blob.Close()

	e.net.SetInput(blob, "image")

	lm := e.net.Forward("landmarks")
	defer lm.Close()

	sc := e.net.Forward("scores")
	defer sc.Close()

	data, err := lm.DataPtrFloat32()
	if err != nil {
		return nil, fmt.Errorf("get landmarks data: %w", err)
	}

	if len(data) < 63 {
		return nil, fmt.Errorf("unexpected landmarks size: %d, expected at least 63", len(data))
	}

	var landmarks [21]models.HandLandmark
	for i := range 21 {
		landmarks[i] = models.HandLandmark{
			X: float64(data[i*3]),
			Y: float64(data[i*3+1]),
			Z: float64(data[i*3+2]),
		}
	}

	scores, err := sc.DataPtrFloat32()
	if err != nil {
		return nil, fmt.Errorf("get scores data: %w", err)
	}

	confidence := 0.5
	if len(scores) > 0 {
		confidence = float64(scores[0])
	}

	handData := &models.HandData{
		Landmarks:   landmarks,
		Confidence:  confidence,
		Orientation: detectOrientation(&landmarks),
		Fingers:     computeFingerStates(&landmarks),
	}

	return handData, nil
}

func (e *dnnExtractor) Close() error {
	if e.net != nil {
		e.net.Close()
		e.net = nil
	}
	return nil
}
