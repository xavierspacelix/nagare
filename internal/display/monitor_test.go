package display

import (
	"testing"
)

func TestNewMapper(t *testing.T) {
	d := NewDiscoverer()
	m := NewMapper(d)
	if m == nil {
		t.Fatal("expected non-nil mapper")
	}
}

func TestMapper_Refresh(t *testing.T) {
	d := NewDiscoverer()
	m := NewMapper(d)

	if err := m.Refresh(); err != nil {
		t.Fatalf("unexpected refresh error: %v", err)
	}

	monitors := m.Monitors()
	if len(monitors) == 0 {
		t.Fatal("expected at least 1 monitor")
	}
}

func TestMapper_ActiveMonitor(t *testing.T) {
	d := NewDiscoverer()
	m := NewMapper(d)
	m.Refresh()

	mon := m.ActiveMonitor()
	if mon.Width == 0 || mon.Height == 0 {
		t.Fatal("expected non-zero monitor dimensions")
	}
}

func TestMapper_SetActiveMonitor(t *testing.T) {
	d := NewDiscoverer()
	m := NewMapper(d)
	m.Refresh()

	m.SetActiveMonitor(-1)
	mon := m.ActiveMonitor()
	if mon.Index != 0 {
		t.Fatal("expected fallback to first monitor")
	}
}

func TestMapper_NormalizeToActive(t *testing.T) {
	d := NewDiscoverer()
	m := NewMapper(d)
	m.Refresh()

	x, y := m.NormalizeToActive(0.5, 0.5)
	if x == 0 && y == 0 {
		t.Fatal("expected non-zero screen position")
	}
}

func TestMapper_TotalWidth(t *testing.T) {
	d := NewDiscoverer()
	m := NewMapper(d)
	m.Refresh()

	if m.TotalWidth() <= 0 {
		t.Fatal("expected positive total width")
	}
}

func TestMapper_TotalHeight(t *testing.T) {
	d := NewDiscoverer()
	m := NewMapper(d)
	m.Refresh()

	if m.TotalHeight() <= 0 {
		t.Fatal("expected positive total height")
	}
}
