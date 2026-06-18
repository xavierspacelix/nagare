package camera

import (
	"testing"
)

func TestDiscover(t *testing.T) {
	m := NewManager(nil)
	cams := m.Discover()
	if len(cams) == 0 {
		t.Fatal("expected at least one camera")
	}
}

func TestOpenClose(t *testing.T) {
	m := NewManager(nil)
	cams := m.Discover()
	if len(cams) == 0 {
		t.Skip("no cameras available")
	}

	err := m.Open(cams[0].ID)
	if err != nil {
		t.Fatal("open camera:", err)
	}
	if !m.IsOpen() {
		t.Fatal("expected camera to be open")
	}

	m.Close()
	if m.IsOpen() {
		t.Fatal("expected camera to be closed")
	}
}

func TestReadFrame(t *testing.T) {
	m := NewManager(nil)
	cams := m.Discover()
	if len(cams) == 0 {
		t.Skip("no cameras available")
	}

	if err := m.Open(cams[0].ID); err != nil {
		t.Fatal("open:", err)
	}
	defer m.Close()

	frame, err := m.Read()
	if err != nil {
		t.Fatal("read frame:", err)
	}
	if frame.Width == 0 || frame.Height == 0 {
		t.Fatal("invalid frame dimensions")
	}
	if len(frame.Data) == 0 {
		t.Fatal("empty frame data")
	}
}
