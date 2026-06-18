package vision

import (
	"fmt"
	"log/slog"
	"time"

	"nagare/models"
)

type OverlayConfig struct {
	Enabled    bool
	WindowName string
}

func DefaultOverlayConfig() OverlayConfig {
	return OverlayConfig{
		Enabled:    true,
		WindowName: "Nagare Debug",
	}
}

type DebugOverlay struct {
	config   OverlayConfig
	fps      *fpsCounter
	tracking bool
	logger   *slog.Logger
	impl     overlayImpl
}

type overlayImpl interface {
	annotate(frame *ProcessedFrame, data *models.HandData, tracking bool) (*ProcessedFrame, error)
	show(frame *ProcessedFrame) error
	close() error
}

func NewDebugOverlay(cfg OverlayConfig, logger *slog.Logger) (*DebugOverlay, error) {
	if logger == nil {
		logger = slog.Default()
	}

	d := &DebugOverlay{
		config: cfg,
		fps:    newFPSCounter(),
		logger: logger,
	}

	impl, err := newOverlayImpl(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("new overlay: %w", err)
	}
	d.impl = impl

	return d, nil
}

func (d *DebugOverlay) Annotate(frame *ProcessedFrame, data *models.HandData) (*ProcessedFrame, error) {
	if frame == nil || len(frame.Data) == 0 {
		return nil, fmt.Errorf("annotate: empty frame")
	}
	d.fps.tick()
	return d.impl.annotate(frame, data, d.tracking)
}

func (d *DebugOverlay) Show(frame *ProcessedFrame) error {
	if frame == nil || len(frame.Data) == 0 {
		return fmt.Errorf("show: empty frame")
	}
	return d.impl.show(frame)
}

func (d *DebugOverlay) SetTracking(active bool) {
	d.tracking = active
}

func (d *DebugOverlay) Close() error {
	return d.impl.close()
}

type fpsCounter struct {
	lastTick time.Time
	current  int
	display  int
}

func newFPSCounter() *fpsCounter {
	return &fpsCounter{lastTick: time.Now()}
}

func (f *fpsCounter) tick() {
	f.current++
	elapsed := time.Since(f.lastTick)
	if elapsed >= time.Second {
		f.display = f.current
		f.current = 0
		f.lastTick = time.Now()
	}
}

func (f *fpsCounter) value() int {
	return f.display
}

type noopOverlayImpl struct{}

func (n *noopOverlayImpl) annotate(frame *ProcessedFrame, data *models.HandData, tracking bool) (*ProcessedFrame, error) {
	return frame, nil
}

func (n *noopOverlayImpl) show(frame *ProcessedFrame) error {
	return nil
}

func (n *noopOverlayImpl) close() error {
	return nil
}

var handConnections = [][2]int{
	{0, 1}, {1, 2}, {2, 3}, {3, 4},
	{0, 5}, {5, 6}, {6, 7}, {7, 8},
	{0, 9}, {9, 10}, {10, 11}, {11, 12},
	{0, 13}, {13, 14}, {14, 15}, {15, 16},
	{0, 17}, {17, 18}, {18, 19}, {19, 20},
	{5, 9}, {9, 13}, {13, 17},
}
