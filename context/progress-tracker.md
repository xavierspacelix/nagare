# Progress Tracker

Update this file after every completed feature. Any AI agent reading this should immediately know what is done, what is in progress, and what is next.

---

## Current Status

**Phase:** Phase 4 — OS Control Layer

**Last Completed:** 14 Volume Controls

**Next:** 15 Keyboard Shortcut Engine

---

## Progress

### Phase 1 — Foundation

* [x] 01 Project Bootstrap
* [x] 02 Tray Application
* [x] 03 Settings Window
* [x] 04 SQLite Persistence

### Phase 2 — Computer Vision Foundation

* [x] 05 Camera Engine
* [x] 06 OpenCV Processing Pipeline
* [x] 07 MediaPipe Hand Tracking
* [x] 08 Debug Overlay

### Phase 3 — Gesture Engine

* [x] 09 Gesture State Machine
* [x] 10 Core Gesture Recognition
* [x] 11 Gesture Stabilization

### Phase 4 — OS Control Layer

* [x] 12 Windows Controller
* [x] 13 Media Controls
* [x] 14 Volume Controls
* [ ] 15 Keyboard Shortcut Engine

### Phase 5 — Performance Optimization

* [ ] 16 Runtime Profiling
* [ ] 17 Pipeline Optimization
* [ ] 18 Reliability Testing

### Phase 6 — Advanced Features

* [ ] 19 Custom Gesture Mapping
* [ ] 20 Multi-Monitor Support
* [ ] 21 Gesture Profiles

### Phase 7 — Cross Platform (macOS)

* [ ] 22 macOS Controller Layer
* [ ] 23 macOS Permissions Workflow
* [ ] 24 Cross Platform Verification

---

## Current MVP Scope

### Included in MVP

* Camera access
* Webcam processing
* Hand landmark detection
* Gesture recognition
* Cursor movement
* Left click
* Right click
* Drag and drop
* Scroll
* Volume control
* Media control
* Tray application
* Settings persistence
* Auto start
* Local-only processing

### Excluded From MVP

* User-trained gestures
* Plugin ecosystem
* Cloud synchronization
* AI-assisted gesture creation

---

## Performance Targets

| Metric          | Target         |
| --------------- | -------------- |
| Startup Time    | < 2 seconds    |
| Memory Usage    | < 150 MB       |
| Idle CPU Usage  | < 10%          |
| FPS             | ≥ 24 FPS       |
| Gesture Latency | Near real-time |

---

## Platform Status

### Windows 11

Status: Not Started

Target:

* Full feature support
* Production release platform

### macOS

Status: Not Started

Target:

* Feature parity with Windows
* Implemented after Windows MVP reaches production quality
* Remains part of the same project roadmap

---

## Decisions Made During Build

* GoCV v0.30.0 required for OpenCV 4.6.0 (Ubuntu Noble). v0.32.1+ uses `cv::aruco::ArucoDetector` from OpenCV 4.7.0+ which does not exist in 4.6.0 headers.
* Chose ONNX Runtime for local inference execution.
* Chose RobotGo for desktop automation layer.
* Chose getlantern/systray for system tray on Windows/macOS.
* Tray icons generated programmatically with Go image/png (solid indigo/gray 32x32 PNG).
* Linux stub tray uses CLI fallback (stdout menu) since appindicator not always available.
* Build tags separate platform-specific tray implementations.

---

## Known Issues

* Camera `TestOpenClose` and `TestReadFrame` fail on Linux without webcam — expected, OpenCV returns "can't open camera" error. Tests still validate the error path.

Example:

* Camera initialization instability on specific USB webcams.
* Performance degradation on low-light environments.

---

## Notes

*Add notes here as the build progresses.*

Examples:

* Keep all webcam processing local.
* Never persist webcam frames.
* Maintain platform abstraction between Windows and macOS controllers.
* All gestures must pass through Gesture Engine state validation before OS actions are executed.

---

## Release Readiness Checklist

### MVP Release

* [ ] Stable webcam acquisition
* [ ] Stable hand tracking
* [ ] Stable cursor control
* [ ] Stable click recognition
* [ ] Stable scroll recognition
* [ ] Stable volume control
* [ ] Tray application completed
* [ ] SQLite persistence completed
* [ ] Performance targets achieved
* [ ] Documentation updated

### Production Release

* [ ] Windows QA completed
* [ ] 8-hour stability test passed
* [ ] Memory leak testing passed
* [ ] CPU usage within target
* [ ] Crash recovery tested
* [ ] Installer package created
* [ ] Versioning process documented

### Cross Platform Release

* [ ] macOS controller completed
* [ ] macOS permission flow completed
* [ ] Feature parity validated
* [ ] Cross-platform QA completed

---

## Last Updated

```text
Phase: Phase 1 — Foundation
Progress: 14 / 24 Features Completed
Next Milestone: MediaPipe Hand Tracking
```
