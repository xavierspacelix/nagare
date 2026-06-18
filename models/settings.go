package models

type Settings struct {
	CameraID       string
	Sensitivity    float64
	Smoothing      float64
	StartupEnabled bool
	ActiveProfile  string
}
