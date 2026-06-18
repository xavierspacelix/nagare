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

type MappingDTO struct {
	ID          int    `json:"id"`
	GestureName string `json:"gesture_name"`
	ActionName  string `json:"action_name"`
	OnState     string `json:"on_state"`
	Enabled     bool   `json:"enabled"`
	CooldownMs  int    `json:"cooldown_ms"`
}

func (s *Service) GetMappingDTOs() ([]MappingDTO, error) {
	mappings, err := s.repo.GetGestureMappings()
	if err != nil {
		return nil, err
	}
	dtos := make([]MappingDTO, len(mappings))
	for i, m := range mappings {
		dtos[i] = MappingDTO{
			ID:          m.ID,
			GestureName: m.GestureName,
			ActionName:  m.ActionName,
			OnState:     m.OnState,
			Enabled:     m.Enabled,
			CooldownMs:  m.CooldownMs,
		}
	}
	return dtos, nil
}

func (s *Service) SaveMappingDTOs(dtos []MappingDTO) error {
	for _, dto := range dtos {
		m := &GestureMapping{
			ID:          dto.ID,
			GestureName: dto.GestureName,
			ActionName:  dto.ActionName,
			OnState:     dto.OnState,
			Enabled:     dto.Enabled,
			CooldownMs:  dto.CooldownMs,
		}
		if err := s.repo.SaveGestureMapping(m); err != nil {
			return err
		}
	}
	return nil
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
