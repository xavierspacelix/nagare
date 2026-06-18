//go:build windows

package controller

import (
	"github.com/go-vgo/robotgo"
)

func New() OSController {
	return &WindowsController{}
}

type WindowsController struct{}

func (w *WindowsController) MoveMouse(x, y int) error {
	robotgo.MoveMouse(x, y)
	return nil
}

func (w *WindowsController) LeftClick() error {
	return robotgo.Click("left")
}

func (w *WindowsController) RightClick() error {
	return robotgo.Click("right")
}

func (w *WindowsController) MouseDown() error {
	return robotgo.MouseDown("left")
}

func (w *WindowsController) MouseUp() error {
	return robotgo.MouseUp("left")
}

func (w *WindowsController) Scroll(ticks int) error {
	robotgo.Scroll(0, ticks)
	return nil
}

func (w *WindowsController) VolumeUp() error {
	robotgo.KeyTap("audio_vol_up")
	return nil
}

func (w *WindowsController) VolumeDown() error {
	robotgo.KeyTap("audio_vol_down")
	return nil
}

func (w *WindowsController) Mute() error {
	robotgo.KeyTap("audio_mute")
	return nil
}

func (w *WindowsController) MediaPlayPause() error {
	robotgo.KeyTap("audio_play")
	return nil
}

func (w *WindowsController) MediaNext() error {
	robotgo.KeyTap("audio_next")
	return nil
}

func (w *WindowsController) MediaPrevious() error {
	robotgo.KeyTap("audio_prev")
	return nil
}

func (w *WindowsController) KeyTap(key string, modifiers ...string) error {
	if len(modifiers) > 0 {
		args := make([]interface{}, len(modifiers))
		for i, m := range modifiers {
			args[i] = m
		}
		return robotgo.KeyTap(key, args...)
	}
	return robotgo.KeyTap(key)
}
