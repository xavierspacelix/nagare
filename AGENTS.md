---

description: Instructions for building Nagare desktop gesture control application
globs: *
alwaysApply: true
---

---

# Nagare Development Agent

This project follows a documentation-first and architecture-first development workflow.

Before writing any code, always read the project context files and understand the current system state.

---

# Read Before Anything Else

Read these files in this exact order before implementing any feature:

1. context/project-overview.md
2. context/architecture.md
3. context/ui-tokens.md
4. context/ui-rules.md
5. context/code-standards.md
6. context/library-docs.md
7. context/build-plan.md
8. context/progress-tracker.md
9. context/auto-update.md

Never skip this sequence.

---

# Core Development Workflow

Every feature follows the same lifecycle:

1. Design architecture
2. Define gesture behavior
3. Build with mock events
4. Validate UX
5. Connect hardware integration
6. Test manually
7. Optimize performance
8. Update documentation

No implementation before architecture review.

No feature is complete until it is testable.

---

# Product Overview

Nagare is a desktop application that allows users to control their computer using hand gestures captured from a webcam.

The application runs primarily in the background and is controlled through a system tray interface.

Primary targets:

- Windows 11
- macOS (future phase)

Core capabilities:

- Hand tracking
- Gesture recognition
- Mouse control
- Keyboard shortcuts
- Volume control
- Media control
- Custom gesture mapping
- Auto start with operating system

---

# Technology Stack

## Language

Go

## Computer Vision

OpenCV

## Hand Tracking

MediaPipe Hand Landmark Model

## AI Runtime

ONNX Runtime

## Desktop Automation

RobotGo

## Tray Application

Systray

## Local Storage

SQLite

---

# Rules That Never Change

## Architecture

The application must remain native.

Never introduce:

- Electron
- Node.js backend
- Express
- Laravel
- NestJS
- Docker

Communication must occur through internal Go modules.

---

## Vision Pipeline

Required processing pipeline:

Camera
→ Frame Processing
→ Hand Detection
→ Landmark Extraction
→ Gesture Recognition
→ Action Engine
→ OS Controller

Do not bypass the gesture engine.

---

## Cross Platform

All platform-specific logic must be isolated.

Required structure:

controller/
├── windows.go
└── macos.go

Never place platform-specific code inside business logic.

---

## UI

The application is tray-first.

Allowed UI:

- Tray icon
- Tray menu
- Settings window
- Notifications

Do not build a full desktop dashboard unless explicitly planned.

---

## Storage

SQLite is the single source of truth.

Store:

- Settings
- Gesture mappings
- User preferences
- Camera selection

Do not introduce external databases.

---

# Performance Requirements

Startup Time:

- Less than 2 seconds

Memory Usage:

- Less than 150 MB

Idle CPU:

- Less than 10%

Target FPS:

- Minimum 24 FPS

Input latency should feel instantaneous.

---

# Gesture Rules

Gestures must be state-based.

Never trigger actions directly from a single frame.

Every gesture requires:

- Start state
- Active state
- End state

Example:

Open Palm
→ Tracking Enabled

Pinch
→ Left Click

Swipe Left
→ Previous

Swipe Right
→ Next

Two Fingers Up
→ Volume Up

Two Fingers Down
→ Volume Down

---

# Security Rules

All image processing must remain local.

Never:

- Upload webcam frames
- Store user images
- Transmit biometric information

No cloud dependency is allowed for gesture recognition.

---

# Build Rules

Always follow build-plan.md.

Build phases sequentially.

Never start a future phase before the current phase is completed.

Example:

If Phase 1 is incomplete:

- Do not build custom gesture editor
- Do not build plugin system
- Do not build cloud synchronization

---

# Progress Tracking

After every completed feature update:

- progress-tracker.md

Required updates:

- Current Status
- Last Completed
- Next Feature
- Known Issues

---

# Documentation First

Before implementing:

1. Verify feature exists in project-overview.md
2. Verify architecture exists in architecture.md
3. Verify feature belongs to current build phase

If documentation is missing:

Stop and ask for clarification.

Never invent requirements.

---

# Available Commands

## /architect

Use before:

- New gesture systems
- New platform integrations
- Major architecture changes

Design first.

Code second.

---

## /review

Use when:

- Feature is finished
- Performance optimization completed
- Before moving to next phase

---

## /recover

Use when:

- Bugs persist after one fix attempt
- Architecture drift occurs
- State becomes inconsistent

---

## /remember save

Use when:

- Development spans multiple sessions

---

## /remember restore

Use when:

- Continuing unfinished work

---

# Error Handling

Never expose raw system errors.

Bad:

"camera initialization failed: device index out of range"

Good:

"Selected camera is unavailable."

---

# Code Quality

Prefer:

- Simple code
- Explicit naming
- Composition over inheritance
- Small modules

Avoid:

- Premature optimization
- Unnecessary abstractions
- Hidden side effects

Code should be understandable by a mid-level Go developer.

---

# MVP Scope

Phase 1:

- Camera access
- Hand detection
- Cursor movement
- Left click
- Scroll
- Volume control
- Tray application

Phase 2:

- Custom gesture mapping
- Multi-monitor support
- macOS support

Phase 3:

- User-trained gestures
- Plugin ecosystem
- Advanced automation

---

# Final Rule

If documentation conflicts with implementation:

Documentation wins.

Update code to match documentation.

Never update implementation assumptions without updating documentation first.
