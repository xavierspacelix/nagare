package gestures

import (
	"testing"
	"time"

	"nagare/models"
)

func newHandData(fingers models.FingerStates, pinch bool) *models.HandData {
	var landmarks [21]models.HandLandmark
	for i := range 21 {
		landmarks[i] = models.HandLandmark{X: 100, Y: float64(200 - i*5), Z: 0}
	}
	if pinch {
		landmarks[4] = models.HandLandmark{X: 100, Y: 100, Z: 0}
		landmarks[8] = models.HandLandmark{X: 100.02, Y: 100, Z: 0}
	} else {
		landmarks[4] = models.HandLandmark{X: 90, Y: 50, Z: 0}
		landmarks[8] = models.HandLandmark{X: 130, Y: 40, Z: 0}
	}

	return &models.HandData{
		Landmarks:   landmarks,
		Confidence:  0.95,
		Orientation: models.HandOrientationRight,
		Fingers:     fingers,
	}
}

func allExtended() models.FingerStates {
	return models.FingerStates{
		models.FingerExtended,
		models.FingerExtended,
		models.FingerExtended,
		models.FingerExtended,
		models.FingerExtended,
	}
}

func allFolded() models.FingerStates {
	return models.FingerStates{
		models.FingerFolded,
		models.FingerFolded,
		models.FingerFolded,
		models.FingerFolded,
		models.FingerFolded,
	}
}

func indexMiddleExtended() models.FingerStates {
	return models.FingerStates{
		models.FingerFolded,
		models.FingerExtended,
		models.FingerExtended,
		models.FingerFolded,
		models.FingerFolded,
	}
}

func TestOpenPalm_IdleToStart(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	m.Process(newHandData(allExtended(), false))

	if len(events) < 1 {
		t.Fatal("expected at least one event")
	}
	if events[0].Gesture != models.GestureOpenPalm {
		t.Fatalf("expected open_palm, got %s", events[0].Gesture)
	}
	if events[0].State != models.GestureStart {
		t.Fatalf("expected GestureStart, got %v", events[0].State)
	}
}

func TestOpenPalm_StartToActive(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	now := time.Now()
	m.clock = func() time.Time { return now }

	data := newHandData(allExtended(), false)

	m.Process(data)
	m.clock = func() time.Time { return now.Add(400 * time.Millisecond) }
	m.Process(data)

	foundActive := false
	for _, e := range events {
		if e.Gesture == models.GestureOpenPalm && e.State == models.GestureActive {
			foundActive = true
			break
		}
	}
	if !foundActive {
		t.Fatal("expected GestureActive for open palm after threshold")
	}
}

func TestOpenPalm_ReleaseBeforeThreshold(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	now := time.Now()
	m.clock = func() time.Time { return now }

	m.Process(newHandData(allExtended(), false))
	m.Process(newHandData(allFolded(), false))

	foundEnd := false
	for _, e := range events {
		if e.Gesture == models.GestureOpenPalm && e.State == models.GestureEnd {
			foundEnd = true
			break
		}
	}
	if !foundEnd {
		t.Fatal("expected GestureEnd for open palm on early release")
	}
}

func TestPinch_FullCycle(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	now := time.Now()
	m.clock = func() time.Time { return now }

	data := newHandData(allExtended(), true)

	m.Process(data)

	m.clock = func() time.Time { return now.Add(200 * time.Millisecond) }
	m.Process(data)

	foundActive := false
	for _, e := range events {
		if e.Gesture == models.GesturePinch && e.State == models.GestureActive {
			foundActive = true
			break
		}
	}
	if !foundActive {
		t.Fatal("expected pinch active event")
	}
}

func TestCooldown_BlocksRapidTrigger(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	now := time.Now()
	m.clock = func() time.Time { return now }

	data := newHandData(indexMiddleExtended(), false)

	m.Process(data)
	m.clock = func() time.Time { return now.Add(200 * time.Millisecond) }
	m.Process(data)

	activeCount := 0
	for _, e := range events {
		if e.Gesture == models.GestureTwoFingerUp && e.State == models.GestureActive {
			activeCount++
		}
	}

	if activeCount > 1 {
		t.Fatalf("expected at most 1 active event due to cooldown, got %d", activeCount)
	}
}

func TestGestureState_DefaultIdle(t *testing.T) {
	m := NewMachine(DefaultConfig(), nil, nil)

	for _, g := range allGestures() {
		if m.State(g) != models.GestureIdle {
			t.Fatalf("expected idle state for %s", g)
		}
	}
}

func TestRecognizer_IsOpenPalm(t *testing.T) {
	r := NewRecognizer()
	data := newHandData(allExtended(), false)

	if !r.IsOpenPalm(data) {
		t.Fatal("expected open palm")
	}
}

func TestRecognizer_IsClosedFist(t *testing.T) {
	r := NewRecognizer()
	data := newHandData(allFolded(), false)

	if !r.IsClosedFist(data) {
		t.Fatal("expected closed fist")
	}
}

func TestRecognizer_IsTwoFingersUp(t *testing.T) {
	r := NewRecognizer()
	data := newHandData(indexMiddleExtended(), false)

	if !r.IsTwoFingersUp(data) {
		t.Fatal("expected two fingers up")
	}
}

func TestRecognizer_LowConfidence(t *testing.T) {
	r := NewRecognizer()
	data := newHandData(allExtended(), false)
	data.Confidence = 0.3

	if r.IsOpenPalm(data) {
		t.Fatal("should not detect with low confidence")
	}
}

func TestCooldownManager_IsReady(t *testing.T) {
	now := time.Now()
	cm := NewCooldownManager(250 * time.Millisecond)
	cm.clock = func() time.Time { return now }

	if !cm.IsReady(models.GesturePinch) {
		t.Fatal("expected ready initially")
	}

	cm.Start(models.GesturePinch)
	if cm.IsReady(models.GesturePinch) {
		t.Fatal("expected not ready after start")
	}

	cm.clock = func() time.Time { return now.Add(300 * time.Millisecond) }
	if !cm.IsReady(models.GesturePinch) {
		t.Fatal("expected ready after cooldown expires")
	}
}

func TestCooldownManager_Remaining(t *testing.T) {
	now := time.Now()
	cm := NewCooldownManager(250 * time.Millisecond)
	cm.clock = func() time.Time { return now }

	cm.Start(models.GesturePinch)
	remaining := cm.Remaining(models.GesturePinch)
	if remaining <= 0 {
		t.Fatal("expected positive remaining time")
	}
}

func TestDefaultMapping_Lookup(t *testing.T) {
	action, ok := LookupMapping(models.GesturePinch, models.GestureActive)
	if !ok {
		t.Fatal("expected mapping found")
	}
	if action != ActionLeftClick {
		t.Fatalf("expected left_click, got %s", action)
	}

	action, ok = LookupMapping(models.GesturePinchHold, models.GestureEnd)
	if !ok {
		t.Fatal("expected mapping found")
	}
	if action != ActionMouseUp {
		t.Fatalf("expected mouse_up, got %s", action)
	}
}

func TestDefaultMapping_NotFound(t *testing.T) {
	_, ok := LookupMapping(models.GesturePinch, models.GestureStart)
	if ok {
		t.Fatal("expected no mapping for start state")
	}
}

func TestMachine_EventHandler(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	m.Process(newHandData(allExtended(), false))

	if len(events) == 0 {
		t.Fatal("expected at least one event")
	}
}
