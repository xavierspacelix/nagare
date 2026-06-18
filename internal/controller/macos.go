//go:build darwin

package controller

import (
	"github.com/go-vgo/robotgo"
	"nagare/internal/display"
)

func New() OSController {
	return &MacController{}
}

type MacController struct{}

func (mc *MacController) MoveMouse(x, y int) error {
	robotgo.MoveMouse(x, y)
	return nil
}

func (mc *MacController) LeftClick() error {
	return robotgo.Click("left")
}

func (mc *MacController) RightClick() error {
	return robotgo.Click("right")
}

func (mc *MacController) MouseDown() error {
	return robotgo.MouseDown("left")
}

func (mc *MacController) MouseUp() error {
	return robotgo.MouseUp("left")
}

func (mc *MacController) Scroll(ticks int) error {
	robotgo.Scroll(0, ticks)
	return nil
}

func (mc *MacController) VolumeUp() error {
	robotgo.KeyTap("audio_vol_up")
	return nil
}

func (mc *MacController) VolumeDown() error {
	robotgo.KeyTap("audio_vol_down")
	return nil
}

func (mc *MacController) Mute() error {
	robotgo.KeyTap("audio_mute")
	return nil
}

func (mc *MacController) MediaPlayPause() error {
	robotgo.KeyTap("audio_play")
	return nil
}

func (mc *MacController) MediaNext() error {
	robotgo.KeyTap("audio_next")
	return nil
}

func (mc *MacController) MediaPrevious() error {
	robotgo.KeyTap("audio_prev")
	return nil
}

func (mc *MacController) KeyTap(key string, modifiers ...string) error {
	if len(modifiers) > 0 {
		args := make([]interface{}, len(modifiers))
		for i, m := range modifiers {
			args[i] = m
		}
		return robotgo.KeyTap(key, args...)
	}
	return robotgo.KeyTap(key)
}

func (mc *MacController) GetMonitors() ([]display.Info, error) {
	num := robotgo.DisplaysNum()
	monitors := make([]display.Info, 0, num)

	for i := range num {
		x, y, monW, monH := robotgo.GetDisplayBounds(i)
		monitors = append(monitors, display.Info{
			Index:   i,
			X:       x,
			Y:       y,
			Width:   monW,
			Height:  monH,
			Primary: i == 0,
		})
	}

	if len(monitors) == 0 {
		sw, sh := robotgo.GetScreenSize()
		monitors = append(monitors, display.Info{
			Index: 0, X: 0, Y: 0,
			Width: sw, Height: sh, Primary: true,
		})
	}

	return monitors, nil
}
