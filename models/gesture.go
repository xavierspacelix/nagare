package models

type Gesture string

const (
	GestureOpenPalm   Gesture = "open_palm"
	GestureClosedFist Gesture = "closed_fist"
	GesturePinch      Gesture = "pinch"
	GesturePinchHold  Gesture = "pinch_hold"
	GestureTwoFingerPinch Gesture = "two_finger_pinch"
	GestureTwoFingerUp    Gesture = "two_fingers_up"
	GestureTwoFingerDown  Gesture = "two_fingers_down"
	GestureSwipeLeft  Gesture = "swipe_left"
	GestureSwipeRight Gesture = "swipe_right"
)

type GestureState int

const (
	GestureIdle   GestureState = 0
	GestureStart  GestureState = 1
	GestureActive GestureState = 2
	GestureEnd    GestureState = 3
)

type GestureEvent struct {
	Gesture    Gesture
	State      GestureState
	Confidence float64
}
