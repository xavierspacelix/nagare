package gestures

import "nagare/models"

type Action string

const (
	ActionNone          Action = ""
	ActionTrackingOn    Action = "tracking_on"
	ActionTrackingOff   Action = "tracking_off"
	ActionLeftClick     Action = "left_click"
	ActionRightClick    Action = "right_click"
	ActionMouseDown     Action = "mouse_down"
	ActionMouseUp       Action = "mouse_up"
	ActionScrollUp      Action = "scroll_up"
	ActionScrollDown    Action = "scroll_down"
	ActionVolumeUp      Action = "volume_up"
	ActionVolumeDown    Action = "volume_down"
	ActionMute          Action = "mute"
	ActionKeyTap        Action = "key_tap"
	ActionMediaPlayPause Action = "media_play_pause"
	ActionMediaNext     Action = "media_next"
	ActionMediaPrev     Action = "media_prev"
)

type Mapping struct {
	Gesture models.Gesture
	Action  Action
	OnState models.GestureState
}

var DefaultMappings = []Mapping{
	{Gesture: models.GestureOpenPalm, Action: ActionTrackingOn, OnState: models.GestureActive},
	{Gesture: models.GestureClosedFist, Action: ActionTrackingOff, OnState: models.GestureActive},
	{Gesture: models.GesturePinch, Action: ActionLeftClick, OnState: models.GestureActive},
	{Gesture: models.GesturePinchHold, Action: ActionMouseDown, OnState: models.GestureActive},
	{Gesture: models.GesturePinchHold, Action: ActionMouseUp, OnState: models.GestureEnd},
	{Gesture: models.GestureTwoFingerPinch, Action: ActionRightClick, OnState: models.GestureActive},
	{Gesture: models.GestureTwoFingerUp, Action: ActionVolumeUp, OnState: models.GestureActive},
	{Gesture: models.GestureTwoFingerDown, Action: ActionVolumeDown, OnState: models.GestureActive},
	{Gesture: models.GestureSwipeLeft, Action: ActionMediaPrev, OnState: models.GestureActive},
	{Gesture: models.GestureSwipeRight, Action: ActionMediaNext, OnState: models.GestureActive},
	{Gesture: models.GestureScrollUp, Action: ActionScrollUp, OnState: models.GestureActive},
	{Gesture: models.GestureScrollDown, Action: ActionScrollDown, OnState: models.GestureActive},
}

func LookupMapping(gesture models.Gesture, state models.GestureState) (Action, bool) {
	for _, m := range DefaultMappings {
		if m.Gesture == gesture && m.OnState == state {
			return m.Action, true
		}
	}
	return ActionNone, false
}

type MappingStore struct {
	custom []Mapping
}

func NewMappingStore() *MappingStore {
	return &MappingStore{}
}

func (ms *MappingStore) SetCustom(mappings []Mapping) {
	ms.custom = mappings
}

func (ms *MappingStore) GetCustom() []Mapping {
	if ms.custom == nil {
		return nil
	}
	return ms.custom
}

func (ms *MappingStore) Lookup(gesture models.Gesture, state models.GestureState) (Action, bool) {
	for _, m := range ms.custom {
		if m.Gesture == gesture && m.OnState == state {
			return m.Action, true
		}
	}
	return LookupMapping(gesture, state)
}

func GestureFromName(name string) (models.Gesture, bool) {
	for _, g := range allGestures() {
		if string(g) == name {
			return g, true
		}
	}
	return "", false
}

func ActionFromName(name string) (Action, bool) {
	for _, a := range AllActions() {
		if string(a) == name {
			return a, true
		}
	}
	return ActionNone, false
}

func AllActions() []Action {
	return []Action{
		ActionTrackingOn, ActionTrackingOff,
		ActionLeftClick, ActionRightClick,
		ActionMouseDown, ActionMouseUp,
		ActionScrollUp, ActionScrollDown,
		ActionVolumeUp, ActionVolumeDown, ActionMute,
		ActionMediaPlayPause, ActionMediaNext, ActionMediaPrev, ActionKeyTap,
	}
}
