package display

import "fmt"

type Info struct {
	Index   int
	X, Y    int
	Width   int
	Height  int
	Primary bool
}

type Discoverer interface {
	Discover() ([]Info, error)
}

type Mapper struct {
	discoverer Discoverer
	monitors   []Info
	activeIdx  int
}

func NewMapper(d Discoverer) *Mapper {
	return &Mapper{discoverer: d}
}

func (m *Mapper) Refresh() error {
	monitors, err := m.discoverer.Discover()
	if err != nil {
		return err
	}
	if len(monitors) == 0 {
		return fmt.Errorf("no monitors found")
	}
	m.monitors = monitors
	return nil
}

func (m *Mapper) Monitors() []Info {
	return m.monitors
}

func (m *Mapper) ActiveMonitor() Info {
	if m.activeIdx < len(m.monitors) {
		return m.monitors[m.activeIdx]
	}
	if len(m.monitors) > 0 {
		return m.monitors[0]
	}
	return Info{}
}

func (m *Mapper) SetActiveMonitor(index int) {
	for i, mon := range m.monitors {
		if mon.Index == index {
			m.activeIdx = i
			return
		}
	}
}

func (m *Mapper) TotalWidth() int {
	if len(m.monitors) == 0 {
		return 0
	}
	maxX := 0
	for _, mon := range m.monitors {
		right := mon.X + mon.Width
		if right > maxX {
			maxX = right
		}
	}
	return maxX
}

func (m *Mapper) TotalHeight() int {
	if len(m.monitors) == 0 {
		return 0
	}
	maxY := 0
	for _, mon := range m.monitors {
		bottom := mon.Y + mon.Height
		if bottom > maxY {
			maxY = bottom
		}
	}
	return maxY
}

func (m *Mapper) NormalizeToActive(nx, ny float64) (int, int) {
	mon := m.ActiveMonitor()
	screenX := mon.X + int(nx*float64(mon.Width))
	screenY := mon.Y + int(ny*float64(mon.Height))
	return screenX, screenY
}
