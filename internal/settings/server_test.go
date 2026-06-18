package settings

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"testing"
	"time"
)

func waitForServer(url string) error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return nil
		}
		time.Sleep(10 * time.Millisecond)
	}
	return fmt.Errorf("server at %s not ready after 100ms", url)
}

func TestServerServesUI(t *testing.T) {
	srv := NewServer(testLogger(t))
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
	srv := NewServer(testLogger(t))
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
	srv := NewServer(testLogger(t))
	url, err := srv.Start()
	if err != nil {
		t.Fatal("start server:", err)
	}
	defer srv.Stop()

	if err := waitForServer(url); err != nil {
		t.Fatal(err)
	}

	updatedBody := `{"startup_enabled":true,"start_minimized":false,"minimize_to_tray":true}`
	req, _ := http.NewRequest(http.MethodPut, url+"/api/settings",
		bytes.NewBufferString(updatedBody))
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
	srv := NewServer(testLogger(t))
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

func testLogger(t *testing.T) *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
