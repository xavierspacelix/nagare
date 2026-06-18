package tray

import (
	"log/slog"
)

type Action int

const (
	ActionStart Action = iota + 1
	ActionStop
	ActionSettings
	ActionRestart
	ActionCheckUpdates
	ActionAbout
	ActionExit
)

func (a Action) String() string {
	switch a {
	case ActionStart:
		return "start"
	case ActionStop:
		return "stop"
	case ActionSettings:
		return "settings"
	case ActionRestart:
		return "restart"
	case ActionCheckUpdates:
		return "check_updates"
	case ActionAbout:
		return "about"
	case ActionExit:
		return "exit"
	default:
		return "unknown"
	}
}

type App struct {
	logger    *slog.Logger
	onOpenSettings func()
}

func New(logger *slog.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Logger() *slog.Logger {
	return a.logger
}

func (a *App) SetOnOpenSettings(fn func()) {
	a.onOpenSettings = fn
}
