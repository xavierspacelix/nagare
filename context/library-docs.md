# Library Docs

Project-specific usage patterns for every third-party library in Nagare.

This file defines how libraries must be used inside this project.

Always check:

1. AGENTS.md
2. MCP documentation (if available)
3. This file

Authority order:

MCP → AGENTS.md → library-docs.md → general knowledge
```

---

# OpenCV

## Purpose

Used only for:

* Camera access
* Frame capture
* Frame preprocessing

Not used for:

* Gesture recognition
* Business logic
* UI rendering

---

## Location

```text
internal/camera
internal/vision
```

Never import OpenCV elsewhere.

---

## Camera Initialization

```go
webcam, err := gocv.OpenVideoCapture(deviceID)
if err != nil {
    return err
}
defer webcam.Close()
```

Rules:

* Always release camera resources
* Never keep unused cameras open
* Only one active camera stream at a time
* Camera selection comes from settings

---

## Frame Processing

Pipeline:

```text
Capture
→ Resize
→ Flip (optional)
→ RGB Conversion
→ Landmark Detection
```

Never perform gesture logic inside frame processing.

---

# MediaPipe Hand Landmark Model

## Purpose

Used for:

* Hand detection
* Landmark extraction

Provides:

```text
21 Hand Landmarks
```

Used as input to Gesture Engine.

---

## Location

```text
internal/vision
```

---

## Rules

Never execute actions directly from landmarks.

Required flow:

```text
Landmarks
→ Gesture Engine
→ Action Mapping
→ Controller
```

---

# ONNX Runtime

## Purpose

Execute MediaPipe Hand Landmark model locally.

---

## Location

```text
internal/vision
```

---

## Initialization

Load model once.

Good:

```go
session := ort.NewSession(...)
```

Bad:

```go
for {
    session := ort.NewSession(...)
}
```

---

## Rules

* Load models at startup
* Reuse sessions
* Never reload per frame
* Never reload per gesture

---

# RobotGo

## Purpose

Desktop automation layer.

Used for:

* Mouse movement
* Mouse click
* Keyboard events
* Media keys
* Volume control

---

## Location

```text
internal/controller
```

Never import RobotGo elsewhere.

---

## Mouse Movement

```go
robotgo.Move(x, y)
```

---

## Click

```go
robotgo.Click()
```

---

## Scroll

```go
robotgo.ScrollDir(...)
```

---

## Rules

RobotGo must be wrapped behind:

```go
type OSController interface
```

Business logic must never call RobotGo directly.

---

# Systray

## Purpose

System tray application.

---

## Location

```text
internal/tray
```

---

## Tray Responsibilities

Allowed:

* Start tracking
* Stop tracking
* Open settings
* Exit application

Not allowed:

* Gesture recognition
* Camera processing
* Business logic

---

## Tray Menu Structure

```text
Nagare

Start Tracking
Stop Tracking
Camera
Settings
Restart Engine
Check Updates
Exit
```

---

# SQLite

## Purpose

Local persistence layer.

---

## Location

```text
internal/settings
```

---

## Stores

* Settings
* Gesture profiles
* Gesture mappings

Never store:

* Camera frames
* Images
* Videos
* Landmark history

---

## Rules

Every query must go through repository layer.

Good:

```text
UI
→ Service
→ Repository
→ SQLite
```

Bad:

```text
UI
→ SQLite
```

---

# Viper

## Purpose

Configuration management.

---

## Location

```text
internal/config
```

---

## Example

```yaml
tracking:
  sensitivity: 0.8

camera:
  default_device: 0
```

---

## Rules

Never hardcode:

* Sensitivity
* Cooldown values
* Camera IDs

---

# slog

## Purpose

Structured logging.

---

## Example

```go
logger.Info(
    "tracking started",
    "camera", cameraID,
)
```

---

## Levels

```text
Debug
Info
Warn
Error
```

---

## Never Log

* Webcam frames
* Landmark arrays
* User screenshots

---

# GoCV

## Purpose

OpenCV bindings for Go.

---

## Rules

Use GoCV as the only OpenCV wrapper.

Never mix:

* CGO wrappers
* Multiple OpenCV bindings

---

# GoReleaser

## Purpose

Build and packaging.

---

## Targets

### Windows

```text
amd64
arm64
```

### macOS

```text
amd64
arm64
```

---

## Rules

Every release must generate:

```text
Windows Installer
Windows ZIP
macOS App Bundle
macOS ZIP
```

---

# Future Libraries

Before adding any library:

Verify:

1. Standard library insufficient
2. Existing dependency insufficient
3. Cross-platform support exists
4. Documentation updated

If a new dependency is added:

Update:

```text
context/library-docs.md
context/tech-stack.md
context/architecture.md
```

before implementation begins.

---

# Final Rule

Libraries are implementation details.

Architecture rules always take precedence over library capabilities.

A library must adapt to Nagare architecture.

Nagare architecture must never adapt to a library.
