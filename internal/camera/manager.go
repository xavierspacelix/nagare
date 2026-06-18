package camera

import (
	"fmt"
	"log/slog"
)

type Manager struct {
	logger *slog.Logger
	active CamDevice
}

type CamDevice interface {
	Open(id int) error
	Read() (*Frame, error)
	Close() error
}

func NewManager(logger *slog.Logger) *Manager {
	if logger == nil {
		logger = slog.Default()
	}
	return &Manager{logger: logger}
}

func (m *Manager) Discover() []Info {
	return discoverCameras()
}

func (m *Manager) Open(id int) error {
	if m.active != nil {
		m.active.Close()
	}
	dev := newDevice()
	if err := dev.Open(id); err != nil {
		return fmt.Errorf("open camera %d: %w", id, err)
	}
	m.active = dev
	m.logger.Info("camera opened", "device_id", id)
	return nil
}

func (m *Manager) Read() (*Frame, error) {
	if m.active == nil {
		return nil, fmt.Errorf("no active camera")
	}
	return m.active.Read()
}

func (m *Manager) Close() {
	if m.active != nil {
		m.active.Close()
		m.active = nil
		m.logger.Info("camera closed")
	}
}

func (m *Manager) IsOpen() bool {
	return m.active != nil
}
