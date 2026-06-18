//go:build stub
// +build stub

package vision

import (
	"fmt"
	"log/slog"

	"nagare/models"
)

type stubExtractor struct {
	config LandmarkConfig
	logger *slog.Logger
}

func newNativeExtractor(cfg LandmarkConfig, logger *slog.Logger) (LandmarkExtractor, error) {
	logger.Info("using stub landmark extractor")
	return &stubExtractor{
		config: cfg,
		logger: logger,
	}, nil
}

func (e *stubExtractor) Extract(frame *ProcessedFrame) (*models.HandData, error) {
	if frame == nil || len(frame.Data) == 0 {
		return nil, fmt.Errorf("empty frame")
	}

	var landmarks [21]models.HandLandmark

	centerX := float64(frame.Width) / 2.0
	centerY := float64(frame.Height) / 2.0

	landmarks[0] = models.HandLandmark{X: centerX, Y: centerY + 80, Z: 0}        // wrist
	landmarks[1] = models.HandLandmark{X: centerX - 30, Y: centerY + 50, Z: 10}  // thumb_cmc
	landmarks[2] = models.HandLandmark{X: centerX - 40, Y: centerY + 20, Z: 20}  // thumb_mcp
	landmarks[3] = models.HandLandmark{X: centerX - 45, Y: centerY - 10, Z: 30}  // thumb_ip
	landmarks[4] = models.HandLandmark{X: centerX - 50, Y: centerY - 35, Z: 40}  // thumb_tip
	landmarks[5] = models.HandLandmark{X: centerX - 10, Y: centerY + 30, Z: 5}   // index_mcp
	landmarks[6] = models.HandLandmark{X: centerX - 5, Y: centerY, Z: 15}        // index_pip
	landmarks[7] = models.HandLandmark{X: centerX, Y: centerY - 20, Z: 25}       // index_dip
	landmarks[8] = models.HandLandmark{X: centerX + 2, Y: centerY - 45, Z: 35}   // index_tip
	landmarks[9] = models.HandLandmark{X: centerX + 5, Y: centerY + 25, Z: 5}    // middle_mcp
	landmarks[10] = models.HandLandmark{X: centerX + 8, Y: centerY - 5, Z: 15}   // middle_pip
	landmarks[11] = models.HandLandmark{X: centerX + 10, Y: centerY - 30, Z: 25} // middle_dip
	landmarks[12] = models.HandLandmark{X: centerX + 12, Y: centerY - 55, Z: 35} // middle_tip
	landmarks[13] = models.HandLandmark{X: centerX + 20, Y: centerY + 20, Z: 5}  // ring_mcp
	landmarks[14] = models.HandLandmark{X: centerX + 22, Y: centerY - 8, Z: 15}  // ring_pip
	landmarks[15] = models.HandLandmark{X: centerX + 24, Y: centerY - 32, Z: 25} // ring_dip
	landmarks[16] = models.HandLandmark{X: centerX + 25, Y: centerY - 50, Z: 32} // ring_tip
	landmarks[17] = models.HandLandmark{X: centerX + 32, Y: centerY + 15, Z: 5}  // pinky_mcp
	landmarks[18] = models.HandLandmark{X: centerX + 35, Y: centerY - 5, Z: 12}  // pinky_pip
	landmarks[19] = models.HandLandmark{X: centerX + 38, Y: centerY - 20, Z: 20} // pinky_dip
	landmarks[20] = models.HandLandmark{X: centerX + 40, Y: centerY - 30, Z: 25} // pinky_tip

	return &models.HandData{
		Landmarks:   landmarks,
		Confidence:  0.95,
		Orientation: models.HandOrientationRight,
		Fingers:     computeFingerStates(&landmarks),
	}, nil
}

func (e *stubExtractor) Close() error {
	return nil
}
