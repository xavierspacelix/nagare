package profiler

import (
	"log/slog"
	"runtime"
	"sync"
	"time"
)

type Profiler struct {
	mu          sync.Mutex
	fps         float64
	frameTime   time.Duration
	latency     time.Duration
	peakLatency time.Duration
	frameCount  int64
	lastFrame   time.Time
	smoothing   float64
}

type Snapshot struct {
	FPS         float64
	FrameTime   time.Duration
	Latency     time.Duration
	PeakLatency time.Duration
	FrameCount  int64
	MemAlloc    uint64
	MemTotal    uint64
	MemSys      uint64
	NumGC       uint32
	NumGoroutine int
	Uptime      time.Duration
}

var startTime = time.Now()

func New() *Profiler {
	return &Profiler{
		smoothing: 0.1,
	}
}

func (p *Profiler) FrameStart() {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now()
	if !p.lastFrame.IsZero() {
		p.frameTime = now.Sub(p.lastFrame)
		instantFPS := 1.0 / p.frameTime.Seconds()
		if p.fps == 0 {
			p.fps = instantFPS
		} else {
			p.fps += p.smoothing * (instantFPS - p.fps)
		}
	}
	p.lastFrame = now
	p.frameCount++
}

func (p *Profiler) FrameEnd() {}

func (p *Profiler) TrackLatency(d time.Duration) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.latency = d
	if d > p.peakLatency {
		p.peakLatency = d
	}
}

func (p *Profiler) Snapshot() Snapshot {
	p.mu.Lock()
	defer p.mu.Unlock()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return Snapshot{
		FPS:           p.fps,
		FrameTime:     p.frameTime,
		Latency:       p.latency,
		PeakLatency:   p.peakLatency,
		FrameCount:    p.frameCount,
		MemAlloc:      m.Alloc,
		MemTotal:      m.TotalAlloc,
		MemSys:        m.Sys,
		NumGC:         m.NumGC,
		NumGoroutine:  runtime.NumGoroutine(),
		Uptime:        time.Since(startTime),
	}
}

func (p *Profiler) LogMetrics(logger *slog.Logger) {
	s := p.Snapshot()
	logger.Info("performance snapshot",
		"fps", round2(s.FPS),
		"frame_time_ms", s.FrameTime.Milliseconds(),
		"latency_ms", s.Latency.Milliseconds(),
		"peak_latency_ms", s.PeakLatency.Milliseconds(),
		"frames", s.FrameCount,
		"mem_alloc_mb", s.MemAlloc/1024/1024,
		"mem_sys_mb", s.MemSys/1024/1024,
		"goroutines", s.NumGoroutine,
		"uptime_s", int(s.Uptime.Seconds()),
	)
}

func round2(v float64) float64 {
	return float64(int(v*100)) / 100
}
