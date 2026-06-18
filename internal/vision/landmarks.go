package vision

import (
	"fmt"
	"log/slog"
	"nagare/models"
)

type LandmarkConfig struct {
	ModelPath     string
	Confidence    float64
	InputWidth    int
	InputHeight   int
}

func DefaultLandmarkConfig() LandmarkConfig {
	return LandmarkConfig{
		ModelPath:   "assets/models/hand_landmark.onnx",
		Confidence:  0.5,
		InputWidth:  224,
		InputHeight: 224,
	}
}

type LandmarkExtractor interface {
	Extract(frame *ProcessedFrame) (*models.HandData, error)
	Close() error
}

func NewLandmarkExtractor(cfg LandmarkConfig, logger *slog.Logger) (LandmarkExtractor, error) {
	if logger == nil {
		logger = slog.Default()
	}

	ext, err := newNativeExtractor(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("new landmark extractor: %w", err)
	}
	return ext, nil
}

func computeFingerStates(landmarks *[21]models.HandLandmark) models.FingerStates {
	var fs models.FingerStates

	thumbTip := landmarks[4]
	thumbIP := landmarks[3]
	indexTip := landmarks[8]
	indexPIP := landmarks[6]
	middleTip := landmarks[12]
	middlePIP := landmarks[10]
	ringTip := landmarks[16]
	ringPIP := landmarks[14]
	pinkyTip := landmarks[20]
	pinkyPIP := landmarks[18]
	wrist := landmarks[0]

	if thumbTip.Y < thumbIP.Y-0.02 {
		fs[models.FingerThumb] = models.FingerExtended
	} else {
		fs[models.FingerThumb] = models.FingerFolded
	}

	fs[models.FingerIndex] = fingerState(indexTip.Y, indexPIP.Y, wrist.Y)
	fs[models.FingerMiddle] = fingerState(middleTip.Y, middlePIP.Y, wrist.Y)
	fs[models.FingerRing] = fingerState(ringTip.Y, ringPIP.Y, wrist.Y)
	fs[models.FingerPinky] = fingerState(pinkyTip.Y, pinkyPIP.Y, wrist.Y)

	return fs
}

func fingerState(tipY, pipY, wristY float64) models.FingerState {
	if tipY < pipY-0.02 {
		return models.FingerExtended
	}
	return models.FingerFolded
}

func detectOrientation(landmarks *[21]models.HandLandmark) models.HandOrientation {
	indexTip := landmarks[8]
	pinkyTip := landmarks[20]
	wrist := landmarks[0]

	cross := (indexTip.X-wrist.X)*(pinkyTip.Y-wrist.Y) -
		(indexTip.Y-wrist.Y)*(pinkyTip.X-wrist.X)

	if cross > 0 {
		return models.HandOrientationRight
	}
	return models.HandOrientationLeft
}
