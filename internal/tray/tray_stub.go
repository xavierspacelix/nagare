//go:build !windows && !darwin
// +build !windows,!darwin

package tray

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (a *App) Run() {
	a.logger.Info("tray running in CLI mode (no desktop environment)")

	a.printMenu()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		switch line {
		case "1":
			a.logger.Info("tray action", "action", ActionStart)
			fmt.Println("→ Tracking started")
		case "2":
			a.logger.Info("tray action", "action", ActionStop)
			fmt.Println("→ Tracking stopped")
		case "3":
			if a.onOpenSettings != nil {
				a.onOpenSettings()
			}
			a.logger.Info("tray action", "action", ActionSettings)
		case "4":
			a.logger.Info("tray action", "action", ActionRestart)
			fmt.Println("→ Engine restarted")
		case "5":
			a.logger.Info("tray action", "action", ActionCheckUpdates)
			fmt.Println("→ Checking for updates...")
		case "6":
			a.logger.Info("tray action", "action", ActionAbout)
			fmt.Println("→ Nagare v0.1.0 — Desktop Gesture Control")
		case "7":
			a.logger.Info("tray action", "action", ActionExit)
			fmt.Println("bye")
			return
		default:
			fmt.Println("unknown option:", line)
		}
		a.printMenu()
	}
}

func (a *App) Stop() {
	a.logger.Info("tray stopping")
	os.Exit(0)
}

func (a *App) printMenu() {
	fmt.Println()
	fmt.Println("=== Nagare Gesture Control ===")
	fmt.Println("Start Tracking  [1]")
	fmt.Println("Stop Tracking   [2]")
	fmt.Println("Open Settings   [3]")
	fmt.Println("Restart Engine  [4]")
	fmt.Println("Check Updates   [5]")
	fmt.Println("About           [6]")
	fmt.Println("Exit            [7]")
	fmt.Print("> ")
}
