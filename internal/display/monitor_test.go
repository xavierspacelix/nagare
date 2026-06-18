package display

import (
	"testing"
)

func testMonitors() []Info {
	return []Info{
		{Index: 0, X: 0, Y: 0, Width: 1920, Height: 1080, Primary: true},
	}
}

func TestNewMapper(t *testing.T) {
	m, err := NewMapper(testMonitors())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if m == nil {
		t.Fatal("expected non-nil mapper")
	}
}

func TestMapper_Monitors(t *testing.T) {
	m, _ := NewMapper(testMonitors())

	monitors := m.Monitors()
	if len(monitors) == 0 {
		t.Fatal("expected at least 1 monitor")
	}
}

func TestMapper_ActiveMonitor(t *testing.T) {
	m, _ := NewMapper(testMonitors())

	mon := m.ActiveMonitor()
	if mon.Width == 0 || mon.Height == 0 {
		t.Fatal("expected non-zero monitor dimensions")
	}
}

func TestMapper_SetActiveMonitor(t *testing.T) {
	m, _ := NewMapper(testMonitors())

	m.SetActiveMonitor(-1)
	mon := m.ActiveMonitor()
	if mon.Index != 0 {
		t.Fatal("expected fallback to first monitor")
	}
}

func TestMapper_NormalizeToActive(t *testing.T) {
	m, _ := NewMapper(testMonitors())

	x, y := m.NormalizeToActive(0.5, 0.5)
	if x == 0 && y == 0 {
		t.Fatal("expected non-zero screen position")
	}
	if x != 960 || y != 540 {
		t.Fatalf("expected (960, 540), got (%d, %d)", x, y)
	}
}

func TestMapper_TotalWidth(t *testing.T) {
	m, _ := NewMapper(testMonitors())

	if m.TotalWidth() != 1920 {
		t.Fatalf("expected 1920, got %d", m.TotalWidth())
	}
}

func TestMapper_TotalHeight(t *testing.T) {
	m, _ := NewMapper(testMonitors())

	if m.TotalHeight() != 1080 {
		t.Fatalf("expected 1080, got %d", m.TotalHeight())
	}
}

func TestMapper_DefaultFallback(t *testing.T) {
	m, err := NewMapper(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if m.ActiveMonitor().Width != 1920 {
		t.Fatal("expected default 1920x1080 fallback")
	}
}
