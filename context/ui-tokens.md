# UI Tokens

Design tokens for Nagare.

All colors, typography, spacing, and component values must be defined through tokens.

Never hardcode colors.

Never use framework default color palettes directly.

---

# Design Philosophy

Nagare is a utility-first desktop application.

Visual principles:

* Minimal
* Native-feeling
* Calm
* Technical
* Functional
* Low-distraction

The application spends most of its lifecycle inside the system tray.

The UI should feel closer to:

* Raycast
* Linear
* Notion Calendar
* Arc
* Warp

Avoid:

* Gaming aesthetics
* Neon themes
* Excessive gradients
* Overly playful UI

---

# Color System

## Primary Brand

Nagare uses a deep indigo accent.

```css
--color-accent: #5B5FF8;
--color-accent-dark: #4347E7;
--color-accent-light: #EEF0FF;
--color-accent-muted: #F6F7FF;
--color-accent-foreground: #FFFFFF;
```

Used for:

* Primary buttons
* Active controls
* Tracking indicators
* Focus states

---

## Backgrounds

```css
--color-background: #F7F8FA;
--color-surface: #FFFFFF;
--color-surface-secondary: #F9FAFB;
--color-surface-tertiary: #F3F4F6;
```

---

## Borders

```css
--color-border: #E5E7EB;
--color-border-light: #F1F5F9;
--color-border-muted: #D1D5DB;
```

---

## Text

```css
--color-text-primary: #111827;
--color-text-secondary: #6B7280;
--color-text-muted: #9CA3AF;
--color-text-dark: #374151;
```

---

## Success

Tracking enabled.

```css
--color-success: #10B981;
--color-success-light: #ECFDF5;
--color-success-foreground: #065F46;
```

---

## Warning

Tracking degraded.

```css
--color-warning: #F59E0B;
--color-warning-light: #FFFBEB;
--color-warning-foreground: #92400E;
```

---

## Error

Camera unavailable.

```css
--color-error: #EF4444;
--color-error-light: #FEF2F2;
--color-error-foreground: #991B1B;
```

---

## Information

System information.

```css
--color-info: #3B82F6;
--color-info-light: #EFF6FF;
--color-info-foreground: #1D4ED8;
```

---

# Typography

Font Family:

```css
Inter
```

---

## Type Scale

| Element         | Size | Weight |
| --------------- | ---- | ------ |
| App Title       | 24px | 700    |
| Section Heading | 18px | 600    |
| Card Heading    | 16px | 600    |
| Body            | 14px | 400    |
| Secondary Text  | 13px | 400    |
| Caption         | 12px | 400    |
| Micro Text      | 11px | 400    |

---

# Border Radius

```css
--radius-sm: 6px;
--radius-md: 10px;
--radius-lg: 14px;
--radius-xl: 18px;
--radius-full: 9999px;
```

---

# Spacing

| Token   | Value |
| ------- | ----- |
| space-1 | 4px   |
| space-2 | 8px   |
| space-3 | 12px  |
| space-4 | 16px  |
| space-5 | 20px  |
| space-6 | 24px  |
| space-8 | 32px  |

---

# Status Colors

## Tracking Status

### Active

```css
background: var(--color-success-light);
text: var(--color-success-foreground);
```

Label:

```text
Tracking Active
```

---

### Idle

```css
background: var(--color-surface-secondary);
text: var(--color-text-secondary);
```

Label:

```text
Tracking Stopped
```

---

### Error

```css
background: var(--color-error-light);
text: var(--color-error-foreground);
```

Label:

```text
Camera Error
```

---

# Gesture Status Indicators

## Gesture Detected

```css
background: var(--color-accent-light);
border: 1px solid var(--color-accent);
text: var(--color-accent);
```

---

## Gesture Executed

```css
background: var(--color-success-light);
border: 1px solid var(--color-success);
text: var(--color-success-foreground);
```

---

# Component Tokens

## Window

```css
background: var(--color-surface);
border: 1px solid var(--color-border);
border-radius: var(--radius-xl);
```

---

## Settings Card

```css
background: var(--color-surface);
border: 1px solid var(--color-border);
border-radius: var(--radius-lg);
padding: 24px;
```

---

## Primary Button

```css
background: var(--color-accent);
color: var(--color-accent-foreground);
border-radius: var(--radius-md);
padding: 10px 16px;
font-weight: 500;
```

---

## Secondary Button

```css
background: var(--color-surface);
border: 1px solid var(--color-border);
color: var(--color-text-primary);
border-radius: var(--radius-md);
padding: 10px 16px;
```

---

## Ghost Button

```css
background: transparent;
color: var(--color-text-secondary);
```

---

## Input

```css
background: var(--color-surface);
border: 1px solid var(--color-border);
border-radius: var(--radius-md);
padding: 10px 12px;
```

---

## Toggle

Enabled:

```css
background: var(--color-accent);
```

Disabled:

```css
background: var(--color-border-muted);
```

---

# Tray Icon States

## Idle

```text
Gray Icon
```

Tracking disabled.

---

## Active

```text
Indigo Icon
```

Tracking enabled.

---

## Gesture Detected

```text
Indigo Icon + Pulse
```

Gesture currently recognized.

---

## Error

```text
Red Icon
```

Camera unavailable.

---

# Camera Preview

Preview window should use:

```css
background: #000000;
border-radius: var(--radius-lg);
```

Landmark overlays:

```css
Hand Points: var(--color-accent)
Hand Connections: var(--color-accent-dark)
Gesture Highlight: var(--color-success)
```

---

# Animation Tokens

Duration:

```css
fast: 120ms;
normal: 180ms;
slow: 250ms;
```

Easing:

```css
ease-out
```

Avoid:

* Bouncy animations
* Spring effects
* Large transitions

Nagare should feel responsive and immediate.

---

# Shadows

## Small

```css
0 1px 2px rgba(0,0,0,0.05)
```

## Medium

```css
0 4px 12px rgba(0,0,0,0.08)
```

## Large

```css
0 12px 32px rgba(0,0,0,0.12)
```

---

# Invariants

* Never use hardcoded colors in components.
* Never use Tailwind default color palettes directly.
* Always use design tokens.
* Accent color is always Nagare Indigo.
* Success always indicates tracking or gesture execution.
* Error always indicates hardware or system failure.
* UI should remain calm and low-distraction.
* Desktop utility experience takes priority over marketing aesthetics.
* Settings window must remain usable on both Windows and macOS.
* Tray icon state must accurately reflect engine state.
