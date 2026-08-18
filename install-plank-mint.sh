#!/usr/bin/env bash
# Plank dock setup for Linux Mint Cinnamon (mac-style bottom dock)
set -euo pipefail

USER_HOME="${HOME}"
PLANK_THEME_DIR="${USER_HOME}/.local/share/plank/themes/MacDock"
PLANK_CONFIG_DIR="${USER_HOME}/.config/plank/dock1"
AUTOSTART_DIR="${USER_HOME}/.config/autostart"

echo "==> Packages"
sudo apt update
sudo dpkg --configure -a 2>/dev/null || true
sudo apt -f install -y 2>/dev/null || true
sudo apt install -y plank

echo "==> User theme (writable, no /usr/share edits)"
mkdir -p "${PLANK_THEME_DIR}"
if [[ -f /usr/share/plank/themes/Default/dock.theme ]]; then
  cp /usr/share/plank/themes/Default/dock.theme "${PLANK_THEME_DIR}/dock.theme"
else
  cat > "${PLANK_THEME_DIR}/dock.theme" <<'EOF'
[PlankTheme]
TopRoundness=4
BottomRoundness=4
LineWidth=0
OuterStrokeColor=0;;;
FillStartColor=85;85;85;180
FillEndColor=85;85;85;180
InnerStrokeColor=0;;;
Font=Sans 10
EOF
fi
chmod u+w "${PLANK_THEME_DIR}/dock.theme"

echo "==> Plank config"
mkdir -p "${PLANK_CONFIG_DIR}"
mkdir -p "${PLANK_CONFIG_DIR}/launchers"

cat > "${PLANK_CONFIG_DIR}/settings" <<'EOF'
[PlankDockItem]
launchers=[]

[PlankDockTheme]
Theme=MacDock
IconSize=48
HideMode=0
Position=3
Monitor=0
Alignment=0.5
ZoomEnabled=true
ZoomPercent=150
TooltipsEnabled=true
FillBehavior=0
Offset=0
EOF

echo "==> Autostart"
mkdir -p "${AUTOSTART_DIR}"
cat > "${AUTOSTART_DIR/plank.desktop}" <<'EOF'
[Desktop Entry]
Type=Application
Name=Plank
Comment=Mac-style dock
Exec=plank
Icon=plank
Terminal=false
Categories=Utility;
X-GNOME-Autostart-enabled=true
X-MATE-Autostart-enabled=true
EOF

echo "==> Restart Plank"
killall plank 2>/dev/null || true
sleep 1
nohup plank >/dev/null 2>&1 &

echo
echo "Done."
echo "- Dock should appear at the bottom."
echo "- Top panel: remove/hide bottom Cinnamon panel if duplicate (Settings -> Desktop -> Panels)."
echo "- Add apps: drag icons from menu onto the dock."
echo "- Warnings about Wnck/GLib in terminal are harmless on Mint 22."
echo "- Do NOT run 'plank --preferences' unless you change Theme to MacDock first."
