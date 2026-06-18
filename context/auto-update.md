# Auto Update System

## Overview

Nagare is a native desktop application distributed outside the Microsoft Store and Mac App Store.

The application runs primarily as a background service with a tray interface and supports automatic update detection and installation workflows.

The auto-update system must be implemented using:

* Go
* GitHub Releases

No cloud backend, VPS, database service, or third-party update server should be required.

The goal is to allow users to receive application updates seamlessly while preserving a fully native and offline-first architecture.

---

# Human Approval Requirement

Before performing any implementation, modification, refactoring, dependency installation, file creation, file deletion, database migration, or code generation, the Agent must first present a proposed plan and wait for explicit user approval.

## Required Workflow

1. Analyze the request.
2. Explain the proposed solution.
3. List affected files, folders, dependencies, and build changes.
4. Identify potential risks or breaking changes.
5. Ask for confirmation.
6. Wait for explicit approval before proceeding.

## Approval Keywords

The Agent may only proceed after receiving one of the following responses:

* approve
* approved
* continue
* proceed
* go ahead
* yes, implement
* yes, proceed

## Forbidden Actions Without Approval

The Agent must NOT perform the following without approval:

* Generate implementation code
* Create new files
* Modify existing files
* Delete files
* Install packages
* Update dependencies
* Change project structure
* Refactor existing code
* Execute terminal commands
* Generate pull requests
* Generate release binaries

## Expected Response Format

Before implementation:

### Proposed Plan

* Step 1
* Step 2
* Step 3

### Files Affected

* internal/updater/...
* internal/config/...
* build/...

### Risks

* Risk A
* Risk B

Please confirm before I proceed.

The Agent must stop after requesting confirmation and wait for user approval.

---

# Architecture Constraints

## Allowed Technologies

* Go
* GitHub Releases
* GitHub API
* Native OS Installers

## Allowed Libraries

* go-github
* net/http
* cobra
* viper

## Not Allowed

* Express.js
* Laravel
* NestJS
* Firebase
* Supabase
* Custom Update Servers
* Electron Updater
* Cloudflare Workers

Nagare must remain completely self-contained.

---

# Update Strategy

Application startup flow:

1. Read installed application version.
2. Check latest release from GitHub Releases.
3. Compare installed version against latest release.
4. Determine update status:
   * Up to date
   * Optional update
   * Mandatory update
5. Display tray notification.
6. Download release package.
7. Verify release integrity.
8. Launch installer.
9. Restart Nagare after installation.

---

# Release Source

Version information must come from GitHub Releases.

Repository structure:

```text
github.com/<organization>/nagare
```

Latest release should always be determined by:

* Latest stable release tag
* Highest version number

Example:

```text
v1.0.0
v1.0.1
v1.1.0
v2.0.0
```

Pre-releases must be ignored unless explicitly enabled.

---

# Release Assets

## Windows

Example:

```text
Nagare-Windows-x64.exe
```

or

```text
Nagare-Setup-x64.exe
```

## macOS

Example:

```text
Nagare-macOS-arm64.dmg
Nagare-macOS-intel.dmg
```

All downloads must use HTTPS.

---

# Update Channels

Supported channels:

* Stable
* Beta

Default:

```text
Stable
```

Users may opt into Beta releases through settings.

---

# Version Rules

Versions must follow Semantic Versioning.

Valid:

```text
v1.0.0
v1.0.1
v1.1.0
v2.0.0
```

Invalid:

```text
version1
release2
build123
```

Version comparison must use semantic version parsing.

Never compare raw strings.

---

# Mandatory Update Rules

If:

```text
current_version < minimum_supported_version
```

Then:

* Nagare must disable gesture processing.
* User must update before continuing.
* Dismiss button must be hidden.

Examples:

```text
Installed:
v1.0.0

Minimum:
v1.2.0
```

Update required.

---

# Optional Update Rules

If:

```text
current_version < latest_version
```

But:

```text
current_version >= minimum_supported_version
```

Then:

* User may continue using Nagare.
* Update notification is dismissible.
* Reminder appears again on next launch.

---

# Release Validation

Before installation:

Verify:

* Release exists
* Download completed
* File checksum matches
* Signature validation succeeds (when available)

Never install unsigned or corrupted releases.

---

# Windows Requirements

Installer format:

```text
EXE
```

Supported:

* Silent installation
* User installation
* Administrator installation

Autostart configuration must be preserved after update.

Current settings database must remain untouched.

---

# macOS Requirements

Installer format:

```text
DMG
```

Supported:

* Apple Silicon
* Intel

LaunchAgent configuration must remain preserved after update.

Current settings database must remain untouched.

---

# Rollback Strategy

If update installation fails:

1. Restore previous executable.
2. Restore previous configuration.
3. Restart previous version.
4. Log update failure.

User settings must never be lost.

---

# Update Frequency

Default checks:

```text
Once per day
```

Additional checks:

* Application startup
* Manual "Check for Updates"

Never poll continuously.

---

# User Experience Rules

Update notifications must be unobtrusive.

Allowed:

* Tray notification
* Native OS notification
* Settings window banner

Not Allowed:

* Full-screen interruptions
* Blocking dialogs for optional updates

---

# Logging

All update activity must be logged.

Examples:

```text
Checking for updates
Latest version found: v1.3.0
Downloading release
Verifying checksum
Launching installer
Update successful
```

Logs must never contain:

* GitHub tokens
* User personal data
* File system secrets

---

# Code Quality Requirements

Generate code using:

* Clean Architecture
* Dependency Injection
* Repository Pattern where applicable
* SOLID principles

Avoid:

* Business logic inside UI code
* Hardcoded versions
* Platform-specific code mixed into shared modules

---

# Deliverables

When implementing update-related features, generate:

* Update service
* GitHub release client
* Version comparison utilities
* Checksum validation
* Update notification UI
* Installer launcher
* Rollback mechanism
* Logging integration
* Unit tests
* Integration tests

All implementations must be production-ready.

---

# Final Rule

Nagare updates must remain:

* Native
* Secure
* Offline-first
* Cross-platform
* Independent from any custom backend infrastructure

GitHub Releases is the single source of truth for application updates.