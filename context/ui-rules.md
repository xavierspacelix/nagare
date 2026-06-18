# UI Rules

Concise rules for building Nagare UI.

Design assets and UI tokens are the source of truth for all visual decisions.

The application is a desktop utility, not a content platform.

UI should remain lightweight, focused, and low-distraction.

---

## Font

Always use Inter as the primary font.

Never use system fonts as the primary font.

Typography should prioritize readability over visual flair.

---

## Layout

Nagare is tray-first.

Primary UI surfaces:

* Tray menu
* Settings window
* Notifications
* Camera preview (optional)

Default settings window:

```text
Width: 900px
Height: 650px
```

Minimum size:

```text
Width: 800px
Height: 600px
```

Content padding:

```text
24px
```

Section spacing:

```text
24px
```

Control spacing:

```text
12px
```

---

## Window Structure

Settings window layout:

```text
Header

Tracking Status

Settings Sections

Footer Actions
```

Avoid deeply nested layouts.

Prefer simple vertical flows.

---

## Header

Header contains:

* Application icon
* Nagare title
* Version badge

Height:

```text
64px
```

Header remains visually lightweight.

Never place large navigation systems in the header.

---

## Navigation

For MVP:

Use settings sections only.

Allowed:

```text
General
Camera
Gestures
Performance
About
```

Avoid:

* Sidebar navigation
* Multi-level navigation
* Dashboard layouts

---

## Cards

Every settings section lives inside a card.

```css
background: var(--color-surface);
border: 1px solid var(--color-border);
border-radius: var(--radius-lg);
padding: 24px;
```

Never use colored card backgrounds.

Color belongs in:

* Status indicators
* Badges
* Icons
* Buttons

Never on card surfaces.

---

## Typography Hierarchy

### Section Heading

```css
font-size: 16px;
font-weight: 600;
```

Used for:

* Settings sections
* Dialog titles
* Group labels

---

### Body Text

```css
font-size: 14px;
font-weight: 400;
```

Used for:

* Settings descriptions
* Content text

---

### Secondary Text

```css
font-size: 12px;
font-weight: 400;
```

Used for:

* Helper text
* Timestamps
* Status details

---

## Status Indicators

Tracking state must always be visible.

### Active

Use success tokens.

Label:

```text
Tracking Active
```

---

### Idle

Use neutral tokens.

Label:

```text
Tracking Stopped
```

---

### Error

Use error tokens.

Label:

```text
Camera Error
```

---

## Buttons

### Primary Button

Used for:

* Start Tracking
* Save Settings
* Apply Changes

Must use accent color.

---

### Secondary Button

Used for:

* Cancel
* Reset
* Close

Must use neutral styling.

---

### Danger Button

Used only for:

* Reset Profile
* Delete Gesture Profile

Must use error tokens.

---

## Form Inputs

All inputs must use UI tokens.

Fields:

```text
Text Input
Select
Toggle
Slider
Number Input
```

Never use custom input styles per page.

All inputs must be visually consistent.

---

## Toggles

Use toggles for binary settings only.

Examples:

```text
Auto Start
Enable Notifications
Enable Smoothing
Enable Debug Overlay
```

Do not use checkboxes unless explicitly required.

---

## Sliders

Use sliders for continuous values.

Examples:

```text
Sensitivity
Cursor Speed
Smoothing
Gesture Cooldown
```

Always display current value.

Example:

```text
Sensitivity: 75%
```

---

## Camera Preview

Camera preview is optional.

Purpose:

* Camera validation
* Gesture debugging
* Landmark visualization

Preview card:

```css
background: #000000;
border-radius: var(--radius-lg);
```

Landmarks must be drawn using UI tokens.

---

## Gesture Display

Recognized gestures may be displayed temporarily.

Example:

```text
Pinch Detected

Swipe Left

Volume Up
```

Display duration:

```text
1-2 seconds
```

Avoid persistent gesture history in MVP.

---

## Notifications

Notifications should be concise.

Good:

```text
Tracking Started
```

```text
Camera Disconnected
```

```text
Gesture Profile Saved
```

Bad:

```text
Nagare tracking service has successfully initialized and is now operating normally.
```

---

## Dialogs

Use dialogs only when necessary.

Allowed:

* Camera unavailable
* Critical errors
* Reset confirmation

Avoid confirmation dialogs for routine actions.

---

## Tray Menu

Tray menu should remain minimal.

Allowed:

```text
Start Tracking

Stop Tracking

Camera

Settings

Restart Engine

Check Updates

Exit
```

Do not place advanced configuration in the tray menu.

Open Settings instead.

---

## Empty States

Every empty state should explain:

* What is missing
* What the user should do next

Example:

```text
No camera detected.

Connect a camera and refresh.
```

---

## Error States

Never show raw system errors.

Bad:

```text
failed to create ONNX session
```

Good:

```text
Gesture engine could not start.
```

---

## Animations

Animations must be subtle.

Allowed:

* Fade
* Opacity transition
* Small scale transition

Avoid:

* Bounce
* Spring effects
* Long transitions

Maximum duration:

```text
250ms
```

---

## Accessibility

All controls must be keyboard accessible.

Provide:

* Focus states
* Keyboard navigation
* Readable contrast

Do not rely on color alone to communicate state.

---

## Cross Platform Consistency

Windows and macOS should provide equivalent experiences.

Platform-specific UI differences are acceptable only when required by native conventions.

Gesture behavior must remain identical.

---

## Do Nots

* Never use hardcoded colors.
* Never bypass UI tokens.
* Never create dashboard-style screens for MVP.
* Never build sidebar navigation.
* Never show raw system errors.
* Never use multiple accent colors.
* Never use gradients on cards.
* Never create settings pages larger than necessary.
* Never add animations that delay interaction.
* Never prioritize visual effects over responsiveness.

---

## Final Rule

Nagare is a utility application.

Every UI decision must prioritize:

1. Clarity
2. Responsiveness
3. Simplicity
4. Accessibility

Visual polish is important.

Operational clarity is more important.
