package settings

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Repository struct {
	db *sql.DB
}

type Settings struct {
	CameraID       string
	Sensitivity    float64
	Smoothing      float64
	StartupEnabled bool
	ActiveProfile  string
	ActiveProfileID int
}

type GestureMapping struct {
	ID          int
	ProfileID   int
	GestureName string
	ActionName  string
	OnState     string
	Enabled     bool
	CooldownMs  int
}

type GestureProfile struct {
	ID          int
	Name        string
	Description string
	IsDefault   bool
}

func NewRepository(dbPath string) (*Repository, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	repo := &Repository{db: db}
	if err := repo.migrate(); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return repo, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) migrate() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS settings (
			id              INTEGER PRIMARY KEY,
			camera_id       TEXT NOT NULL DEFAULT '',
			sensitivity     REAL NOT NULL DEFAULT 0.8,
			smoothing       REAL NOT NULL DEFAULT 0.5,
			startup_enabled INTEGER NOT NULL DEFAULT 0,
			active_profile  TEXT NOT NULL DEFAULT 'default',
			created_at      TEXT NOT NULL DEFAULT (datetime('now')),
			updated_at      TEXT NOT NULL DEFAULT (datetime('now'))
		)`,
		`CREATE TABLE IF NOT EXISTS gesture_mappings (
			id           INTEGER PRIMARY KEY,
			gesture_name TEXT NOT NULL,
			action_name  TEXT NOT NULL,
			on_state     TEXT NOT NULL DEFAULT 'active',
			enabled      INTEGER NOT NULL DEFAULT 1,
			cooldown_ms  INTEGER NOT NULL DEFAULT 200,
			created_at   TEXT NOT NULL DEFAULT (datetime('now'))
		)`,
		`CREATE TABLE IF NOT EXISTS gesture_profiles (
			id          INTEGER PRIMARY KEY,
			name        TEXT NOT NULL,
			description TEXT NOT NULL DEFAULT '',
			is_default  INTEGER NOT NULL DEFAULT 0,
			created_at  TEXT NOT NULL DEFAULT (datetime('now'))
		)`,
		`CREATE TABLE IF NOT EXISTS application_logs (
			id         INTEGER PRIMARY KEY,
			level      TEXT NOT NULL,
			source     TEXT NOT NULL,
			message    TEXT NOT NULL,
			created_at TEXT NOT NULL DEFAULT (datetime('now'))
		)`,
	}

	for _, q := range queries {
		if _, err := r.db.Exec(q); err != nil {
			return fmt.Errorf("exec migration: %w", err)
		}
	}

	r.db.Exec(`ALTER TABLE gesture_mappings ADD COLUMN on_state TEXT NOT NULL DEFAULT 'active'`)
	r.db.Exec(`ALTER TABLE gesture_mappings ADD COLUMN profile_id INTEGER NOT NULL DEFAULT 1`)

	return r.seed()
}

func (r *Repository) seed() error {
	if err := r.seedSettings(); err != nil {
		return err
	}
	if err := r.seedMappings(); err != nil {
		return err
	}
	return r.seedProfiles()
}

func (r *Repository) seedSettings() error {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM settings").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	_, err = r.db.Exec(`INSERT INTO settings (camera_id, sensitivity, smoothing, startup_enabled, active_profile)
		VALUES ('', 0.8, 0.5, 0, 'default')`)
	if err != nil {
		return fmt.Errorf("seed settings: %w", err)
	}
	return nil
}

func (r *Repository) seedMappings() error {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM gesture_mappings").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	defaults := []struct {
		gesture string
		action  string
		state   string
	}{
		{"open_palm", "tracking_on", "active"},
		{"closed_fist", "tracking_off", "active"},
		{"pinch", "left_click", "active"},
		{"pinch_hold", "mouse_down", "active"},
		{"pinch_hold", "mouse_up", "end"},
		{"two_finger_pinch", "right_click", "active"},
		{"two_finger_up", "volume_up", "active"},
		{"two_finger_down", "volume_down", "active"},
		{"swipe_left", "media_prev", "active"},
		{"swipe_right", "media_next", "active"},
		{"scroll_up", "scroll_up", "active"},
		{"scroll_down", "scroll_down", "active"},
	}

	for _, d := range defaults {
		_, err = r.db.Exec(`INSERT INTO gesture_mappings (gesture_name, action_name, on_state, enabled, cooldown_ms)
			VALUES (?, ?, ?, 1, 250)`, d.gesture, d.action, d.state)
		if err != nil {
			return fmt.Errorf("seed mapping %s: %w", d.gesture, err)
		}
	}
	return nil
}

func (r *Repository) seedProfiles() error {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM gesture_profiles").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	profiles := []struct {
		name        string
		description string
		isDefault   bool
	}{
		{"Default", "Standard gesture control settings", true},
		{"Work", "Optimized for productivity applications", false},
		{"Gaming", "Gesture controls for gaming", false},
		{"Presentation", "Simplified controls for presentations", false},
	}

	for _, p := range profiles {
		isDefault := 0
		if p.isDefault {
			isDefault = 1
		}
		_, err := r.db.Exec(`INSERT INTO gesture_profiles (name, description, is_default)
			VALUES (?, ?, ?)`, p.name, p.description, isDefault)
		if err != nil {
			return fmt.Errorf("seed profile %s: %w", p.name, err)
		}
	}
	return nil
}

func (r *Repository) GetSettings() (*Settings, error) {
	var s Settings
	var startupEnabled int
	err := r.db.QueryRow(`SELECT s.camera_id, s.sensitivity, s.smoothing, s.startup_enabled, s.active_profile,
		COALESCE((SELECT p.id FROM gesture_profiles p WHERE p.name = s.active_profile), 1)
		FROM settings s ORDER BY s.id LIMIT 1`).Scan(
		&s.CameraID, &s.Sensitivity, &s.Smoothing, &startupEnabled, &s.ActiveProfile, &s.ActiveProfileID,
	)
	if err != nil {
		return nil, fmt.Errorf("get settings: %w", err)
	}
	s.StartupEnabled = startupEnabled != 0
	return &s, nil
}

func (r *Repository) SaveSettings(s *Settings) error {
	startupEnabled := 0
	if s.StartupEnabled {
		startupEnabled = 1
	}
	_, err := r.db.Exec(`UPDATE settings SET
		camera_id = ?, sensitivity = ?, smoothing = ?,
		startup_enabled = ?, active_profile = ?,
		updated_at = ?
		WHERE id = (SELECT id FROM settings ORDER BY id LIMIT 1)`,
		s.CameraID, s.Sensitivity, s.Smoothing,
		startupEnabled, s.ActiveProfile,
		time.Now().UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("save settings: %w", err)
	}
	return nil
}

func (r *Repository) GetGestureMappings() ([]GestureMapping, error) {
	rows, err := r.db.Query(`SELECT id, profile_id, gesture_name, action_name, on_state, enabled, cooldown_ms
		FROM gesture_mappings ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("get gesture mappings: %w", err)
	}
	defer rows.Close()

	var mappings []GestureMapping
	for rows.Next() {
		var m GestureMapping
		if err := rows.Scan(&m.ID, &m.ProfileID, &m.GestureName, &m.ActionName, &m.OnState, &m.Enabled, &m.CooldownMs); err != nil {
			return nil, fmt.Errorf("scan mapping: %w", err)
		}
		mappings = append(mappings, m)
	}
	return mappings, rows.Err()
}

func (r *Repository) GetGestureMappingsByProfile(profileID int) ([]GestureMapping, error) {
	rows, err := r.db.Query(`SELECT id, profile_id, gesture_name, action_name, on_state, enabled, cooldown_ms
		FROM gesture_mappings WHERE profile_id = ? ORDER BY id`, profileID)
	if err != nil {
		return nil, fmt.Errorf("get gesture mappings by profile: %w", err)
	}
	defer rows.Close()

	var mappings []GestureMapping
	for rows.Next() {
		var m GestureMapping
		if err := rows.Scan(&m.ID, &m.ProfileID, &m.GestureName, &m.ActionName, &m.OnState, &m.Enabled, &m.CooldownMs); err != nil {
			return nil, fmt.Errorf("scan mapping: %w", err)
		}
		mappings = append(mappings, m)
	}
	return mappings, rows.Err()
}

func (r *Repository) SaveGestureMapping(m *GestureMapping) error {
	_, err := r.db.Exec(`UPDATE gesture_mappings SET
		gesture_name = ?, action_name = ?, on_state = ?, enabled = ?, cooldown_ms = ?
		WHERE id = ?`,
		m.GestureName, m.ActionName, m.OnState, m.Enabled, m.CooldownMs, m.ID,
	)
	if err != nil {
		return fmt.Errorf("save gesture mapping: %w", err)
	}
	return nil
}

func (r *Repository) GetGestureProfile(id int) (*GestureProfile, error) {
	var p GestureProfile
	err := r.db.QueryRow(`SELECT id, name, description, is_default
		FROM gesture_profiles WHERE id = ?`, id).Scan(&p.ID, &p.Name, &p.Description, &p.IsDefault)
	if err != nil {
		return nil, fmt.Errorf("get profile %d: %w", id, err)
	}
	return &p, nil
}

func (r *Repository) SaveGestureProfile(p *GestureProfile) error {
	isDefault := 0
	if p.IsDefault {
		isDefault = 1
	}
	_, err := r.db.Exec(`UPDATE gesture_profiles SET
		name = ?, description = ?, is_default = ?
		WHERE id = ?`,
		p.Name, p.Description, isDefault, p.ID,
	)
	if err != nil {
		return fmt.Errorf("save profile: %w", err)
	}
	return nil
}

func (r *Repository) CreateGestureProfile(p *GestureProfile) (int, error) {
	isDefault := 0
	if p.IsDefault {
		isDefault = 1
	}
	res, err := r.db.Exec(`INSERT INTO gesture_profiles (name, description, is_default)
		VALUES (?, ?, ?)`, p.Name, p.Description, isDefault)
	if err != nil {
		return 0, fmt.Errorf("create profile: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *Repository) DeleteGestureProfile(id int) error {
	_, err := r.db.Exec(`DELETE FROM gesture_profiles WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete profile: %w", err)
	}
	return nil
}

func (r *Repository) GetGestureProfiles() ([]GestureProfile, error) {
	rows, err := r.db.Query(`SELECT id, name, description, is_default
		FROM gesture_profiles ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("get profiles: %w", err)
	}
	defer rows.Close()

	var profiles []GestureProfile
	for rows.Next() {
		var p GestureProfile
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.IsDefault); err != nil {
			return nil, fmt.Errorf("scan profile: %w", err)
		}
		profiles = append(profiles, p)
	}
	return profiles, rows.Err()
}
