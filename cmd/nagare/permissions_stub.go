//go:build !darwin

package main

import "log/slog"

func checkPermissions(logger *slog.Logger) {
	logger.Debug("platform does not require macOS permission checks")
}
