package vision

import (
	"testing"

	"nagare/models"
)

func TestNewLandmarkExtractor_Stub(t *testing.T) {
	ext, err := NewLandmarkExtractor(DefaultLandmarkConfig(), nil)
	if err != nil {
		t.Fatal("new extractor:", err)
	}
	defer ext.Close()
}

func TestStubExtract_ValidFrame(t *testing.T) {
	ext, err := NewLandmarkExtractor(DefaultLandmarkConfig(), nil)
	if err != nil {
		t.Fatal("new extractor:", err)
	}
	defer ext.Close()

	frame := &ProcessedFrame{
		Data:   make([]byte, 640*480*3),
		Width:  640,
		Height: 480,
	}

	data, err := ext.Extract(frame)
	if err != nil {
		t.Fatal("extract:", err)
	}

	if data.Confidence != 0.95 {
		t.Fatalf("expected confidence 0.95, got %f", data.Confidence)
	}
	if data.Orientation != models.HandOrientationRight {
		t.Fatalf("expected right hand, got %v", data.Orientation)
	}
}

func TestStubExtract_LandmarkPositions(t *testing.T) {
	ext, err := NewLandmarkExtractor(DefaultLandmarkConfig(), nil)
	if err != nil {
		t.Fatal("new extractor:", err)
	}
	defer ext.Close()

	frame := &ProcessedFrame{
		Data:   make([]byte, 640*480*3),
		Width:  640,
		Height: 480,
	}

	data, err := ext.Extract(frame)
	if err != nil {
		t.Fatal("extract:", err)
	}

	if data.Landmarks[0].X != 320 {
		t.Fatalf("expected wrist X=320, got %f", data.Landmarks[0].X)
	}
	if data.Landmarks[0].Y != 320 {
		t.Fatalf("expected wrist Y=320, got %f", data.Landmarks[0].Y)
	}
}

func TestStubExtract_EmptyFrame(t *testing.T) {
	ext, err := NewLandmarkExtractor(DefaultLandmarkConfig(), nil)
	if err != nil {
		t.Fatal("new extractor:", err)
	}
	defer ext.Close()

	_, err = ext.Extract(nil)
	if err == nil {
		t.Fatal("expected error for nil frame")
	}

	_, err = ext.Extract(&ProcessedFrame{Data: nil})
	if err == nil {
		t.Fatal("expected error for empty frame")
	}
}

func TestComputeFingerStates(t *testing.T) {
	landmarks := &[21]models.HandLandmark{}
	for i := range 21 {
		landmarks[i] = models.HandLandmark{X: 0, Y: float64(200 - i*10), Z: 0}
	}

	fs := computeFingerStates(landmarks)

	if fs.Index() != models.FingerExtended {
		t.Fatal("expected index extended")
	}
	if fs.Middle() != models.FingerExtended {
		t.Fatal("expected middle extended")
	}
}

func TestDetectOrientation(t *testing.T) {
	var right [21]models.HandLandmark
	right[0] = models.HandLandmark{X: 100, Y: 200, Z: 0}
	right[8] = models.HandLandmark{X: 120, Y: 100, Z: 0}
	right[20] = models.HandLandmark{X: 160, Y: 120, Z: 0}

	orientation := detectOrientation(&right)
	if orientation != models.HandOrientationRight {
		t.Fatal("expected right hand orientation")
	}
}

func TestDefaultLandmarkConfig(t *testing.T) {
	cfg := DefaultLandmarkConfig()
	if cfg.InputWidth != 224 {
		t.Fatalf("expected InputWidth 224, got %d", cfg.InputWidth)
	}
	if cfg.InputHeight != 224 {
		t.Fatalf("expected InputHeight 224, got %d", cfg.InputHeight)
	}
	if cfg.ModelPath != "assets/models/hand_landmark.onnx" {
		t.Fatalf("unexpected model path: %s", cfg.ModelPath)
	}
}
