//go:build windows || darwin
// +build windows darwin

package tray

import (
	"github.com/getlantern/systray"
)

type menuItems struct {
	startTracking *systray.MenuItem
	stopTracking  *systray.MenuItem
	openSettings  *systray.MenuItem
	restartEngine *systray.MenuItem
	checkUpdates  *systray.MenuItem
	about         *systray.MenuItem
	quit          *systray.MenuItem
}

func (a *App) Run() {
	systray.Run(a.onReady, a.onExit)
}

func (a *App) Stop() {
	systray.Quit()
}

func (a *App) onReady() {
	systray.SetIcon(generateGrayIcon())
	systray.SetTitle("Nagare")
	systray.SetTooltip("Nagare Gesture Control")

	items := a.buildMenu()
	go a.listen(items)

	a.logger.Info("tray ready")
}

func (a *App) onExit() {
	a.logger.Info("tray stopped")
}

func (a *App) buildMenu() menuItems {
	start := systray.AddMenuItem("Start Tracking", "Enable gesture tracking")
	stop := systray.AddMenuItem("Stop Tracking", "Disable gesture tracking")
	stop.Disable()
	systray.AddSeparator()
	camera := systray.AddMenuItem("Camera", "Select camera")
	camera.Disable()
	systray.AddSeparator()
	settings := systray.AddMenuItem("Open Settings", "Configure Nagare")
	restart := systray.AddMenuItem("Restart Engine", "Restart the gesture engine")
	systray.AddSeparator()
	updates := systray.AddMenuItem("Check Updates", "Check for new version")
	systray.AddSeparator()
	about := systray.AddMenuItem("About", "About Nagare")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Exit", "Exit Nagare")

	return menuItems{
		startTracking: start,
		stopTracking:  stop,
		openSettings:  settings,
		restartEngine: restart,
		checkUpdates:  updates,
		about:         about,
		quit:          quit,
	}
}

func (a *App) listen(items menuItems) {
	for {
		select {
		case <-items.startTracking.ClickedCh:
			items.startTracking.Disable()
			items.stopTracking.Enable()
			systray.SetIcon(generateIcon())
			a.handle(ActionStart)

		case <-items.stopTracking.ClickedCh:
			items.stopTracking.Enable()
			items.startTracking.Disable()
			systray.SetIcon(generateGrayIcon())
			a.handle(ActionStop)

		case <-items.openSettings.ClickedCh:
			if a.onOpenSettings != nil {
				a.onOpenSettings()
			}
			a.handle(ActionSettings)

		case <-items.restartEngine.ClickedCh:
			a.handle(ActionRestart)

		case <-items.checkUpdates.ClickedCh:
			a.handle(ActionCheckUpdates)

		case <-items.about.ClickedCh:
			a.handle(ActionAbout)

		case <-items.quit.ClickedCh:
			a.handle(ActionExit)
			a.Stop()
		}
	}
}

func (a *App) handle(action Action) {
	a.logger.Info("tray action", "action", action)
	switch action {
	case ActionStart:
		if a.onStart != nil {
			go a.onStart()
		}
	case ActionStop:
		if a.onStop != nil {
			go a.onStop()
		}
	case ActionRestart:
		if a.onRestart != nil {
			go a.onRestart()
		}
	}
}
