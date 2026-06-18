package actions

import (
	"testing"

	"nagare/internal/controller"
	"nagare/internal/gestures"
	"nagare/models"
)

func TestEngine_LeftClickOnPinch(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	event := models.GestureEvent{
		Gesture:    models.GesturePinch,
		State:      models.GestureActive,
		Confidence: 0.9,
	}

	e.Handle(event)

	if stub.LastAction != "left_click" {
		t.Fatalf("expected left_click, got %s", stub.LastAction)
	}
}

func TestEngine_RightClickOnTwoFingerPinch(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	e.Handle(models.GestureEvent{
		Gesture: models.GestureTwoFingerPinch,
		State:   models.GestureActive,
	})

	if stub.LastAction != "right_click" {
		t.Fatalf("expected right_click, got %s", stub.LastAction)
	}
}

func TestEngine_MouseDownOnPinchHold(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	e.Handle(models.GestureEvent{
		Gesture: models.GesturePinchHold,
		State:   models.GestureActive,
	})

	if stub.LastAction != "mouse_down" {
		t.Fatalf("expected mouse_down, got %s", stub.LastAction)
	}
}

func TestEngine_MouseUpOnPinchHoldEnd(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	e.Handle(models.GestureEvent{
		Gesture: models.GesturePinchHold,
		State:   models.GestureEnd,
	})

	if stub.LastAction != "mouse_up" {
		t.Fatalf("expected mouse_up, got %s", stub.LastAction)
	}
}

func TestEngine_VolumeUp(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	e.Handle(models.GestureEvent{
		Gesture: models.GestureTwoFingerUp,
		State:   models.GestureActive,
	})

	if stub.LastAction != "volume_up" {
		t.Fatalf("expected volume_up, got %s", stub.LastAction)
	}
}

func TestEngine_VolumeDown(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	e.Handle(models.GestureEvent{
		Gesture: models.GestureTwoFingerDown,
		State:   models.GestureActive,
	})

	if stub.LastAction != "volume_down" {
		t.Fatalf("expected volume_down, got %s", stub.LastAction)
	}
}

func TestEngine_MediaPlayPause(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	e.Handle(models.GestureEvent{
		Gesture: models.GestureOpenPalm,
		State:   models.GestureActive,
	})

	if stub.LastAction != "" {
		t.Fatalf("expected no action for open palm (handled by engine), got %s", stub.LastAction)
	}
}

func TestEngine_TrackingOnOff(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)

	if e.IsTracking() {
		t.Fatal("expected tracking off initially")
	}

	e.Handle(models.GestureEvent{
		Gesture: models.GestureOpenPalm,
		State:   models.GestureActive,
	})

	if !e.IsTracking() {
		t.Fatal("expected tracking on after open palm")
	}

	e.Handle(models.GestureEvent{
		Gesture: models.GestureClosedFist,
		State:   models.GestureActive,
	})

	if e.IsTracking() {
		t.Fatal("expected tracking off after closed fist")
	}
}

func TestEngine_BlocksActionsWhenNotTracking(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)

	e.Handle(models.GestureEvent{
		Gesture: models.GesturePinch,
		State:   models.GestureActive,
	})

	if stub.LastAction != "" {
		t.Fatalf("expected no action when not tracking, got %s", stub.LastAction)
	}
}

func TestEngine_UnknownGesture(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	action, ok := gestures.LookupMapping(models.GesturePinch, models.GestureStart)
	if ok {
		t.Logf("start mapping: %s", action)
	}
	e.Handle(models.GestureEvent{
		Gesture: models.GesturePinch,
		State:   models.GestureStart,
	})

	if stub.LastAction != "" {
		t.Fatalf("expected no action for start state, got %s", stub.LastAction)
	}
}

func TestEngine_MediaNextPrev(t *testing.T) {
	stub := controller.NewStubController()
	e := NewEngine(stub, nil)
	e.SetTracking(true)

	e.Handle(models.GestureEvent{
		Gesture: models.GestureSwipeLeft,
		State:   models.GestureActive,
	})
	if stub.LastAction != "media_prev" {
		t.Fatalf("expected media_prev, got %s", stub.LastAction)
	}

	e.Handle(models.GestureEvent{
		Gesture: models.GestureSwipeRight,
		State:   models.GestureActive,
	})
	if stub.LastAction != "media_next" {
		t.Fatalf("expected media_next, got %s", stub.LastAction)
	}
}
