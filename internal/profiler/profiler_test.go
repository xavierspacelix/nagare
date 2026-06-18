package profiler

import (
	"log/slog"
	"testing"
	"time"
)

func TestNewProfiler(t *testing.T) {
	p := New()
	if p == nil {
		t.Fatal("expected non-nil profiler")
	}
}

func TestFrameStart(t *testing.T) {
	p := New()

	p.FrameStart()
	time.Sleep(10 * time.Millisecond)
	p.FrameStart()

	s := p.Snapshot()
	if s.FrameCount != 2 {
		t.Fatalf("expected 2 frames, got %d", s.FrameCount)
	}
	if s.FPS == 0 {
		t.Fatal("expected non-zero FPS")
	}
	if s.FrameTime == 0 {
		t.Fatal("expected non-zero frame time")
	}
}

func TestTrackLatency(t *testing.T) {
	p := New()
	p.TrackLatency(15 * time.Millisecond)
	p.TrackLatency(5 * time.Millisecond)

	s := p.Snapshot()
	if s.Latency != 5*time.Millisecond {
		t.Fatalf("expected 5ms latency, got %v", s.Latency)
	}
	if s.PeakLatency != 15*time.Millisecond {
		t.Fatalf("expected 15ms peak latency, got %v", s.PeakLatency)
	}
}

func TestSnapshot(t *testing.T) {
	p := New()
	p.FrameStart()

	s := p.Snapshot()
	if s.FrameCount == 0 {
		t.Fatal("expected frame count > 0")
	}
	if s.MemAlloc == 0 {
		t.Fatal("expected non-zero memory alloc")
	}
	if s.NumGoroutine == 0 {
		t.Fatal("expected at least 1 goroutine")
	}
	if s.Uptime <= 0 {
		t.Fatal("expected positive uptime")
	}
}

func TestLogMetrics(t *testing.T) {
	p := New()

	p.FrameStart()
	p.TrackLatency(10 * time.Millisecond)
	p.LogMetrics(slog.Default())

	s := p.Snapshot()
	if s.FrameCount == 0 {
		t.Fatal("expected frame count > 0 after log")
	}
}
