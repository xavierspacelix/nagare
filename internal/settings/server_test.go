package settings

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"
)

func testServer(t *testing.T) *SettingsServer {
	repo, err := NewRepository(filepath(t))
	if err != nil {
		t.Fatal("new repo:", err)
	}
	t.Cleanup(func() { repo.Close(); os.Remove(filepath(t)) })

	svc := NewService(repo)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewServer(logger, svc)
}

func filepath(t *testing.T) string {
	return fmt.Sprintf("%s/test_%d.db", os.TempDir(), time.Now().UnixNano())
}

func waitForServer(url string) error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return nil
		}
		time.Sleep(10 * time.Millisecond)
	}
	return fmt.Errorf("server at %s not ready", url)
}

func TestServerServesUI(t *testing.T) {
	srv := testServer(t)
	url, err := srv.Start()
	if err != nil {
		t.Fatal("start server:", err)
	}
	defer srv.Stop()

	if err := waitForServer(url); err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get(url + "/")
	if err != nil {
		t.Fatal("get ui:", err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	if len(b) == 0 {
		t.Fatal("empty UI response")
	}
}

func TestServerServesSettings(t *testing.T) {
	srv := testServer(t)
	url, err := srv.Start()
	if err != nil {
		t.Fatal("start server:", err)
	}
	defer srv.Stop()

	if err := waitForServer(url); err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get(url + "/api/settings")
	if err != nil {
		t.Fatal("get settings:", err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	if len(b) == 0 {
		t.Fatal("empty settings response")
	}
}

func TestServerUpdatesSettings(t *testing.T) {
	srv := testServer(t)
	url, err := srv.Start()
	if err != nil {
		t.Fatal("start server:", err)
	}
	defer srv.Stop()

	if err := waitForServer(url); err != nil {
		t.Fatal(err)
	}

	body := `{"startup_enabled":true,"start_minimized":false,"minimize_to_tray":true}`
	req, _ := http.NewRequest(http.MethodPut, url+"/api/settings",
		bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("put settings:", err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	if len(b) == 0 {
		t.Fatal("empty response")
	}
}

func TestServerServesCameras(t *testing.T) {
	srv := testServer(t)
	url, err := srv.Start()
	if err != nil {
		t.Fatal("start server:", err)
	}
	defer srv.Stop()

	if err := waitForServer(url); err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get(url + "/api/cameras")
	if err != nil {
		t.Fatal("get cameras:", err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	if len(b) == 0 {
		t.Fatal("empty cameras response")
	}
}

func TestRepositoryPersists(t *testing.T) {
	path := filepath(t)
	repo, err := NewRepository(path)
	if err != nil {
		t.Fatal("new repo:", err)
	}
	defer repo.Close()
	defer os.Remove(path)

	s, err := repo.GetSettings()
	if err != nil {
		t.Fatal("get settings:", err)
	}
	if s.Sensitivity != 0.8 {
		t.Fatalf("expected 0.8, got %f", s.Sensitivity)
	}

	s.StartupEnabled = true
	s.Sensitivity = 0.5
	if err := repo.SaveSettings(s); err != nil {
		t.Fatal("save settings:", err)
	}

	repo.Close()

	repo2, err := NewRepository(path)
	if err != nil {
		t.Fatal("reopen repo:", err)
	}
	defer repo2.Close()

	s2, err := repo2.GetSettings()
	if err != nil {
		t.Fatal("get settings after reopen:", err)
	}
	if !s2.StartupEnabled {
		t.Fatal("expected startup_enabled=true after reopen")
	}
	if s2.Sensitivity != 0.5 {
		t.Fatalf("expected sensitivity=0.5, got %f", s2.Sensitivity)
	}
}
