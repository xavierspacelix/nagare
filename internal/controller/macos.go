//go:build darwin

package controller

import (
	"github.com/go-vgo/robotgo"
)

func New() OSController {
	return &MacController{}
}

type MacController struct{}

func (c *MacController) MoveMouse(x, y int) error {
	robotgo.MoveMouse(x, y)
	return nil
}

func (c *MacController) LeftClick() error {
	return robotgo.Click("left")
}

func (c *MacController) RightClick() error {
	return robotgo.Click("right")
}

func (c *MacController) MouseDown() error {
	return robotgo.MouseDown("left")
}

func (c *MacController) MouseUp() error {
	return robotgo.MouseUp("left")
}

func (c *MacController) Scroll(ticks int) error {
	robotgo.Scroll(0, ticks)
	return nil
}

func (c *MacController) VolumeUp() error {
	robotgo.KeyTap("audio_vol_up")
	return nil
}

func (c *MacController) VolumeDown() error {
	robotgo.KeyTap("audio_vol_down")
	return nil
}

func (c *MacController) Mute() error {
	robotgo.KeyTap("audio_mute")
	return nil
}

func (c *MacController) MediaPlayPause() error {
	robotgo.KeyTap("audio_play")
	return nil
}

func (c *MacController) MediaNext() error {
	robotgo.KeyTap("audio_next")
	return nil
}

func (c *MacController) MediaPrevious() error {
	robotgo.KeyTap("audio_prev")
	return nil
}

func (c *MacController) KeyTap(key string, modifiers ...string) error {
	if len(modifiers) > 0 {
		args := make([]interface{}, len(modifiers))
		for i, m := range modifiers {
			args[i] = m
		}
		return robotgo.KeyTap(key, args...)
	}
	return robotgo.KeyTap(key)
}
