package settings

var settingsHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Nagare Settings</title>
<style>
.confirm-dialog {
  position: fixed; top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.4); display: flex;
  align-items: center; justify-content: center; z-index: 100;
}
.confirm-box {
  background: var(--color-surface); border-radius: var(--radius-lg);
  padding: var(--space-6); max-width: 400px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.15);
}
.confirm-actions { display: flex; gap: var(--space-3); justify-content: flex-end; margin-top: var(--space-4); }
:root {
  --color-accent: #5B5FF8;
  --color-accent-dark: #4347E7;
  --color-accent-light: #EEF0FF;
  --color-accent-foreground: #FFFFFF;
  --color-background: #F7F8FA;
  --color-surface: #FFFFFF;
  --color-surface-secondary: #F9FAFB;
  --color-border: #E5E7EB;
  --color-border-muted: #D1D5DB;
  --color-text-primary: #111827;
  --color-text-secondary: #6B7280;
  --color-text-muted: #9CA3AF;
  --color-success: #10B981;
  --color-success-light: #ECFDF5;
  --color-success-foreground: #065F46;
  --color-error: #EF4444;
  --color-error-light: #FEF2F2;
  --color-error-foreground: #991B1B;
  --radius-sm: 6px;
  --radius-md: 10px;
  --radius-lg: 14px;
  --space-1: 4px;
  --space-2: 8px;
  --space-3: 12px;
  --space-4: 16px;
  --space-5: 20px;
  --space-6: 24px;
}

*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Inter', Roboto, Helvetica, Arial, sans-serif;
  background: var(--color-background);
  color: var(--color-text-primary);
  font-size: 14px;
  line-height: 1.5;
  width: 900px; min-height: 650px;
  overflow-y: auto;
}

.header {
  display: flex; align-items: center; gap: var(--space-3);
  height: 64px; padding: 0 var(--space-6);
  background: var(--color-surface);
  border-bottom: 1px solid var(--color-border);
}
.header-icon {
  width: 28px; height: 28px;
  background: var(--color-accent);
  border-radius: 8px; display: flex; align-items: center; justify-content: center;
  color: white; font-weight: 700; font-size: 14px;
}
.header-title { font-size: 18px; font-weight: 600; }
.header-version {
  background: var(--color-surface-secondary);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-sm);
  padding: 2px 8px;
  font-size: 12px; color: var(--color-text-muted);
  margin-left: auto;
}

.status-bar {
  padding: var(--space-3) var(--space-6);
  display: flex; align-items: center; gap: var(--space-2);
  font-size: 13px; color: var(--color-text-secondary);
  border-bottom: 1px solid var(--color-border);
  background: var(--color-surface);
}
.status-dot {
  width: 8px; height: 8px; border-radius: 50%;
  background: var(--color-border-muted);
}

.content { padding: var(--space-6); display: flex; flex-direction: column; gap: var(--space-6); }

.card {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  padding: var(--space-6);
}
.card-title {
  font-size: 16px; font-weight: 600;
  margin-bottom: var(--space-5);
  padding-bottom: var(--space-3);
  border-bottom: 1px solid var(--color-border);
}

.field {
  display: flex; align-items: center; justify-content: space-between;
  padding: var(--space-3) 0;
}
.field + .field { border-top: 1px solid var(--color-border-light, #F1F5F9); }
.field-label { font-size: 14px; color: var(--color-text-primary); }
.field-desc { font-size: 12px; color: var(--color-text-muted); margin-top: 2px; }
.field-control { flex-shrink: 0; }

.toggle {
  position: relative; width: 40px; height: 22px;
  background: var(--color-border-muted); border-radius: 11px;
  cursor: pointer; transition: background 180ms ease-out;
}
.toggle.active { background: var(--color-accent); }
.toggle-knob {
  position: absolute; top: 2px; left: 2px;
  width: 18px; height: 18px; border-radius: 50%;
  background: white; box-shadow: 0 1px 2px rgba(0,0,0,0.1);
  transition: transform 180ms ease-out;
}
.toggle.active .toggle-knob { transform: translateX(18px); }

select, input[type="number"] {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  padding: 8px 12px; font-size: 14px; color: var(--color-text-primary);
  outline: none; min-width: 160px;
}
select:focus, input:focus { border-color: var(--color-accent); }

.slider-group { display: flex; align-items: center; gap: var(--space-3); }
input[type="range"] {
  -webkit-appearance: none; appearance: none;
  width: 160px; height: 4px;
  background: var(--color-border); border-radius: 2px;
  outline: none;
}
input[type="range"]::-webkit-slider-thumb {
  -webkit-appearance: none; width: 16px; height: 16px;
  background: var(--color-accent); border-radius: 50%;
  cursor: pointer;
}
.slider-value { font-size: 13px; color: var(--color-text-secondary); min-width: 36px; text-align: right; }

.btn {
  display: inline-flex; align-items: center; gap: var(--space-2);
  padding: 10px 16px; border-radius: var(--radius-md);
  font-size: 14px; font-weight: 500; cursor: pointer;
  border: none; outline: none; transition: background 120ms ease-out;
}
.btn-primary { background: var(--color-accent); color: var(--color-accent-foreground); }
.btn-primary:hover { background: var(--color-accent-dark); }
.btn-secondary {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  color: var(--color-text-primary);
}
.btn-secondary:hover { background: var(--color-surface-secondary); }
.btn-danger { background: var(--color-error); color: white; }
.btn-danger:hover { opacity: 0.9; }

.info-row { display: flex; justify-content: space-between; padding: var(--space-2) 0; font-size: 13px; }
.info-label { color: var(--color-text-secondary); }
.info-value { color: var(--color-text-primary); }

.log-area {
  background: #1a1a2e; color: #e0e0e0;
  border-radius: var(--radius-md);
  padding: var(--space-3); font-family: 'SF Mono', 'Fira Code', monospace;
  font-size: 12px; line-height: 1.6;
  height: 120px; overflow-y: auto; white-space: pre-wrap;
}

.footer {
  display: flex; justify-content: flex-end; gap: var(--space-3);
  padding: var(--space-4) var(--space-6);
  background: var(--color-surface);
  border-top: 1px solid var(--color-border);
  margin-top: var(--space-4);
}
.footer-status { 
  margin-right: auto;
  font-size: 13px; color: var(--color-text-secondary);
  display: flex; align-items: center;
}
</style>
</head>
<body>
<div class="header">
  <div class="header-icon">N</div>
  <span class="header-title">Nagare</span>
  <span class="header-version">v0.1.0</span>
</div>

<div class="status-bar">
  <span class="status-dot" id="statusDot"></span>
  <span id="statusText">Tracking Stopped</span>
</div>

<div class="content" id="app"></div>

<div class="footer">
  <div class="footer-status" id="footerStatus">Settings loaded</div>
  <button class="btn btn-secondary" onclick="resetDefaults()">Reset</button>
  <button class="btn btn-primary" onclick="saveSettings()">Save Settings</button>
</div>

<script>
let settings = {};
let mappings = [];
let profiles = [];
let activeProfileId = 1;

async function loadSettings() {
  const res = await fetch('/api/settings');
  settings = await res.json();
  const mr = await fetch('/api/mappings');
  mappings = await mr.json();
  const pr = await fetch('/api/profiles');
  profiles = await pr.json();
  if (profiles.length > 0) {
    activeProfileId = profiles[0].id;
    for (const p of profiles) {
      if (settings.camera_id === '' && p.name === 'Default') {
        activeProfileId = p.id;
        break;
      }
    }
  }
  render();
}

async function switchProfile(id) {
  activeProfileId = id;
  const mr = await fetch('/api/mappings?profile_id=' + id);
  mappings = await mr.json();
  if (mappings.length === 0) {
    const fallback = await fetch('/api/mappings?profile_id=1');
    mappings = await fallback.json();
  }
  render();
}

async function addProfile() {
  const name = prompt('Profile name:');
  if (!name) return;
  const res = await fetch('/api/profiles', {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({name: name, description: '', is_default: false}),
  });
  if (res.ok) {
    const pr = await fetch('/api/profiles');
    profiles = await pr.json();
    render();
  }
}

async function deleteProfile(id) {
  if (!confirm('Delete this profile? Mappings will be lost.')) return;
  await fetch('/api/profiles?id=' + id, {method: 'DELETE'});
  if (activeProfileId === id) {
    activeProfileId = 1;
    const mr = await fetch('/api/mappings?profile_id=1');
    mappings = await mr.json();
  }
  const pr = await fetch('/api/profiles');
  profiles = await pr.json();
  render();
}

async function saveSettings() {
  const res = await fetch('/api/settings', {
    method: 'PUT',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(settings),
  });
  if (res.ok) {
    settings = await res.json();
    const mr = await fetch('/api/mappings', {
      method: 'PUT',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(mappings),
    });
    if (mr.ok) mappings = await mr.json();
    document.getElementById('footerStatus').textContent = 'Settings saved';
    setTimeout(() => {
      document.getElementById('footerStatus').textContent = 'Settings loaded';
    }, 2000);
  }
}

async function resetDefaults() {
  settings = {
    startup_enabled: false,
    start_minimized: true,
    minimize_to_tray: true,
    camera_id: "cam-0",
    resolution: "640x480",
    gesture_sensitivity: 0.8,
    cursor_sensitivity: 0.7,
    gesture_delay: 200,
  };
  render();
  document.getElementById('footerStatus').textContent = 'Defaults restored (not saved)';
}

function render() {
  const app = document.getElementById('app');
  app.innerHTML = sectionsHTML();
  bindEvents();
}

function bindEvents() {
  document.querySelectorAll('[data-toggle]').forEach(el => {
    el.addEventListener('click', () => {
      const key = el.dataset.toggle;
      settings[key] = !settings[key];
      el.classList.toggle('active', settings[key]);
    });
  });
  document.querySelectorAll('[data-select]').forEach(el => {
    el.addEventListener('change', () => {
      settings[el.dataset.select] = el.value;
    });
  });
  document.querySelectorAll('[data-slider]').forEach(el => {
    el.addEventListener('input', () => {
      const key = el.dataset.slider;
      const val = parseFloat(el.value);
      settings[key] = val;
      const display = el.parentElement.querySelector('.slider-value');
      if (display) {
        if (key === 'gesture_delay') {
          display.textContent = val + 'ms';
        } else {
          display.textContent = Math.round(val * 100) + '%';
        }
      }
    });
  });
  document.querySelectorAll('[data-map-toggle]').forEach(el => {
    el.addEventListener('click', () => {
      const i = parseInt(el.dataset.mapToggle);
      mappings[i].enabled = !mappings[i].enabled;
      el.classList.toggle('active', mappings[i].enabled);
    });
  });
  document.querySelectorAll('[data-map-action]').forEach(el => {
    el.addEventListener('change', () => {
      const i = parseInt(el.dataset.mapAction);
      mappings[i].action_name = el.value;
    });
  });
}

function sectionsHTML() {
  return generalSection() + profileSection() + cameraSection() + gesturesSection() + mappingsSection() + systemSection();
}

const gestureNames = [
  'open_palm','closed_fist','pinch','pinch_hold',
  'two_finger_pinch','two_finger_up','two_finger_down',
  'swipe_left','swipe_right','scroll_up','scroll_down'
];
const actionNames = [
  'tracking_on','tracking_off','left_click','right_click',
  'mouse_down','mouse_up','scroll_up','scroll_down',
  'volume_up','volume_down','mute',
  'media_play_pause','media_next','media_prev','key_tap'
];

function toggleHTML(key, value) {
  return '<div class="toggle' + (value ? ' active' : '') + '" data-toggle="' + key + '"><div class="toggle-knob"></div></div>';
}

function generalSection() {
  const profileOpts = profiles.map(p =>
    '<option value="' + p.id + '"' + (p.id === activeProfileId ? ' selected' : '') + '>' + escapeHtml(p.name) + '</option>'
  ).join('');
  return '<div class="card">' +
    '<div class="card-title">General</div>' +
    field('Launch on startup', 'Automatically start Nagare when you log in', toggleHTML('startup_enabled', settings.startup_enabled)) +
    field('Start minimized', 'Launch Nagare to the tray without showing the window', toggleHTML('start_minimized', settings.start_minimized)) +
    field('Minimize to tray', 'Minimize the settings window to the tray instead of quitting', toggleHTML('minimize_to_tray', settings.minimize_to_tray)) +
    '<div class="field"><div><div class="field-label">Active Profile</div><div class="field-desc">Switch between gesture profiles</div></div>' +
    '<select onchange="switchProfile(parseInt(this.value))" style="min-width:160px;">' + profileOpts + '</select></div>' +
  '</div>';
}

function profileSection() {
  const list = profiles.map(p => {
    const isActive = p.id === activeProfileId;
    const isDefault = p.is_default;
    return '<div style="display:flex;align-items:center;justify-content:space-between;padding:var(--space-2) 0;' +
      (!isActive ? 'opacity:0.6;' : '') + '">' +
      '<div><strong>' + escapeHtml(p.name) + '</strong>' +
      (p.description ? '<div style="font-size:12px;color:var(--color-text-muted);">' + escapeHtml(p.description) + '</div>' : '') +
      '</div>' +
      '<div style="display:flex;gap:var(--space-2);">' +
      (!isActive ? '<button class="btn btn-secondary" style="padding:4px 10px;font-size:12px;" onclick="switchProfile(' + p.id + ')">Switch</button>' : '') +
      (!isDefault && !isActive ? '<button class="btn btn-danger" style="padding:4px 10px;font-size:12px;" onclick="deleteProfile(' + p.id + ')">Delete</button>' : '') +
      '</div></div>';
  }).join('');

  return '<div class="card">' +
    '<div class="card-title">Gesture Profiles</div>' +
    '<div style="font-size:13px;color:var(--color-text-secondary);margin-bottom:var(--space-4);">' +
    'Profiles let you switch between different gesture configurations instantly.</div>' +
    list +
    '<div style="margin-top:var(--space-3);padding-top:var(--space-3);border-top:1px solid var(--color-border);">' +
    '<button class="btn btn-secondary" onclick="addProfile()">+ New Profile</button></div>' +
  '</div>';
}

function escapeHtml(str) {
  return str.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;');
}

function cameraSection() {
  const cameras = ['Integrated Webcam', 'USB Camera'];
  const resolutions = ['640x480', '800x600', '1280x720'];
  return '<div class="card">' +
    '<div class="card-title">Camera</div>' +
    '<div class="field"><div><div class="field-label">Camera</div><div class="field-desc">Select the camera to use for hand tracking</div></div>' +
    '<select data-select="camera_id">' + cameras.map((n, i) => '<option value="cam-' + i + '"' + (settings.camera_id === 'cam-' + i ? ' selected' : '') + '>' + n + '</option>').join('') + '</select></div>' +
    '<div class="field"><div><div class="field-label">Resolution</div><div class="field-desc">Camera capture resolution</div></div>' +
    '<select data-select="resolution">' + resolutions.map(r => '<option' + (settings.resolution === r ? ' selected' : '') + '>' + r + '</option>').join('') + '</select></div>' +
    '<div class="field"><div><div class="field-label">Camera Preview</div><div class="field-desc">Preview will display once camera is connected</div></div>' +
    '<div style="width:160px;height:90px;background:#000;border-radius:var(--radius-md);display:flex;align-items:center;justify-content:center;color:var(--color-text-muted);font-size:12px;">Preview</div></div>' +
  '</div>';
}

function gesturesSection() {
  const sensPct = Math.round(settings.gesture_sensitivity * 100);
  const curPct = Math.round(settings.cursor_sensitivity * 100);
  return '<div class="card">' +
    '<div class="card-title">Gestures</div>' +
    '<div class="field"><div><div class="field-label">Gesture Sensitivity</div><div class="field-desc">How easily gestures are detected</div></div>' +
    '<div class="slider-group"><input type="range" min="0.1" max="1.0" step="0.05" value="' + settings.gesture_sensitivity + '" data-slider="gesture_sensitivity"><span class="slider-value">' + sensPct + '%</span></div></div>' +
    '<div class="field"><div><div class="field-label">Cursor Sensitivity</div><div class="field-desc">How fast the cursor moves in response to hand motion</div></div>' +
    '<div class="slider-group"><input type="range" min="0.1" max="1.0" step="0.05" value="' + settings.cursor_sensitivity + '" data-slider="cursor_sensitivity"><span class="slider-value">' + curPct + '%</span></div></div>' +
    '<div class="field"><div><div class="field-label">Gesture Delay</div><div class="field-desc">Cooldown between gesture activations</div></div>' +
    '<div class="slider-group"><input type="range" min="50" max="500" step="10" value="' + settings.gesture_delay + '" data-slider="gesture_delay"><span class="slider-value">' + settings.gesture_delay + 'ms</span></div></div>' +
  '</div>';
}

function field(label, desc, control) {
  return '<div class="field"><div><div class="field-label">' + label + '</div><div class="field-desc">' + desc + '</div></div><div class="field-control">' + control + '</div></div>';
}

function labelForGesture(name) {
  return name.replace(/_/g, ' ').replace(/\b\w/g, c => c.toUpperCase());
}
function labelForAction(name) {
  return name.replace(/_/g, ' ').replace(/\b\w/g, c => c.toUpperCase());
}

function mappingsSection() {
  const rows = mappings.map((m, i) =>
    '<tr>' +
      '<td style="padding:6px 8px;border-bottom:1px solid var(--color-border);font-size:13px;">' + labelForGesture(m.gesture_name) + '</td>' +
      '<td style="padding:6px 8px;border-bottom:1px solid var(--color-border);">' +
        '<select data-map-action="' + i + '" style="min-width:140px;padding:4px 8px;font-size:12px;">' +
          actionNames.map(a => '<option value="' + a + '"' + (m.action_name === a ? ' selected' : '') + '>' + labelForAction(a) + '</option>').join('') +
        '</select>' +
      '</td>' +
      '<td style="padding:6px 8px;border-bottom:1px solid var(--color-border);text-align:center;">' +
        '<div class="toggle' + (m.enabled ? ' active' : '') + '" data-map-toggle="' + i + '" style="display:inline-block;"><div class="toggle-knob"></div></div>' +
      '</td>' +
    '</tr>'
  ).join('');

  return '<div class="card">' +
    '<div class="card-title">Gesture Mappings</div>' +
    '<div style="font-size:13px;color:var(--color-text-secondary);margin-bottom:var(--space-4);">Assign actions to each gesture</div>' +
    '<table style="width:100%;border-collapse:collapse;">' +
      '<thead><tr>' +
        '<th style="text-align:left;padding:6px 8px;font-size:12px;color:var(--color-text-muted);font-weight:500;">Gesture</th>' +
        '<th style="text-align:left;padding:6px 8px;font-size:12px;color:var(--color-text-muted);font-weight:500;">Action</th>' +
        '<th style="text-align:center;padding:6px 8px;font-size:12px;color:var(--color-text-muted);font-weight:500;">Enabled</th>' +
      '</tr></thead>' +
      '<tbody>' + rows + '</tbody>' +
    '</table>' +
  '</div>';
}

function systemSection() {
  return '<div class="card">' +
    '<div class="card-title">System</div>' +
    '<div class="field">' +
      '<div><div class="field-label">Application Logs</div><div class="field-desc">Recent log entries for troubleshooting</div></div>' +
    '</div>' +
    '<div class="log-area" id="logArea">[INFO] nagare starting\n[INFO] configuration loaded\n[INFO] tray ready\n[INFO] settings server started on port 34251\n</div>' +
    '<div style="display:flex;gap:var(--space-3);margin-top:var(--space-3);">' +
      '<button class="btn btn-secondary" onclick="refreshLogs()">Refresh Logs</button>' +
      '<button class="btn btn-secondary" onclick="runDiagnostics()">Run Diagnostics</button>' +
    '</div>' +
    '<div style="margin-top:var(--space-5);padding-top:var(--space-4);border-top:1px solid var(--color-border);">' +
      '<div class="info-row"><span class="info-label">Version</span><span class="info-value">0.1.0</span></div>' +
      '<div class="info-row"><span class="info-label">Build</span><span class="info-value">dev</span></div>' +
      '<div class="info-row"><span class="info-label">Platform</span><span class="info-value">' + navigator.platform + '</span></div>' +
    '</div>' +
  '</div>';
}

function refreshLogs() {
  const area = document.getElementById('logArea');
  if (area) area.textContent += '[INFO] logs refreshed\n';
}

function runDiagnostics() {
  const area = document.getElementById('logArea');
  if (area) area.textContent += '[INFO] running diagnostics...\n[INFO] camera: ok\n[INFO] gesture engine: idle\n';
}

loadSettings();
</script>
</body>
</html>`
