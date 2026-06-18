//go:build stub
// +build stub

package vision

import (
	"log/slog"
)

func newOverlayImpl(cfg OverlayConfig, logger *slog.Logger) (overlayImpl, error) {
	return &noopOverlayImpl{}, nil
}
