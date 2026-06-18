//go:build darwin

package main

import (
	"log/slog"

	"nagare/internal/permissions"
)

func checkPermissions(logger *slog.Logger) {
	logger.Info("checking macOS permissions")

	if permissions.Check(permissions.KindCamera) == permissions.StatusNotDetermined {
		logger.Warn("camera permission not determined yet")
	}

	if permissions.Check(permissions.KindCamera) != permissions.StatusGranted {
		logger.Warn("camera permission not granted — enable in System Preferences > Security & Privacy > Privacy > Camera")
	}

	if permissions.Check(permissions.KindAccessibility) != permissions.StatusGranted {
		logger.Warn("accessibility permission not granted — enable in System Preferences > Security & Privacy > Privacy > Accessibility")
	}

	logger.Info("permission check complete",
		"camera", permissions.StatusText(permissions.KindCamera),
		"accessibility", permissions.StatusText(permissions.KindAccessibility),
	)
}
