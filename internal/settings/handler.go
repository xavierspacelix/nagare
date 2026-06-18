package settings

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var handlerLogger = slog.New(slog.NewTextHandler(os.Stderr, nil))

func encodeJSON(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		handlerLogger.Warn("json encode failed", "error", err)
	}
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
		appSettings, err := s.service.GetAppSettings()
		if err != nil {
			http.Error(w, `{"error":"load settings"}`, http.StatusInternalServerError)
			return
		}
		encodeJSON(w, appSettings)

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
		encodeJSON(w, updated)

	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func (s *SettingsServer) handleCameras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encodeJSON(w, MockCameras)
}

func (s *SettingsServer) handleMappings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	profileIDStr := r.URL.Query().Get("profile_id")

	switch r.Method {
	case http.MethodGet:
		var mappings []MappingDTO
		var err error
		if profileIDStr != "" {
			var profileID int
			if _, e := fmt.Sscanf(profileIDStr, "%d", &profileID); e == nil {
				mappings, err = s.service.GetMappingsByProfile(profileID)
			} else {
				mappings, err = s.service.GetMappingDTOs()
			}
		} else {
			mappings, err = s.service.GetMappingDTOs()
		}
		if err != nil {
			http.Error(w, `{"error":"load mappings"}`, http.StatusInternalServerError)
			return
		}
		encodeJSON(w, mappings)

	case http.MethodPut:
		var mappings []MappingDTO
		if err := json.NewDecoder(r.Body).Decode(&mappings); err != nil {
			http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
			return
		}
		if err := s.service.SaveMappingDTOs(mappings); err != nil {
			http.Error(w, `{"error":"save mappings"}`, http.StatusInternalServerError)
			return
		}
		encodeJSON(w, mappings)

	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func (s *SettingsServer) handleProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		profiles, err := s.service.GetProfiles()
		if err != nil {
			http.Error(w, `{"error":"load profiles"}`, http.StatusInternalServerError)
			return
		}
		encodeJSON(w, profiles)

	case http.MethodPost:
		var dto ProfileDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
			return
		}
		created, err := s.service.CreateProfile(dto)
		if err != nil {
			http.Error(w, `{"error":"create profile"}`, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		encodeJSON(w, created)

	case http.MethodPut:
		var dto ProfileDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
			return
		}
		if err := s.service.SaveProfile(dto); err != nil {
			http.Error(w, `{"error":"save profile"}`, http.StatusInternalServerError)
			return
		}
		encodeJSON(w, dto)

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, `{"error":"missing id"}`, http.StatusBadRequest)
			return
		}
		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
			return
		}
		if err := s.service.DeleteProfile(id); err != nil {
			http.Error(w, `{"error":"delete profile"}`, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}
