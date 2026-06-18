package pipeline

import (
	"log/slog"
	"time"

	"nagare/internal/actions"
	"nagare/internal/camera"
	"nagare/internal/controller"
	"nagare/internal/display"
	"nagare/internal/gestures"
	"nagare/internal/profiler"
	"nagare/internal/vision"
	"nagare/models"
)

type Runner struct {
	camera    *camera.Manager
	pipeline  *vision.Pipeline
	landmark  vision.LandmarkExtractor
	overlay   *vision.DebugOverlay
	machine   *gestures.Machine
	engine   *actions.Engine
	ctrl      controller.OSController
	profiler  *profiler.Profiler
	display   *display.Mapper
	logger    *slog.Logger

	running  bool
	tracking bool
	stopCh   chan struct{}
}

func NewRunner(
	cm *camera.Manager,
	pl *vision.Pipeline,
	lm vision.LandmarkExtractor,
	ov *vision.DebugOverlay,
	mc *gestures.Machine,
	eng *actions.Engine,
	ctrl controller.OSController,
	prof *profiler.Profiler,
	dm *display.Mapper,
	logger *slog.Logger,
) *Runner {
	if logger == nil {
		logger = slog.Default()
	}
	return &Runner{
		camera:   cm,
		pipeline: pl,
		landmark: lm,
		overlay:  ov,
		machine:  mc,
		engine:   eng,
		ctrl:     ctrl,
		profiler: prof,
		display:  dm,
		logger:   logger,
		stopCh:   make(chan struct{}),
	}
}

func (r *Runner) Start() error {
	if r.running {
		return nil
	}

	if !r.camera.IsOpen() {
		if err := r.camera.Open(0); err != nil {
			return err
		}
	}

	if r.display != nil {
		if err := r.display.Refresh(); err != nil {
			r.logger.Warn("display refresh failed", "error", err)
		}
	}

	r.running = true
	r.stopCh = make(chan struct{})
	go r.loop()
	r.logger.Info("pipeline started")
	return nil
}

func (r *Runner) Stop() {
	if !r.running {
		return
	}
	r.running = false
	close(r.stopCh)
	r.camera.Close()
	r.logger.Info("pipeline stopped")
}

func (r *Runner) IsRunning() bool {
	return r.running
}

func (r *Runner) loop() {
	for {
		select {
		case <-r.stopCh:
			return
		default:
		}

		r.profiler.FrameStart()

		raw, err := r.camera.Read()
		if err != nil {
			r.logger.Warn("camera read failed", "error", err)
			r.profiler.FrameEnd()
			time.Sleep(33 * time.Millisecond)
			continue
		}

		processed, err := r.pipeline.Process(raw)
		if err != nil {
			r.logger.Warn("pipeline process failed", "error", err)
			r.profiler.FrameEnd()
			continue
		}

		latencyStart := time.Now()

		handData, err := r.landmark.Extract(processed)
		if err != nil {
			r.logger.Warn("landmark extraction failed", "error", err)
			r.profiler.FrameEnd()
			continue
		}

		if handData != nil {
			r.machine.Process(handData)
			if r.tracking {
				r.moveCursor(handData)
			}
		}

		r.profiler.TrackLatency(time.Since(latencyStart))

		if r.overlay != nil {
			annotated, err := r.overlay.Annotate(processed, handData)
			if err == nil {
				r.overlay.Show(annotated)
			}
		}

		r.profiler.FrameEnd()
	}
}

func (r *Runner) moveCursor(data *models.HandData) {
	if data == nil || len(data.Landmarks) < 9 {
		return
	}

	indexTip := data.Landmarks[8]

	nx := indexTip.X / 224.0
	ny := indexTip.Y / 224.0

	if nx < 0 {
		nx = 0
	} else if nx > 1 {
		nx = 1
	}
	if ny < 0 {
		ny = 0
	} else if ny > 1 {
		ny = 1
	}

	if r.display != nil {
		sx, sy := r.display.NormalizeToActive(nx, ny)
		r.ctrl.MoveMouse(sx, sy)
	} else {
		sx := int(nx * 1920)
		sy := int(ny * 1080)
		r.ctrl.MoveMouse(sx, sy)
	}
}

func (r *Runner) SetTracking(active bool) {
	r.tracking = active
	r.engine.SetTracking(active)
	if r.overlay != nil {
		r.overlay.SetTracking(active)
	}
}
