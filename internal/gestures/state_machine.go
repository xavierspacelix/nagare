package gestures

import (
	"log/slog"
	"time"

	"nagare/models"
)

type EventHandler func(event models.GestureEvent)

type Config struct {
	OpenPalmThreshold    time.Duration
	PinchThreshold       time.Duration
	StabilizerFrames     int
	DefaultMinConfidence float64
	GestureMinConfidence map[models.Gesture]float64
}

func DefaultConfig() Config {
	return Config{
		OpenPalmThreshold:    300 * time.Millisecond,
		PinchThreshold:       100 * time.Millisecond,
		StabilizerFrames:     3,
		DefaultMinConfidence: 0.5,
		GestureMinConfidence: nil,
	}
}

type Machine struct {
	config  Config
	onEvent EventHandler
	states  map[models.Gesture]*gestureState
	reco    *Recognizer
	cm      *CooldownManager
	stab    *Stabilizer
	clock   func() time.Time
	logger  *slog.Logger
}

type gestureState struct {
	current  models.GestureState
	since    time.Time
}

func NewMachine(cfg Config, handler EventHandler, logger *slog.Logger) *Machine {
	if logger == nil {
		logger = slog.Default()
	}

	m := &Machine{
		config:  cfg,
		onEvent: handler,
		states:  make(map[models.Gesture]*gestureState),
		reco:    NewRecognizer(),
		cm:      NewCooldownManager(250 * time.Millisecond),
		stab:    NewStabilizer(cfg.StabilizerFrames),
		clock:   time.Now,
		logger:  logger,
	}

	for _, g := range allGestures() {
		m.states[g] = &gestureState{current: models.GestureIdle}
	}

	return m
}

func allGestures() []models.Gesture {
	return []models.Gesture{
		models.GestureOpenPalm,
		models.GestureClosedFist,
		models.GesturePinch,
		models.GesturePinchHold,
		models.GestureTwoFingerPinch,
		models.GestureTwoFingerUp,
		models.GestureTwoFingerDown,
		models.GestureSwipeLeft,
		models.GestureSwipeRight,
		models.GestureScrollUp,
		models.GestureScrollDown,
	}
}

func (m *Machine) Process(data *models.HandData) {
	now := m.clock()

	m.processGesture(now, data, models.GestureOpenPalm, m.reco.IsOpenPalm)
	m.processGesture(now, data, models.GestureClosedFist, m.reco.IsClosedFist)
	m.processGesture(now, data, models.GesturePinch, m.reco.IsPinch)
	m.processGesture(now, data, models.GesturePinchHold, m.reco.IsPinch)
	m.processGesture(now, data, models.GestureTwoFingerPinch, m.reco.IsTwoFingerPinch)
	m.processGesture(now, data, models.GestureTwoFingerUp, m.reco.IsTwoFingersUp)
	m.processGesture(now, data, models.GestureTwoFingerDown, m.reco.IsTwoFingersDown)
	m.processGesture(now, data, models.GestureSwipeLeft, m.reco.IsSwipeLeft)
	m.processGesture(now, data, models.GestureSwipeRight, m.reco.IsSwipeRight)
	m.processGesture(now, data, models.GestureScrollUp, m.reco.IsScrollUp)
	m.processGesture(now, data, models.GestureScrollDown, m.reco.IsScrollDown)
}

func (m *Machine) confidenceFor(gesture models.Gesture) float64 {
	if m.config.GestureMinConfidence != nil {
		if c, ok := m.config.GestureMinConfidence[gesture]; ok {
			return c
		}
	}
	return m.config.DefaultMinConfidence
}

func (m *Machine) processGesture(now time.Time, data *models.HandData, gesture models.Gesture, check func(*models.HandData) bool) {
	state := m.states[gesture]

	if data == nil || data.Confidence < m.confidenceFor(gesture) {
		if state.current != models.GestureIdle {
			state.current = models.GestureEnd
			state.since = now
			m.stab.Reset(gesture)
			m.stab.Record(gesture, false)
			m.emit(gesture, models.GestureEnd, data)
		}
		return
	}

	raw := check(data)
	stable := m.stab.Record(gesture, raw)

	switch state.current {
	case models.GestureIdle:
		if !m.cm.IsReady(gesture) {
			return
		}
		if stable {
			state.current = models.GestureStart
			state.since = now
			m.emit(gesture, models.GestureStart, data)
		}

	case models.GestureStart:
		if !stable {
			state.current = models.GestureIdle
			m.emit(gesture, models.GestureEnd, data)
			break
		}

		threshold := m.config.PinchThreshold
		if gesture == models.GestureOpenPalm {
			threshold = m.config.OpenPalmThreshold
		} else if gesture == models.GesturePinchHold {
			threshold = m.config.PinchThreshold
		}

		if now.Sub(state.since) >= threshold {
			state.current = models.GestureActive
			m.cm.Start(gesture)
			m.emit(gesture, models.GestureActive, data)

			if gesture == models.GestureSwipeLeft || gesture == models.GestureSwipeRight {
				m.reco.ResetSwipe()
			}
		}

	case models.GestureActive:
		if !stable {
			state.current = models.GestureEnd
			state.since = now
			m.cm.Start(gesture)
			m.emit(gesture, models.GestureEnd, data)

			if gesture == models.GesturePinchHold {
				m.cm.Start(gesture)
			}
		}

	case models.GestureEnd:
		state.current = models.GestureIdle
	}
}

func (m *Machine) emit(gesture models.Gesture, state models.GestureState, data *models.HandData) {
	confidence := 0.0
	if data != nil {
		confidence = data.Confidence
	}

	event := models.GestureEvent{
		Gesture:    gesture,
		State:      state,
		Confidence: confidence,
	}

	if m.onEvent != nil {
		m.onEvent(event)
	}
}

func (m *Machine) State(gesture models.Gesture) models.GestureState {
	s, ok := m.states[gesture]
	if !ok {
		return models.GestureIdle
	}
	return s.current
}

func (m *Machine) Recognizer() *Recognizer {
	return m.reco
}

func (m *Machine) Stabilizer() *Stabilizer {
	return m.stab
}
