# Project Overview

## About the Project

Nagare is a desktop gesture control application that enables users to operate their computers using hand gestures captured through a webcam.

The application runs primarily in the background and provides a lightweight tray-based experience. Using computer vision and hand landmark tracking, Nagare translates natural hand movements into desktop actions such as cursor movement, clicks, scrolling, media controls, and volume adjustments.

The entire experience is designed to be local-first, privacy-focused, and responsive enough for daily desktop use without requiring specialized hardware.

---

## The Problem It Solves

Traditional desktop interaction depends heavily on physical devices such as a mouse, keyboard, trackpad, or presentation remote.

In many situations users want:

* Hands-free computer interaction
* Touchless media control
* Better presentation control
* Accessibility assistance
* Alternative input methods

Existing gesture control solutions are often:

* Expensive
* Hardware dependent
* Cloud connected
* Inaccurate
* Resource intensive

Nagare provides a lightweight desktop-native alternative that runs locally using a standard webcam.

---

## Application Components

```text
System Tray
Settings Window
Gesture Engine
Camera Manager
Gesture Recognition
Action Engine
Desktop Controller
```

---

## Navigation

Nagare is a tray-first application.

Primary interactions:

```text
Tray Icon

├── Start Tracking
├── Stop Tracking
├── Camera Selection
├── Sensitivity
├── Open Settings
├── Restart Engine
├── Check Updates
└── Exit
```

No dashboard is required for MVP.

---

## Core User Flow

### First Launch

* User installs Nagare
* Application starts in tray
* Camera discovery runs automatically
* User selects preferred camera
* Default gesture profile is loaded
* Tracking remains disabled until explicitly enabled

---

### Enable Tracking

* User clicks Start Tracking
* Camera stream initializes
* Gesture engine loads
* Hand tracking becomes active
* Tray icon indicates active state

---

### Cursor Control

* User raises their hand
* Gesture engine detects landmarks
* Cursor follows mapped finger position
* Movement is smoothed to reduce jitter

---

### Click Interaction

* User performs Pinch gesture
* Gesture engine validates activation threshold
* Cooldown validation occurs
* Left click is triggered

---

### Drag and Drop

* User performs Pinch Hold gesture
* Mouse button is held down
* Cursor movement continues
* Release gesture triggers mouse up

---

### Scrolling

* User performs Scroll gesture
* Vertical movement is translated into wheel events
* Smooth scrolling is applied

---

### Volume Control

* User performs Two Fingers Up

* Volume increases

* User performs Two Fingers Down

* Volume decreases

---

### Media Control

* Swipe Right

  * Next track

* Swipe Left

  * Previous track

* Open Palm

  * Play/Pause

---

### Settings Management

* User opens Settings
* Updates sensitivity
* Changes camera
* Enables auto start
* Updates gesture mappings

Changes are stored locally and applied immediately.

---

## Gesture Recognition Flow

### Camera Capture

* Webcam provides video frames
* OpenCV processes incoming frames
* Frames are passed to the gesture engine

### Landmark Detection

* MediaPipe Hand Landmark model detects:

  * Palm position
  * Finger positions
  * Hand orientation

### Gesture Recognition

* Landmarks are converted into gesture states
* Gesture state machine validates:

  * Activation threshold
  * Active state
  * Exit condition
  * Cooldown period

### Action Execution

* Gesture is mapped to an action
* Action is executed through platform adapter
* Desktop responds immediately

---

## Data Architecture

### Application Settings

Stored in SQLite.

Contains:

* Camera selection
* Sensitivity settings
* Cursor smoothing
* Auto start configuration
* Active gesture profile

---

### Gesture Profiles

Stored in SQLite.

Contains:

* Gesture definitions
* Action mappings
* Cooldown values
* Profile preferences

---

### Application Logs

Stored locally.

Contains:

* Startup events
* Camera events
* Gesture engine events
* Error logs

No camera images are stored.

---

## Features In Scope

* Camera discovery
* Camera selection
* Hand tracking
* Landmark extraction
* Gesture recognition
* Cursor movement
* Left click
* Right click
* Drag and drop
* Scrolling
* Volume control
* Media controls
* Tray application
* Settings window
* Auto start
* Local configuration storage
* Windows support
* macOS support

---

## Features Out of Scope

* Cloud processing
* Remote control via internet
* User account system
* Online synchronization
* Webcam recording
* Gesture analytics dashboard
* Team collaboration
* Mobile application
* Subscription system
* Browser extension
* Plugin marketplace
* Voice control
* Face recognition
* Biometric storage

---

## Gesture Examples

### Cursor Mode

```text
Index Finger
→ Move Cursor
```

### Click

```text
Pinch
→ Left Click
```

### Drag

```text
Pinch Hold
→ Mouse Down

Release
→ Mouse Up
```

### Scroll

```text
Two Finger Vertical Motion
→ Scroll
```

### Media Control

```text
Swipe Left
→ Previous

Swipe Right
→ Next

Open Palm
→ Play/Pause
```

### Volume Control

```text
Two Fingers Up
→ Volume Up

Two Fingers Down
→ Volume Down
```

---

## Privacy Principles

All processing occurs locally.

Nagare never:

* Uploads camera frames
* Stores user images
* Stores video recordings
* Sends biometric information
* Requires cloud inference

The application must remain fully functional without internet access.

---

## Target User

A desktop user who:

* Wants touchless computer interaction
* Uses presentations frequently
* Wants media controls without touching peripherals
* Values privacy
* Uses a webcam-equipped computer
* Wants an alternative input method

---

## Success Criteria

* User can install and start using Nagare in under 5 minutes
* Hand tracking initializes reliably
* Cursor movement feels natural
* Gesture recognition feels accurate and responsive
* Clicks and scrolling behave consistently
* Media controls respond immediately
* Application remains under performance targets
* No cloud dependency is required
* Camera frames never leave the device
* Windows and macOS provide equivalent behavior
* Application remains stable during long-running sessions
* Tray-based workflow remains simple and intuitive
