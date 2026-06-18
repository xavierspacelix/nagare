# Build Plan

## Core Principle

Nagare is built as a **background-first native desktop application**.

Every system must be visible, testable, measurable, and documented before progressing to the next phase.

Development order:

1. Architecture
2. Mock gesture events
3. Functional implementation
4. Performance validation
5. Real-world testing
6. Documentation updates

No gesture feature is considered complete until it works reliably with a real webcam.

---

# Phase 1 — Foundation

## 01 Project Bootstrap

Create the complete project structure and core runtime.

### Deliverables

* Go workspace initialized
* Module structure created
* Dependency management configured
* Build scripts configured
* Logging system configured
* Configuration loader implemented

### Validation

* Application starts successfully
* Configuration loads correctly
* Logs written correctly

---

## 02 Tray Application

Build the tray-first application shell.

### UI

* Tray icon
* Start Gesture Control
* Stop Gesture Control
* Settings
* About
* Exit

### Logic

* Single instance application
* Tray lifecycle management
* Background service management

### Validation

* App runs entirely from tray
* No main dashboard window required

---

## 03 Settings Window

Build complete settings UI using mock data.

### UI

#### General

* Launch on startup
* Start minimized
* Minimize to tray

#### Camera

* Camera selector
* Camera preview
* Resolution selector

#### Gestures

* Gesture sensitivity
* Cursor sensitivity
* Gesture delay

#### System

* Logs
* Diagnostics
* Version information

### Validation

* All settings screens navigable
* Mock settings persist locally

---

## 04 SQLite Persistence

Implement local persistence layer.

### Logic

Store:

* Application settings
* Gesture mappings
* Selected camera
* Performance preferences

### Validation

* Settings persist across restarts

---

# Phase 2 — Computer Vision Foundation

## 05 Camera Engine

Implement webcam capture system.

### Logic

* Camera enumeration
* Camera selection
* Frame acquisition
* Frame buffering

### Validation

* Live camera feed stable
* Camera switching works

---

## 06 OpenCV Processing Pipeline

Implement image processing pipeline.

### Logic

Pipeline:

```text
Camera
→ Frame Capture
→ Frame Normalization
→ Frame Processing
→ Render Debug Overlay
```

### Validation

* Stable frame processing
* No memory leaks

---

## 07 MediaPipe Hand Tracking

Integrate MediaPipe Hand Landmark Model.

### Logic

Extract:

* Hand landmarks
* Finger states
* Hand orientation
* Tracking confidence

### Validation

* Reliable hand detection
* Stable landmark tracking

---

## 08 Debug Overlay

Build developer debugging interface.

### UI

Display:

* Camera feed
* Hand landmarks
* Finger labels
* Tracking status
* FPS counter

### Validation

* Overlay updates in real time

---

# Phase 3 — Gesture Engine

## 09 Gesture State Machine

Implement gesture lifecycle engine.

### Logic

Every gesture must support:

```text
Idle
→ Start
→ Active
→ End
```

Never trigger actions from a single frame.

### Validation

* State transitions observable

---

## 10 Core Gesture Recognition

Implement MVP gestures.

### Supported Gestures

#### Open Palm

```text
Enable Tracking
```

#### Closed Fist

```text
Disable Tracking
```

#### Pinch

```text
Left Click
```

#### Pinch Hold

```text
Drag
```

#### Two Finger Pinch

```text
Right Click
```

#### Two Finger Vertical Movement

```text
Scroll
```

### Validation

* Gesture recognition stable
* False positives minimized

---

## 11 Gesture Stabilization

Implement filtering system.

### Logic

* Confidence thresholds
* Debouncing
* Temporal smoothing
* Gesture cooldowns

### Validation

* Reduced accidental triggers

---

# Phase 4 — OS Control Layer

## 12 Windows Controller

Implement Windows 11 native control engine.

### Actions

* Cursor movement
* Left click
* Right click
* Drag and drop
* Scroll

### Validation

* Input feels native
* No visible lag

---

## 13 Media Controls

Implement media integrations.

### Actions

* Play/Pause
* Next Track
* Previous Track
* Mute

### Validation

* Works with common media apps

---

## 14 Volume Controls

Implement system volume control.

### Gestures

#### Two Fingers Up

```text
Volume Up
```

#### Two Fingers Down

```text
Volume Down
```

### Validation

* Smooth volume adjustment

---

## 15 Keyboard Shortcut Engine

Implement shortcut execution.

### Examples

* Alt + Tab
* Win + D
* Copy
* Paste
* Custom shortcuts

### Validation

* Shortcut execution reliable

---

# Phase 5 — Performance Optimization

## 16 Runtime Profiling

Measure:

* CPU usage
* Memory usage
* FPS
* Gesture latency

### Targets

```text
Startup < 2s
Memory < 150MB
Idle CPU < 10%
FPS ≥ 24
```

---

## 17 Pipeline Optimization

Optimize:

* Frame processing
* Landmark extraction
* Gesture recognition
* OS interaction

### Validation

Targets consistently achieved.

---

## 18 Reliability Testing

Perform long-running tests.

### Scenarios

* 1 hour runtime
* 4 hour runtime
* 8 hour runtime

### Validation

* No crashes
* No memory leaks
* No FPS degradation

---

# Phase 6 — Advanced Features

## 19 Custom Gesture Mapping

### UI

Allow users to assign actions to gestures.

### Logic

```text
Gesture
→ Action
```

Mappings stored in SQLite.

---

## 20 Multi-Monitor Support

### Logic

* Screen discovery
* Monitor boundaries
* Cursor mapping

### Validation

Works on 2+ monitor setups.

---

## 21 Gesture Profiles

### Features

* Work profile
* Gaming profile
* Presentation profile

### Validation

Profile switching works instantly.

---

# Phase 7 — Cross Platform (macOS)

> macOS is developed within the same project lifecycle but only after Windows MVP reaches production quality.

## 22 macOS Controller Layer

Implement native macOS control engine.

### Actions

* Cursor movement
* Mouse actions
* Scroll
* Keyboard shortcuts
* Media controls

### Validation

Equivalent behavior to Windows implementation.

---

## 23 macOS Permissions Workflow

Implement handling for:

* Camera permissions
* Accessibility permissions
* Input monitoring permissions

### Validation

User onboarding flow is clear and reliable.

---

## 24 Cross Platform Verification

Validate identical behavior across:

* Windows 11
* macOS

### Validation

* Same gestures
* Same mappings
* Same settings storage model
* Same user experience

---

# Feature Count

| Phase                                | Features |
| ------------------------------------ | -------: |
| Phase 1 — Foundation                 |        4 |
| Phase 2 — Computer Vision Foundation |        4 |
| Phase 3 — Gesture Engine             |        3 |
| Phase 4 — OS Control Layer           |        4 |
| Phase 5 — Performance Optimization   |        3 |
| Phase 6 — Advanced Features          |        3 |
| Phase 7 — Cross Platform (macOS)     |        3 |
| **Total**                            |   **24** |

---

# MVP Release Scope

Windows 11 MVP includes:

* Tray application
* Camera engine
* MediaPipe hand tracking
* Cursor control
* Left click
* Right click
* Drag and drop
* Scroll
* Volume control
* Media control
* Settings persistence
* Auto start
* Performance monitoring

---

# Production Readiness Criteria

Nagare is production-ready when:

* Startup time < 2 seconds
* Memory usage < 150 MB
* Idle CPU < 10%
* FPS ≥ 24
* Gesture latency feels instantaneous
* No webcam frames leave the device
* Application runs reliably for 8+ hours
* Windows 11 support is stable
* macOS support passes parity validation with Windows implementation
