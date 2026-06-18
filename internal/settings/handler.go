package settings

import (
	"encoding/json"
	"net/http"
)

func (s *SettingsServer) handleUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(settingsHTML))
}

func (s *SettingsServer) handleSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		appSettings, err := s.service.GetAppSettings()
		if err != nil {
			http.Error(w, `{"error":"load settings"}`, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(appSettings)

	case http.MethodPut:
		var updated AppSettings
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
			return
		}
		if err := s.service.SaveAppSettings(&updated); err != nil {
			http.Error(w, `{"error":"save settings"}`, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updated)

	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func (s *SettingsServer) handleCameras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MockCameras)
}
