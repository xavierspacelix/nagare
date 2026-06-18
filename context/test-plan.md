# Test Plan — Nagare Hardware Verification

## Prerequisites

### Windows 11
- Go 1.24+ installed
- MinGW-w64 GCC for CGo (install via `choco install mingw`)
- OpenCV 4.6.0 installed (matching gocv v0.30.0)
- Webcam connected

### macOS
- Go 1.24+ installed
- Xcode Command Line Tools
- OpenCV 4.6.0 installed
- Webcam connected (built-in or external)

### Setup
```bash
# Install dependencies
go mod download

# Verify build
./scripts/build.sh linux    # stub build (quick check)
./scripts/build.sh test      # run all tests
```

---

## Phase 1 — Build Verification

| Test | Command | Expected |
|------|---------|----------|
| Stub build | `./scripts/build.sh linux` | Binary at `build/nagare-linux-stub` |
| Tests pass | `./scripts/build.sh test` | All tests pass |
| Vet clean | `./scripts/build.sh vet` | No warnings |

### Windows-specific
```bash
# Requires MinGW-w64
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o nagare.exe ./cmd/nagare/
```

### macOS-specific
```bash
# Run on macOS with Xcode
GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o nagare ./cmd/nagare/
```

---

## Phase 2 — Startup Verification

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 1 | Binary launches | Run `nagare.exe` | Tray icon appears |
| 2 | Settings opens | Click tray → Open Settings | Settings window opens in browser |
| 3 | Camera discovery | Open Settings → Camera section | Camera dropdown populated |
| 4 | No crash on close | Close settings or Exit tray | Process exits cleanly |

---

## Phase 3 — Camera & Vision

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 5 | Camera starts | Click Start Tracking | Camera LED turns on |
| 6 | Debug overlay | Enable debug mode | Overlay window shows camera feed |
| 7 | Hand detected | Hold hand in frame | Landmark dots on hand |
| 8 | No camera | Disconnect webcam, restart | Graceful error, no crash |
| 9 | Camera switch | Change camera in Settings | Stream switches |

---

## Phase 4 — Cursor Control

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 10 | Cursor follows hand | Open Palm, move hand | Cursor moves naturally |
| 11 | Left click | Pinch gesture | Left click registers |
| 12 | Right click | Two-finger pinch | Right click registers |
| 13 | Drag and drop | Pinch hold → move → release | Drags and drops |
| 14 | Scroll up | Scroll up gesture | Scrolls up |
| 15 | Scroll down | Scroll down gesture | Scrolls down |

---

## Phase 5 — Media & Volume

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 16 | Volume up | Two fingers up gesture | Volume increases |
| 17 | Volume down | Two fingers down gesture | Volume decreases |
| 18 | Mute | Mute action (if mapped) | Volume mutes |
| 19 | Media next | Swipe right | Next track |
| 20 | Media prev | Swipe left | Previous track |
| 21 | Play/pause | Open Palm (tracking on) | Toggle play/pause |

---

## Phase 6 — Multi-Monitor

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 22 | Cursor on primary | Hand moves on primary monitor | Cursor stays on primary |
| 23 | Cursor crosses to secondary | Hand moves to right edge | Cursor moves to secondary |
| 24 | Gesture on secondary | Perform pinch on secondary | Click on secondary |

---

## Phase 7 — Gesture Profiles

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 25 | Profile switch | Settings → Profiles → Switch | Mappings update instantly |
| 26 | Custom mapping | Change mapping in Settings | Gesture triggers new action |
| 27 | Profile persists | Restart app | Mappings preserved |

---

## Phase 8 — Performance

| # | Test | Target | Method |
|---|------|--------|--------|
| 28 | Startup time | < 2 seconds | Time from launch to tray ready |
| 29 | Tracking FPS | ≥ 24 FPS | Debug overlay FPS counter |
| 30 | Memory usage | < 150 MB | Task Manager / Activity Monitor |
| 31 | Idle CPU | < 10% | Task Manager when tracking disabled |
| 32 | Gesture latency | < 100ms | Visual: pinch → immediate click |

---

## Phase 9 — Stability

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 33 | 1-hour session | Run with tracking for 1 hour | No crash, no memory leak |
| 34 | 8-hour session | Run with tracking for 8 hours | Stable, FPS consistent |
| 35 | Start/Stop cycle | Start/Stop tracking 50 times | No crash, clean state each time |
| 36 | Restart cycle | Restart engine 20 times | No crash |
| 37 | Settings save/load | Change settings, restart | All settings preserved |

---

## Phase 10 — macOS-specific

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 38 | Camera permission | First launch | Permission prompt appears |
| 39 | Accessibility permission | Enable tracking | Accessibility request |
| 40 | Permission check | Settings → Status | Shows correct permission status |
| 41 | All controls | Repeat Phase 4-5 tests | Same behavior as Windows |

---

## Edge Cases

| # | Test | Steps | Expected |
|---|------|-------|----------|
| 42 | No webcam | Launch without camera | Graceful message, no crash |
| 43 | Multiple hands | Two hands in frame | Tracks primary hand only |
| 44 | Low light | Dim room | Reduced confidence but no crash |
| 45 | Rapid gestures | Quick successive pinches | Cooldown prevents repeat |
| 46 | Small hand/far away | Hand far from camera | Lower confidence, no false positive |
| 47 | Close to camera | Hand fills frame | Partial landmarks, graceful |
| 48 | Corrupt DB | Delete nagare.db, restart | Fresh DB created with defaults |

---

## Logging

During all tests, check `stdout` and application logs for:
- No `[ERROR]` messages during normal operation
- Warnings are acceptable for expected conditions (no camera, low confidence)
- No panic or stack traces

---

## Known Issues

- MediaPipe ONNX model (`assets/models/hand_landmark.onnx`) is not included
- Without model, landmark extraction falls back gracefully (stub returns nil)
- Cursor smoothing and scaling may need tuning per monitor setup
- Custom gesture training (user-trained gestures) is Phase 3 stretch goal
