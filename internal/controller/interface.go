package controller

type OSController interface {
	MoveMouse(x, y int) error
	LeftClick() error
	RightClick() error
	MouseDown() error
	MouseUp() error
	Scroll(ticks int) error
	VolumeUp() error
	VolumeDown() error
	Mute() error
	MediaPlayPause() error
	MediaNext() error
	MediaPrevious() error
	KeyTap(key string, modifiers ...string) error
}
