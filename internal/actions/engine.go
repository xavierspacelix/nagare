package actions

import (
	"log/slog"

	"nagare/internal/controller"
	"nagare/internal/gestures"
	"nagare/models"
)

type Engine struct {
	ctrl   controller.OSController
	logger *slog.Logger
	tracking bool
}

func NewEngine(ctrl controller.OSController, logger *slog.Logger) *Engine {
	if logger == nil {
		logger = slog.Default()
	}
	return &Engine{
		ctrl:   ctrl,
		logger: logger,
	}
}

func (e *Engine) Handle(event models.GestureEvent) {
	if !e.tracking && event.Gesture != models.GestureOpenPalm {
		return
	}

	action, ok := gestures.LookupMapping(event.Gesture, event.State)
	if !ok {
		return
	}

	e.logger.Debug("executing action",
		"gesture", event.Gesture,
		"state", event.State,
		"action", action,
		"confidence", event.Confidence,
	)

	var err error
	switch action {
	case gestures.ActionTrackingOn:
		e.tracking = true
		e.logger.Info("tracking enabled")
		return
	case gestures.ActionTrackingOff:
		e.tracking = false
		e.logger.Info("tracking disabled")
		return
	case gestures.ActionLeftClick:
		err = e.ctrl.LeftClick()
	case gestures.ActionRightClick:
		err = e.ctrl.RightClick()
	case gestures.ActionMouseDown:
		err = e.ctrl.MouseDown()
	case gestures.ActionMouseUp:
		err = e.ctrl.MouseUp()
	case gestures.ActionScrollUp:
		err = e.ctrl.Scroll(1)
	case gestures.ActionScrollDown:
		err = e.ctrl.Scroll(-1)
	case gestures.ActionVolumeUp:
		err = e.ctrl.VolumeUp()
	case gestures.ActionVolumeDown:
		err = e.ctrl.VolumeDown()
	case gestures.ActionMediaPlayPause:
		err = e.ctrl.MediaPlayPause()
	case gestures.ActionMediaNext:
		err = e.ctrl.MediaNext()
	case gestures.ActionMediaPrev:
		err = e.ctrl.MediaPrevious()
	default:
		e.logger.Warn("unknown action", "action", action)
		return
	}

	if err != nil {
		e.logger.Error("action failed",
			"action", action,
			"error", err,
		)
	}
}

func (e *Engine) SetTracking(active bool) {
	e.tracking = active
}

func (e *Engine) IsTracking() bool {
	return e.tracking
}
