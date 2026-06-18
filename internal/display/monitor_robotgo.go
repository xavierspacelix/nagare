//go:build windows || darwin

package display

import "github.com/go-vgo/robotgo"

type RobotGoDiscoverer struct{}

func NewDiscoverer() Discoverer {
	return &RobotGoDiscoverer{}
}

func (d *RobotGoDiscoverer) Discover() ([]Info, error) {
	num := robotgo.DisplaysNum()
	monitors := make([]Info, 0, num)

	for i := range num {
		x, y, w, h := robotgo.GetDisplayBounds(i)
		monitors = append(monitors, Info{
			Index:   i,
			X:       x,
			Y:       y,
			Width:   w,
			Height:  h,
			Primary: i == 0,
		})
	}

	if len(monitors) == 0 {
		w, h := robotgo.GetScreenSize()
		monitors = append(monitors, Info{
			Index: 0, X: 0, Y: 0,
			Width: w, Height: h, Primary: true,
		})
	}

	return monitors, nil
}
