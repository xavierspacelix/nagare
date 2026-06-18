# Architecture

## Stack

| Layer              | Tool                  | Purpose                                   |
| ------------------ | --------------------- | ----------------------------------------- |
| Language           | Go                    | Main application runtime                  |
| Computer Vision    | OpenCV                | Camera access and frame processing        |
| Hand Tracking      | MediaPipe Hands Model | 21-point hand landmark detection          |
| AI Runtime         | ONNX Runtime          | Local inference execution                 |
| Desktop Automation | RobotGo               | Mouse, keyboard, volume and media control |
| Tray Application   | Systray               | System tray interface                     |
| Local Database     | SQLite                | Local configuration and mappings          |
| Logging            | slog                  | Structured application logging            |
| Configuration      | Viper                 | Configuration management                  |
| Packaging          | GoReleaser            | Cross-platform release builds             |

---

## Folder Structure

```text
/
├── AGENTS.md
├── context/
│   ├── project-overview.md
│   ├── architecture.md
│   ├── tech-stack.md
│   ├── gesture-specification.md
│   ├── platform-rules.md
│   ├── code-standards.md
│   ├── library-docs.md
│   ├── build-plan.md
│   ├── progress-tracker.md
│   └── performance-targets.md
│
├── cmd/
│   └── nagare/
│       └── main.go
│
├── internal/
│
│   ├── camera/
│   │   ├── camera.go
│   │   ├── manager.go
│   │   └── discovery.go
│
│   ├── vision/
│   │   ├── preprocessing.go
│   │   ├── landmarks.go
│   │   └── inference.go
│
│   ├── gestures/
│   │   ├── recognizer.go
│   │   ├── state_machine.go
│   │   ├── cooldown.go
│   │   └── mapping.go
│
│   ├── actions/
│   │   ├── mouse.go
│   │   ├── keyboard.go
│   │   ├── media.go
│   │   └── volume.go
│
│   ├── controller/
│   │   ├── interface.go
│   │   ├── windows.go
│   │   └── macos.go
│
│   ├── tray/
│   │   ├── tray.go
│   │   ├── menu.go
│   │   └── notifications.go
│
│   ├── settings/
│   │   ├── repository.go
│   │   ├── service.go
│   │   └── defaults.go
│
│   ├── updater/
│   │   ├── updater.go
│   │   └── releases.go
│
│   └── logging/
│       └── logger.go
│
├── models/
│   ├── hand_landmark.go
│   ├── gesture.go
│   ├── settings.go
│   └── camera.go
│
├── storage/
│   └── nagare.db
│
├── assets/
│   ├── icons/
│   └── models/
│       └── hand_landmark.onnx
│
└── releases/
```

---

## System Boundaries

| Folder      | Owns                                               |
| ----------- | -------------------------------------------------- |
| camera/     | Camera discovery, connection and frame acquisition |
| vision/     | Frame processing and landmark extraction           |
| gestures/   | Gesture recognition and state transitions          |
| actions/    | Gesture-to-action translation                      |
| controller/ | Platform-specific desktop control                  |
| tray/       | System tray interaction                            |
| settings/   | Local configuration management                     |
| updater/    | Application update system                          |
| models/     | Shared domain models                               |

---

## Data Flow

### Gesture Recognition

```text
Camera
    ↓
Frame Capture
    ↓
OpenCV Processing
    ↓
ONNX Runtime
    ↓
Hand Landmark Detection
    ↓
Gesture Recognition
    ↓
Action Mapping
    ↓
OS Controller
```

---

### Cursor Movement

```text
Hand Tracking
    ↓
Index Finger Position
    ↓
Screen Coordinate Mapping
    ↓
RobotGo
    ↓
Mouse Move
```

---

### Gesture Click

```text
Pinch Gesture
    ↓
Gesture State Machine
    ↓
Cooldown Validation
    ↓
OS Controller
    ↓
Left Click
```

---

### Volume Control

```text
Two Finger Up
    ↓
Gesture Recognition
    ↓
Volume Action
    ↓
OS Controller
    ↓
Volume Up
```

---

### Settings Update

```text
Tray Menu
    ↓
Settings Service
    ↓
SQLite
    ↓
Reload Gesture Engine
```

---

## SQLite Schema

### settings

| Column          | Type     | Notes                |
| --------------- | -------- | -------------------- |
| id              | integer  | Primary key          |
| camera_id       | text     | Active camera        |
| sensitivity     | real     | Tracking sensitivity |
| smoothing       | real     | Cursor smoothing     |
| startup_enabled | boolean  | Auto start           |
| active_profile  | text     | Current profile      |
| created_at      | datetime |                      |
| updated_at      | datetime |                      |

---

### gesture_mappings

| Column       | Type     | Notes                 |
| ------------ | -------- | --------------------- |
| id           | integer  | Primary key           |
| gesture_name | text     | Pinch, SwipeLeft, etc |
| action_name  | text     | Click, VolumeUp       |
| enabled      | boolean  |                       |
| cooldown_ms  | integer  |                       |
| created_at   | datetime |                       |

---

### gesture_profiles

| Column      | Type     | Notes        |
| ----------- | -------- | ------------ |
| id          | integer  | Primary key  |
| name        | text     | Profile name |
| description | text     | Optional     |
| is_default  | boolean  |              |
| created_at  | datetime |              |

---

### application_logs

| Column     | Type     | Notes             |
| ---------- | -------- | ----------------- |
| id         | integer  | Primary key       |
| level      | text     | info, warn, error |
| source     | text     | module name       |
| message    | text     |                   |
| created_at | datetime |                   |

---

## Camera Lifecycle

```text
Application Start
      ↓
Discover Cameras
      ↓
Load Selected Camera
      ↓
Open Camera Stream
      ↓
Start Processing Loop
      ↓
Gesture Detection
      ↓
Action Execution
```

---

## Platform Adapter Pattern

```go
type OSController interface {
    MoveMouse(x, y int)
    LeftClick()
    RightClick()
    MouseDown()
    MouseUp()
    Scroll(delta int)
    VolumeUp()
    VolumeDown()
    MediaPlayPause()
    MediaNext()
    MediaPrevious()
}
```

Implementations:

```text
controller/
├── interface.go
├── windows.go
└── macos.go
```

Business logic must never directly call OS APIs.

---

## Tray Menu Structure

```text
Nagare

├── Start Tracking
├── Stop Tracking
├── Camera
│   ├── Camera 1
│   ├── Camera 2
│   └── Camera 3
│
├── Sensitivity
│
├── Open Settings
│
├── Restart Engine
│
├── Check Updates
│
└── Exit
```

---

## Performance Targets

| Metric           | Target   |
| ---------------- | -------- |
| Startup          | < 2 sec  |
| Memory Usage     | < 150 MB |
| Idle CPU         | < 10%    |
| Tracking FPS     | ≥ 30 FPS |
| Gesture Response | < 100 ms |

---

## Invariants

Rules the AI agent must never violate:

* Business logic never directly calls Windows APIs.
* Business logic never directly calls macOS APIs.
* Gesture recognition must always pass through the state machine.
* Every gesture requires a cooldown strategy.
* Camera frames are never uploaded externally.
* No cloud dependency for gesture recognition.
* Platform implementations are isolated inside controller/.
* SQLite is the only application database.
* RobotGo usage is isolated behind OSController.
* Tray components never perform gesture recognition.
* Gesture recognition never performs UI operations.
* All image processing remains local.
* Application must remain functional without internet access.
* Windows and macOS must share the same gesture engine.
* A feature is not production-ready until both platform adapters exist.
