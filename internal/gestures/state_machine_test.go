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

func stabilize(m *Machine, data *models.HandData) {
	for range 3 {
		m.Process(data)
	}
}

func TestOpenPalm_IdleToStart(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	stabilize(m, newHandData(allExtended(), false))

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

	stabilize(m, data)

	m.clock = func() time.Time { return now.Add(400 * time.Millisecond) }
	stabilize(m, data)

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

	stabilize(m, newHandData(allExtended(), false))

	now2 := time.Now()
	m.clock = func() time.Time { return now2 }
	stabilize(m, newHandData(allFolded(), false))

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

	stabilize(m, data)

	m.clock = func() time.Time { return now.Add(200 * time.Millisecond) }
	stabilize(m, data)

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

	stabilize(m, data)
	m.clock = func() time.Time { return now.Add(200 * time.Millisecond) }
	stabilize(m, data)

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

func TestStabilizer_RequiresMultipleFrames(t *testing.T) {
	s := NewStabilizer(3)

	stable := s.Record(models.GestureOpenPalm, true)
	if stable {
		t.Fatal("expected not stable after 1 frame")
	}

	stable = s.Record(models.GestureOpenPalm, true)
	if stable {
		t.Fatal("expected not stable after 2 frames")
	}

	stable = s.Record(models.GestureOpenPalm, true)
	if !stable {
		t.Fatal("expected stable after 3 frames")
	}
}

func TestStabilizer_NoiseFiltering(t *testing.T) {
	s := NewStabilizer(3)

	s.Record(models.GestureOpenPalm, true)
	s.Record(models.GestureOpenPalm, true)
	stable := s.Record(models.GestureOpenPalm, false)
	if !stable {
		t.Fatal("expected stable with majority true")
	}

	s2 := NewStabilizer(3)
	s2.Record(models.GestureOpenPalm, true)
	s2.Record(models.GestureOpenPalm, false)
	stable = s2.Record(models.GestureOpenPalm, false)
	if stable {
		t.Fatal("expected not stable with minority true")
	}
}

func TestStabilizer_Reset(t *testing.T) {
	s := NewStabilizer(3)

	s.Record(models.GestureOpenPalm, true)
	s.Record(models.GestureOpenPalm, true)
	s.Record(models.GestureOpenPalm, true)
	s.Reset(models.GestureOpenPalm)

	stable := s.Record(models.GestureOpenPalm, true)
	if stable {
		t.Fatal("expected not stable after reset")
	}
}

func TestStabilizer_ResetAll(t *testing.T) {
	s := NewStabilizer(3)

	s.Record(models.GestureOpenPalm, true)
	s.Record(models.GestureOpenPalm, true)
	s.Record(models.GestureOpenPalm, true)
	s.ResetAll()

	stable := s.Record(models.GestureOpenPalm, true)
	if stable {
		t.Fatal("expected not stable after full reset")
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

func TestConfidenceGate_BlocksLowConfidence(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	data := newHandData(allExtended(), false)
	data.Confidence = 0.3

	stabilize(m, data)

	for _, e := range events {
		if e.Gesture == models.GestureOpenPalm {
			t.Fatal("expected no open palm events with low confidence")
		}
	}
}

func TestConfidenceGate_InterruptsActiveOnLowConfidence(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	now := time.Now()
	m.clock = func() time.Time { return now }

	data := newHandData(allExtended(), false)
	stabilize(m, data)

	m.clock = func() time.Time { return now.Add(400 * time.Millisecond) }
	stabilize(m, data)

	activeFound := false
	for _, e := range events {
		if e.Gesture == models.GestureOpenPalm && e.State == models.GestureActive {
			activeFound = true
		}
	}
	if !activeFound {
		t.Fatal("expected active before confidence drop")
	}

	data.Confidence = 0.3
	m.clock = func() time.Time { return now.Add(500 * time.Millisecond) }
	stabilize(m, data)

	endFound := false
	for _, e := range events {
		if e.Gesture == models.GestureOpenPalm && e.State == models.GestureEnd {
			endFound = true
		}
	}
	if !endFound {
		t.Fatal("expected end event after confidence drop")
	}
}

func TestConfidenceGate_PerGestureThreshold(t *testing.T) {
	cfg := DefaultConfig()
	cfg.GestureMinConfidence = map[models.Gesture]float64{
		models.GesturePinch: 0.8,
	}
	events := []models.GestureEvent{}
	m := NewMachine(cfg, func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	now := time.Now()
	m.clock = func() time.Time { return now }

	data := newHandData(allExtended(), true)
	data.Confidence = 0.6

	stabilize(m, data)
	m.clock = func() time.Time { return now.Add(200 * time.Millisecond) }
	stabilize(m, data)

	for _, e := range events {
		if e.Gesture == models.GesturePinch {
			t.Fatal("expected no pinch with confidence below per-gesture threshold")
		}
	}
}

func TestDebounce_PreventsRapidRetrigger(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	now := time.Now()
	m.clock = func() time.Time { return now }

	data := newHandData(indexMiddleExtended(), false)
	closed := newHandData(allFolded(), false)

	stabilize(m, data)
	m.clock = func() time.Time { return now.Add(200 * time.Millisecond) }
	stabilize(m, data)

	m.clock = func() time.Time { return now.Add(300 * time.Millisecond) }
	stabilize(m, closed)

	m.clock = func() time.Time { return now.Add(400 * time.Millisecond) }
	stabilize(m, data)

	startCount := 0
	for _, e := range events {
		if e.Gesture == models.GestureTwoFingerUp && e.State == models.GestureStart {
			startCount++
		}
	}
	if startCount > 1 {
		t.Fatalf("expected at most 1 start event due to cooldown, got %d", startCount)
	}
}

func TestCooldown_DoesNotInterruptActive(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	now := time.Now()
	m.clock = func() time.Time { return now }

	data := newHandData(allExtended(), false)
	stabilize(m, data)
	m.clock = func() time.Time { return now.Add(400 * time.Millisecond) }
	stabilize(m, data)

	m.clock = func() time.Time { return now.Add(500 * time.Millisecond) }
	stabilize(m, data)

	m.clock = func() time.Time { return now.Add(600 * time.Millisecond) }
	stabilize(m, data)

	activeCount := 0
	for _, e := range events {
		if e.Gesture == models.GestureOpenPalm && e.State == models.GestureActive {
			activeCount++
		}
	}
	if activeCount != 1 {
		t.Fatalf("expected exactly 1 active event, got %d (cooldown should not interrupt active)", activeCount)
	}
}

func TestMappingStore_FallbackToDefault(t *testing.T) {
	ms := NewMappingStore()

	action, ok := ms.Lookup(models.GesturePinch, models.GestureActive)
	if !ok {
		t.Fatal("expected default mapping found")
	}
	if action != ActionLeftClick {
		t.Fatalf("expected left_click, got %s", action)
	}
}

func TestMappingStore_CustomOverridesDefault(t *testing.T) {
	ms := NewMappingStore()
	ms.SetCustom([]Mapping{
		{Gesture: models.GesturePinch, Action: ActionRightClick, OnState: models.GestureActive},
	})

	action, ok := ms.Lookup(models.GesturePinch, models.GestureActive)
	if !ok {
		t.Fatal("expected custom mapping found")
	}
	if action != ActionRightClick {
		t.Fatalf("expected right_click, got %s", action)
	}
}

func TestMappingStore_CustomNotFound(t *testing.T) {
	ms := NewMappingStore()
	ms.SetCustom([]Mapping{
		{Gesture: models.GesturePinch, Action: ActionLeftClick, OnState: models.GestureActive},
	})

	_, ok := ms.Lookup(models.GestureSwipeLeft, models.GestureActive)
	if !ok {
		t.Fatal("expected fallback to default")
	}
}

func TestMappingStore_GetCustom(t *testing.T) {
	ms := NewMappingStore()
	if ms.GetCustom() != nil {
		t.Fatal("expected nil initially")
	}

	custom := []Mapping{
		{Gesture: models.GesturePinch, Action: ActionLeftClick, OnState: models.GestureActive},
	}
	ms.SetCustom(custom)
	if len(ms.GetCustom()) != 1 {
		t.Fatal("expected 1 custom mapping")
	}
}

func TestGestureFromName(t *testing.T) {
	g, ok := GestureFromName("pinch")
	if !ok {
		t.Fatal("expected gesture found")
	}
	if g != models.GesturePinch {
		t.Fatalf("expected pinch, got %s", g)
	}

	_, ok = GestureFromName("nonexistent")
	if ok {
		t.Fatal("expected not found")
	}
}

func TestActionFromName(t *testing.T) {
	a, ok := ActionFromName("left_click")
	if !ok {
		t.Fatal("expected action found")
	}
	if a != ActionLeftClick {
		t.Fatalf("expected left_click, got %s", a)
	}

	_, ok = ActionFromName("nonexistent")
	if ok {
		t.Fatal("expected not found")
	}
}

func TestMachine_EventHandler(t *testing.T) {
	events := []models.GestureEvent{}
	m := NewMachine(DefaultConfig(), func(e models.GestureEvent) {
		events = append(events, e)
	}, nil)

	stabilize(m, newHandData(allExtended(), false))

	if len(events) == 0 {
		t.Fatal("expected at least one event")
	}
}
