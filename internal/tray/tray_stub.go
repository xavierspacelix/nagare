//go:build !windows && !darwin
// +build !windows,!darwin

package tray

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func (a *App) Run() {
	a.logger.Info("tray running in CLI mode (no desktop environment)")

	setupCLIHandlers()
	waitForQuit()
}

func (a *App) Stop() {
	a.logger.Info("tray stopping")
	os.Exit(0)
}

func setupCLIHandlers() {
	fmt.Println("=== Nagare Gesture Control ===")
	fmt.Println("Start Tracking  [1]")
	fmt.Println("Stop Tracking   [2]")
	fmt.Println("Open Settings   [3]")
	fmt.Println("Restart Engine  [4]")
	fmt.Println("Check Updates   [5]")
	fmt.Println("About           [6]")
	fmt.Println("Exit            [7]")
	fmt.Println("==============================")
}

func waitForQuit() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
