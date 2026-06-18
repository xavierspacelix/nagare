# Code Standards

Implementation rules and conventions for the entire project.

The AI agent must follow these in every session without exception.

These rules prevent architectural drift across development sessions.

---

# Engineering Mindset

The AI agent on this project operates as a senior engineer.

This means:

* Think before implementing
* Read context files first
* Never assume requirements
* Scope is sacred
* Every feature must be testable
* Clean over clever
* One feature at a time
* Optimize only after validation
* Fail safely
* Design before coding

Before implementing any feature:

1. Read project-overview.md
2. Read architecture.md
3. Verify build phase
4. Verify performance targets

Never skip these steps.

---

# Go Standards

## General

* Use Go 1.24+
* Enable static analysis
* Use `gofmt`
* Use `go vet`
* Code must compile without warnings

---

## Naming

### Packages

Use lowercase.

Good:

```go
camera
gestures
controller
settings
```

Bad:

```go
Camera
GestureEngine
```

---

### Interfaces

Interfaces describe behavior.

Good:

```go
type GestureRecognizer interface
type OSController interface
```

Bad:

```go
type IGestureRecognizer interface
```

Never use the `I` prefix.

---

### Structs

Use PascalCase.

```go
type CameraManager struct
type GestureEngine struct
```

---

### Functions

Use descriptive names.

Good:

```go
DetectGesture()
StartTracking()
LoadSettings()
```

Bad:

```go
Run()
DoThing()
Handle()
```

---

## Error Handling

Never ignore errors.

Bad:

```go
result, _ := loadSettings()
```

Good:

```go
result, err := loadSettings()
if err != nil {
    return err
}
```

---

Every external operation must return errors.

Examples:

* Camera access
* SQLite access
* File operations
* ONNX Runtime initialization
* Configuration loading

---

Wrap errors with context.

```go
return fmt.Errorf("load settings: %w", err)
```

---

Never panic for expected failures.

Bad:

```go
panic(err)
```

Good:

```go
return fmt.Errorf("camera unavailable: %w", err)
```

---

# Project Structure

Never violate architecture boundaries.

Allowed:

```text
UI
 ↓
Application
 ↓
Domain
 ↓
Infrastructure
 ↓
Platform
```

Not allowed:

```text
UI
 ↓
Windows API
```

---

Business logic must never directly call:

* RobotGo
* Windows APIs
* macOS APIs
* SQLite

Use interfaces.

---

# Dependency Injection

Prefer constructor injection.

Good:

```go
func NewGestureEngine(
    recognizer GestureRecognizer,
    controller OSController,
) *GestureEngine
```

Avoid:

```go
var globalController *WindowsController
```

Never use global mutable state.

---

# Concurrency

Use goroutines carefully.

Allowed:

* Camera loop
* Frame processing
* Gesture engine
* Event queue

Avoid spawning uncontrolled goroutines.

Bad:

```go
go processFrame(frame)
```

inside every frame callback.

---

Always provide cancellation.

Use:

```go
context.Context
```

for long-running operations.

---

# Logging

Use structured logging only.

Preferred:

```go
logger.Info(
    "camera started",
    "camera_id", cameraID,
)
```

Avoid:

```go
fmt.Println("camera started")
```

---

Log levels:

```text
Debug
Info
Warn
Error
```

---

Never log:

* Camera frames
* Landmark data dumps
* User screenshots

---

# Camera Rules

Camera code belongs only in:

```text
internal/camera
```

Never access OpenCV directly from:

* UI
* Tray
* Settings

---

Camera lifecycle:

```text
Discover
Open
Stream
Process
Close
```

Always release camera resources.

---

# Vision Rules

OpenCV and ONNX Runtime code belongs only in:

```text
internal/vision
```

Never perform inference inside:

```text
controller/
tray/
settings/
```

---

Frame pipeline:

```text
Capture
→ Preprocess
→ Landmark Detection
→ Gesture Recognition
```

Never bypass the pipeline.

---

# Gesture Engine Rules

All gestures must use state machines.

Every gesture requires:

* Start State
* Active State
* End State

Never trigger actions from a single frame.

Bad:

```text
Pinch detected
→ Click
```

Good:

```text
Pinch Start
→ Pinch Active
→ Threshold Reached
→ Click
```

---

Every gesture requires:

* Threshold
* Cooldown
* Exit condition

---

# Platform Rules

Platform code belongs only in:

```text
internal/controller
```

Required:

```text
controller/
├── interface.go
├── windows.go
└── macos.go
```

Business logic must only interact with:

```go
type OSController interface
```

Never import platform files outside controller.

---

# Storage Rules

SQLite is the single source of truth.

Store:

* Settings
* Gesture mappings
* Profiles

Do not store:

* Camera frames
* User images
* Landmark history

---

Database access belongs only in:

```text
internal/settings
```

---

# UI Rules

UI is tray-first.

Allowed UI:

* Tray icon
* Tray menu
* Settings window
* Notifications

Never implement:

* Dashboard
* Analytics page
* Multi-window workflow

unless explicitly planned.

---

# Configuration

All configuration must be externalized.

Examples:

```yaml
camera:
  default_device: 0

tracking:
  sensitivity: 0.8
```

Never hardcode:

* Camera IDs
* Sensitivity values
* Cooldowns

---

# Performance Rules

Target FPS:

```text
≥ 30 FPS
```

Maximum startup:

```text
< 2 seconds
```

Maximum idle memory:

```text
< 150 MB
```

---

Avoid:

* Excessive allocations
* Repeated model loading
* Repeated camera initialization

Load expensive resources once.

---

# Testing

Every feature must be testable.

Required tests:

* Gesture recognition logic
* Mapping logic
* Settings management

Prefer:

```go
*_test.go
```

---

Never release a feature that cannot be manually verified.

---

# Comments

Code should explain itself.

Comments explain:

WHY

not

WHAT

Bad:

```go
// increment counter
counter++
```

Good:

```go
// debounce repeated gesture activations
counter++
```

---

Never leave:

```go
TODO
FIXME
HACK
```

in committed code.

---

# Dependencies

Before adding a package:

Ask:

1. Does Go standard library solve this?
2. Does existing project code solve this?
3. Is the dependency actively maintained?
4. Is the dependency cross-platform?

Approved dependencies:

* OpenCV
* ONNX Runtime
* RobotGo
* Systray
* SQLite driver
* Viper
* slog

Do not introduce additional dependencies without updating documentation.

---

# Import Rules

Standard library first.

```go
import (
    "context"
    "fmt"
)
```

Then external packages.

```go
import (
    "github.com/go-vgo/robotgo"
)
```

Then internal packages.

```go
import (
    "nagare/internal/controller"
)
```

---

# Security Rules

All image processing remains local.

Never:

* Upload frames
* Store screenshots
* Store video recordings
* Send biometric data

Nagare must remain fully functional offline.

---

# Final Rule

If implementation conflicts with documentation:

Documentation wins.

Update code to match documentation.

Never update implementation assumptions without updating documentation first.
