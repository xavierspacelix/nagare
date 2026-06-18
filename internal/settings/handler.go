package settings

import (
	"encoding/json"
	"net/http"
)

type mockSettings struct {
	StartupEnabled bool    `json:"startup_enabled"`
	StartMinimized bool    `json:"start_minimized"`
	MinimizeToTray bool    `json:"minimize_to_tray"`
	CameraID       string  `json:"camera_id"`
	Resolution     string  `json:"resolution"`
	GestureSens    float64 `json:"gesture_sensitivity"`
	CursorSens     float64 `json:"cursor_sensitivity"`
	GestureDelay   int     `json:"gesture_delay"`
}

type cameraInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var currentSettings = mockSettings{
	StartupEnabled: false,
	StartMinimized: true,
	MinimizeToTray: true,
	CameraID:       "cam-0",
	Resolution:     "640x480",
	GestureSens:    0.8,
	CursorSens:     0.7,
	GestureDelay:   200,
}

var mockCameras = []cameraInfo{
	{ID: "cam-0", Name: "Integrated Webcam"},
	{ID: "cam-1", Name: "USB Camera"},
}

func (s *SettingsServer) handleUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(settingsHTML))
}

func (s *SettingsServer) handleSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(currentSettings)

	case http.MethodPut:
		var updated mockSettings
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
			return
		}
		currentSettings = updated
		json.NewEncoder(w).Encode(currentSettings)

	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func (s *SettingsServer) handleCameras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockCameras)
}
