package main

import (
	"os"
	"os/exec"
	"runtime"

	"nagare/internal/config"
	"nagare/internal/logging"
	"nagare/internal/settings"
	"nagare/internal/tray"
)

func main() {
	logger := logging.New()
	logger.Info("nagare starting")

	_, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	dbPath := "nagare.db"
	repo, err := settings.NewRepository(dbPath)
	if err != nil {
		logger.Error("failed to open database", "error", err)
		os.Exit(1)
	}
	defer repo.Close()
	logger.Info("database opened", "path", dbPath)

	svc := settings.NewService(repo)
	settingsServer := settings.NewServer(logger, svc)
	app := tray.New(logger)

	app.SetOnOpenSettings(func() {
		url, err := settingsServer.Start()
		if err != nil {
			logger.Error("failed to start settings server", "error", err)
			return
		}
		logger.Info("opening settings", "url", url)
		if err := openBrowser(url); err != nil {
			logger.Error("failed to open browser", "error", err)
		}
	})

	app.Run()
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()
}
