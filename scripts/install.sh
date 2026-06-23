#!/usr/bin/env bash

set -e

REPO="kolcman/linux-admin-guide"
BINARY_NAME="linux-guide"
INSTALL_PATH="/usr/local/bin"

echo "🐧 Installing Linux Admin Guide..."

# Detect OS
OS="$(uname -s)"
ARCH="$(uname -m)"

if [[ "$OS" != "Linux" ]]; then
  echo "❌ This tool supports Linux only."
  exit 1
fi

# Normalize architecture
if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
else
  echo "❌ Unsupported architecture: $ARCH"
  exit 1
fi

# Download URL (GitHub latest release)
URL="https://github.com/${REPO}/releases/latest/download/${BINARY_NAME}"

TMP_FILE="/tmp/${BINARY_NAME}"

echo "⬇️ Downloading from: $URL"

curl -L --fail "$URL" -o "$TMP_FILE"

chmod +x "$TMP_FILE"

# Move to bin
echo "📦 Installing to ${INSTALL_PATH}..."

sudo mv "$TMP_FILE" "${INSTALL_PATH}/${BINARY_NAME}"

echo "✅ Installed successfully!"

echo ""
echo "👉 Run:"
echo "   ${BINARY_NAME}"
echo ""
echo "🚀 Done."