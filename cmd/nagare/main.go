package main

import (
	"os"

	"nagare/internal/config"
	"nagare/internal/logging"
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

	app := tray.New(logger)
	app.Run()
}
