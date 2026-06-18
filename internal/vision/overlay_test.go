package vision

import (
	"testing"
	"time"

	"nagare/models"
)

func TestNewDebugOverlay(t *testing.T) {
	cfg := DefaultOverlayConfig()
	cfg.Enabled = false
	d, err := NewDebugOverlay(cfg, nil)
	if err != nil {
		t.Fatal("new overlay:", err)
	}
	defer d.Close()
}

func TestAnnotate_NoData(t *testing.T) {
	d, err := NewDebugOverlay(OverlayConfig{Enabled: false}, nil)
	if err != nil {
		t.Fatal("new overlay:", err)
	}
	defer d.Close()

	frame := &ProcessedFrame{
		Data:   make([]byte, 100*100*3),
		Width:  100,
		Height: 100,
	}

	result, err := d.Annotate(frame, nil)
	if err != nil {
		t.Fatal("annotate:", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if len(result.Data) == 0 {
		t.Fatal("expected non-empty result data")
	}
}

func TestAnnotate_WithHandData(t *testing.T) {
	d, err := NewDebugOverlay(OverlayConfig{Enabled: false}, nil)
	if err != nil {
		t.Fatal("new overlay:", err)
	}
	defer d.Close()

	frame := &ProcessedFrame{
		Data:   make([]byte, 640*480*3),
		Width:  640,
		Height: 480,
	}

	var landmarks [21]models.HandLandmark
	for i := range 21 {
		landmarks[i] = models.HandLandmark{X: 320, Y: 240, Z: 0}
	}

	data := &models.HandData{
		Landmarks:   landmarks,
		Confidence:  0.92,
		Orientation: models.HandOrientationRight,
		Fingers:     models.FingerStates{models.FingerExtended, models.FingerExtended, models.FingerExtended, models.FingerExtended, models.FingerExtended},
	}

	result, err := d.Annotate(frame, data)
	if err != nil {
		t.Fatal("annotate:", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestAnnotate_EmptyFrame(t *testing.T) {
	d, err := NewDebugOverlay(OverlayConfig{Enabled: false}, nil)
	if err != nil {
		t.Fatal("new overlay:", err)
	}
	defer d.Close()

	_, err = d.Annotate(nil, nil)
	if err == nil {
		t.Fatal("expected error for nil frame")
	}
}

func TestSetTracking(t *testing.T) {
	d, err := NewDebugOverlay(OverlayConfig{Enabled: false}, nil)
	if err != nil {
		t.Fatal("new overlay:", err)
	}
	defer d.Close()

	d.SetTracking(true)
	if !d.tracking {
		t.Fatal("expected tracking true")
	}

	d.SetTracking(false)
	if d.tracking {
		t.Fatal("expected tracking false")
	}
}

func TestFPSCounter(t *testing.T) {
	fps := newFPSCounter()

	fps.tick()
	fps.tick()
	fps.tick()

	time.Sleep(1100 * time.Millisecond)

	fps.tick()

	if fps.value() > 0 {
		t.Logf("FPS: %d", fps.value())
	}
}

func TestDefaultOverlayConfig(t *testing.T) {
	cfg := DefaultOverlayConfig()
	if !cfg.Enabled {
		t.Fatal("expected overlay enabled by default")
	}
	if cfg.WindowName != "Nagare Debug" {
		t.Fatalf("unexpected window name: %s", cfg.WindowName)
	}
}

func TestShow_NilFrame(t *testing.T) {
	d, err := NewDebugOverlay(OverlayConfig{Enabled: false}, nil)
	if err != nil {
		t.Fatal("new overlay:", err)
	}
	defer d.Close()

	err = d.Show(nil)
	if err == nil {
		t.Fatal("expected error for nil frame")
	}
}
