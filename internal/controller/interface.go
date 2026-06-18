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
	MediaPlayPause() error
	MediaNext() error
	MediaPrevious() error
}
