package main

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	"nagare/models"

	"nagare/internal/actions"
	"nagare/internal/camera"
	"nagare/internal/config"
	"nagare/internal/controller"
	"nagare/internal/display"
	"nagare/internal/gestures"
	"nagare/internal/logging"
	"nagare/internal/pipeline"
	"nagare/internal/profiler"
	"nagare/internal/settings"
	"nagare/internal/tray"
	"nagare/internal/vision"
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

	cm := camera.NewManager(logger)
	pl := vision.NewPipeline(vision.DefaultConfig())
	lm, err := vision.NewLandmarkExtractor(vision.DefaultLandmarkConfig(), logger)
	if err != nil {
		logger.Warn("landmark extractor not available", "error", err)
		lm = nil
	}

	ov, err := vision.NewDebugOverlay(vision.DefaultOverlayConfig(), logger)
	if err != nil {
		logger.Warn("debug overlay not available", "error", err)
		ov = nil
	}

	checkPermissions(logger)

	ctrl := controller.New()
	eng := actions.NewEngine(ctrl, logger)
	prof := profiler.New()

	loadMappings := func(profileID int) {
		mappingStore := gestures.NewMappingStore()
		repoMappings, err := repo.GetGestureMappingsByProfile(profileID)
		if err != nil || len(repoMappings) == 0 {
			repoMappings, err = repo.GetGestureMappingsByProfile(1)
		}
		if err == nil {
			custom := make([]gestures.Mapping, 0, len(repoMappings))
			for _, rm := range repoMappings {
				g, ok := gestures.GestureFromName(rm.GestureName)
				if !ok {
					continue
				}
				a, ok := gestures.ActionFromName(rm.ActionName)
				if !ok {
					continue
				}
				state := models.GestureActive
				if rm.OnState == "end" {
					state = models.GestureEnd
				}
				custom = append(custom, gestures.Mapping{Gesture: g, Action: a, OnState: state})
			}
			mappingStore.SetCustom(custom)
		}
		eng.SetMappings(mappingStore)
	}

	profileID, _ := svc.GetActiveProfileID()
	if profileID == 0 {
		profileID = 1
	}
	loadMappings(profileID)

	machineCfg := gestures.DefaultConfig()
	machine := gestures.NewMachine(machineCfg, eng.HandleGesture, logger)

	monitors, err := ctrl.GetMonitors()
	if err != nil {
		logger.Warn("failed to get monitors", "error", err)
	}
	dm, err := display.NewMapper(monitors)
	if err != nil {
		logger.Warn("failed to create display mapper", "error", err)
	}
	pipe := pipeline.NewRunner(cm, pl, lm, ov, machine, eng, ctrl, prof, dm, logger)

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
	app.SetOnStart(func() {
		if err := pipe.Start(); err != nil {
			logger.Error("failed to start pipeline", "error", err)
			return
		}
		pipe.SetTracking(true)
		logger.Info("gesture control started")
	})
	app.SetOnStop(func() {
		pipe.SetTracking(false)
		pipe.Stop()
		logger.Info("gesture control stopped")
	})
	app.SetOnRestart(func() {
		pipe.Stop()
		time.Sleep(500 * time.Millisecond)
		if err := pipe.Start(); err != nil {
			logger.Error("failed to restart pipeline", "error", err)
		}
		logger.Info("gesture control restarted")
	})

	logger.Info("nagare ready", "version", "0.1.0", "platform", runtime.GOOS)
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
