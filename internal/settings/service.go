package settings

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type AppSettings struct {
	StartupEnabled bool    `json:"startup_enabled"`
	StartMinimized bool    `json:"start_minimized"`
	MinimizeToTray bool    `json:"minimize_to_tray"`
	CameraID       string  `json:"camera_id"`
	Resolution     string  `json:"resolution"`
	GestureSens    float64 `json:"gesture_sensitivity"`
	CursorSens     float64 `json:"cursor_sensitivity"`
	GestureDelay   int     `json:"gesture_delay"`
}

type CameraInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *Service) GetAppSettings() (*AppSettings, error) {
	repo, err := s.repo.GetSettings()
	if err != nil {
		return defaultAppSettings(), nil
	}

	return &AppSettings{
		StartupEnabled: repo.StartupEnabled,
		CameraID:       repo.CameraID,
		GestureSens:    repo.Sensitivity,
		CursorSens:     repo.Sensitivity * 0.875,
		GestureDelay:   200,
		StartMinimized: true,
		MinimizeToTray: true,
		Resolution:     "640x480",
	}, nil
}

func (s *Service) SaveAppSettings(as *AppSettings) error {
	return s.repo.SaveSettings(&Settings{
		CameraID:       as.CameraID,
		Sensitivity:    as.GestureSens,
		Smoothing:      0.5,
		StartupEnabled: as.StartupEnabled,
		ActiveProfile:  "default",
	})
}

func defaultAppSettings() *AppSettings {
	return &AppSettings{
		StartupEnabled: false,
		StartMinimized: true,
		MinimizeToTray: true,
		CameraID:       "",
		Resolution:     "640x480",
		GestureSens:    0.8,
		CursorSens:     0.7,
		GestureDelay:   200,
	}
}

var MockCameras = []CameraInfo{
	{ID: "cam-0", Name: "Integrated Webcam"},
	{ID: "cam-1", Name: "USB Camera"},
}
