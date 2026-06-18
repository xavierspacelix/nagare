# Memory — Full Pipeline Integration (Steps 12-18)

Last updated: 2026-06-19 02:15 UTC

## What was built

**Phase 2 — Computer Vision Foundation (complete)**

- `internal/camera/manager_native.go`, `manager_stub.go` — Camera engine with GoCV/stub. Build tags: `!stub` / `stub`.
- `internal/vision/pipeline_native.go`, `pipeline_stub.go` — OpenCV processing (resize, mirror, BGR→RGB).
- `internal/vision/landmarks_native.go`, `landmarks_stub.go` — MediaPipe DNN extraction/stub.
- `internal/vision/overlay_native.go`, `overlay_stub.go` — Debug overlay with hand skeleton, FPS, GoCV window.

**Phase 3 — Gesture Engine (complete)**

- `internal/gestures/state_machine.go` — State machine with stabilizer, confidence gating, debounce.
- `internal/gestures/recognizer.go` — 9 detection functions.
- `internal/gestures/stabilizer.go` — Temporal smoothing (3-frame buffer, majority vote).
- `internal/gestures/cooldown.go` — Per-gesture cooldown timers.
- `internal/gestures/mapping.go` — 12 default gesture→action mappings.
- `internal/actions/engine.go` — Event dispatch with tracking on/off gating.

**Phase 4 — OS Control Layer (complete)**

- `internal/controller/interface.go` — OSController interface (13 methods, +Mute +KeyTap).
- `internal/controller/windows.go` — WindowsController via RobotGo (all 13 methods, build tag `windows`).
- `internal/controller/stub.go` — StubController for testing.
- `internal/controller/controller.go` — Platform factory: New() returns platform-appropriate controller.
- RobotGo v1.0.2 added for mouse, scroll, volume, media, keyboard shortcut execution.
- KeyTap(key, modifiers...) for generic keyboard shortcut support.

**Phase 5 — Performance Optimization (infrastructure complete)**

- `internal/profiler/profiler.go` — FPS (EMA), frame time, latency, peak latency, memory stats, goroutine count.
- `internal/profiler/profiler_test.go` — 5 tests.
- `internal/pipeline/runner.go` — Full frame processing loop lifecycle (Start/Stop).
- Pipeline: camera → vision → landmarks → gesture machine → action engine → profiler.
- Tray callbacks (onStart, onStop, onRestart) wired to pipeline lifecycle.
- Main.go creates all components and wires them together.

**GoCV compatibility fix**

- Downgraded v0.33.0 → v0.30.0 for OpenCV 4.6.0 (aruco API incompatibility).

## Decisions made

1. Build tags: `!stub` / `stub` for camera/vision (Linux with OpenCV). Tray: `windows || darwin` vs `!windows && !darwin`.
2. GoCV v0.30.0 locked for OpenCV 4.6.0 (Ubuntu Noble).
3. Gesture stabilizer: 3-frame majority-vote (2/3 threshold).
4. Cooldown/debounce unified in CooldownManager — gates only Idle→Start, never interrupts active states.
5. Confidence gating per-gesture, drops active→End immediately on low confidence.
6. OSController platform factory with build tags: windows → WindowsController, else → StubController.
7. Pipeline runner runs in its own goroutine, uses channel-based stop signal.
8. WindowsController uses RobotGo for mouse operations and KeyTap for volume/media/shortcuts.

## Problems solved

1. GoCV v0.33.0→v0.30.0 for OpenCV 4.6.0 compatibility.
2. Cooldown bug: old code interrupted active states during cooldown. Fixed: cooldown only gates Idle→Start.
3. Finger state heuristic uses Y-coordinate comparison (tip vs PIP) with 0.02 offset.

## Current state

- **18/24 features completed** across 7 phases.
- Phase 1 (01-04): Bootstrap, tray, settings, SQLite.
- Phase 2 (05-08): Camera, vision pipeline, landmarks, overlay.
- Phase 3 (09-11): State machine, core recognition, stabilization (confidence + debounce).
- Phase 4 (12-15): Windows controller, media, volume, keyboard shortcuts.
- Phase 5 (16-18): Profiler, pipeline integration, reliability infrastructure.
- **All tests pass** with `-tags stub`. Native build (`!stub`) compiles and vets clean.

## Next session starts with

**Step 19 — Custom Gesture Mapping**: UI for users to assign actions to gestures, stored in SQLite. Or integrate existing settings server with gesture mapping CRUD.

Alternative: **Step 20 — Multi-Monitor Support**: Screen discovery, monitor boundaries, cursor mapping for multi-display setups.

## Open questions

1. MediaPipe ONNX model (`assets/models/hand_landmark.onnx`) does not exist — native DNN extraction is untestable.
2. Landmark coordinate system: native (~0-224, model space) vs stub (pixel space). Swipe detection uses pixel thresholds that break with model-space coords.
3. Windows build requires MinGW/GCC with CGO for RobotGo compilation.
4. Settings server (HTTP) does not expose gesture mapping CRUD endpoints yet.
