package pipeline

import (
	"log/slog"
	"testing"

	"nagare/internal/actions"
	"nagare/internal/camera"
	"nagare/internal/controller"
	"nagare/internal/display"
	"nagare/internal/gestures"
	"nagare/internal/profiler"
	"nagare/internal/vision"
)

func newTestRunner(logger *slog.Logger) *Runner {
	cm := camera.NewManager(logger)
	pl := vision.NewPipeline(vision.DefaultConfig())
	ctrl := controller.NewStubController()
	eng := actions.NewEngine(ctrl, logger)
	prof := profiler.New()
	mc := gestures.NewMachine(gestures.DefaultConfig(), eng.HandleGesture, logger)
	mon := []display.Info{{Index: 0, X: 0, Y: 0, Width: 1920, Height: 1080, Primary: true}}
	dm, _ := display.NewMapper(mon)
	return NewRunner(cm, pl, nil, nil, mc, eng, ctrl, prof, dm, logger)
}

func TestNewRunner(t *testing.T) {
	r := newTestRunner(slog.Default())
	if r == nil {
		t.Fatal("expected non-nil runner")
	}
}

func TestRunner_StartStop(t *testing.T) {
	r := newTestRunner(slog.Default())

	if r.IsRunning() {
		t.Fatal("expected not running initially")
	}

	if err := r.Start(); err != nil {
		t.Fatalf("unexpected start error: %v", err)
	}

	if !r.IsRunning() {
		t.Fatal("expected running after start")
	}

	r.Stop()

	if r.IsRunning() {
		t.Fatal("expected not running after stop")
	}
}

func TestRunner_SetTracking(t *testing.T) {
	logger := slog.Default()
	cm := camera.NewManager(logger)
	pl := vision.NewPipeline(vision.DefaultConfig())
	ctrl := controller.NewStubController()
	eng := actions.NewEngine(ctrl, logger)
	prof := profiler.New()
	mc := gestures.NewMachine(gestures.DefaultConfig(), eng.HandleGesture, logger)
	dm, _ := display.NewMapper(nil)

	r := NewRunner(cm, pl, nil, nil, mc, eng, ctrl, prof, dm, logger)
	r.SetTracking(true)

	if !eng.IsTracking() {
		t.Fatal("expected tracking enabled")
	}

	r.SetTracking(false)
	if eng.IsTracking() {
		t.Fatal("expected tracking disabled")
	}
}
