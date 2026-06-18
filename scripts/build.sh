#!/usr/bin/env bash
set -euo pipefail

APP="nagare"
OUTPUT_DIR="./build"

usage() {
  cat <<EOF
Usage: $0 <target>

Targets:
  linux            Build for Linux (stub — no camera/tracking)
  linux-native     Build for Linux with OpenCV (requires gocv)
  windows          Build for Windows 11 (requires MinGW-w64)
  macos            Build for macOS (requires Xcode)
  all              Build all targets
  test             Run stub tests
  vet              Vet all code
  clean            Remove build output

Examples:
  $0 linux       # Build stub version for Linux
  $0 windows     # Cross-compile for Windows (requires MinGW)
  $0 test        # Run test suite with stub tags
EOF
  exit 1
}

ensure_output() {
  mkdir -p "$OUTPUT_DIR"
}

build_linux() {
  echo "==> Building for Linux (stub)..."
  ensure_output
  CGO_ENABLED=0 go build -tags stub -o "$OUTPUT_DIR/${APP}-linux-stub" ./cmd/nagare/
  echo "    -> $OUTPUT_DIR/${APP}-linux-stub"
}

build_linux_native() {
  echo "==> Building for Linux (native, requires OpenCV 4.6.0)..."
  ensure_output
  go build -tags '!stub' -o "$OUTPUT_DIR/${APP}-linux-native" ./cmd/nagare/
  echo "    -> $OUTPUT_DIR/${APP}-linux-native"
}

build_windows() {
  echo "==> Building for Windows 11..."
  echo "    NOTE: Requires MinGW-w64 (x86_64-w64-mingw32-gcc)"
  echo "    Install: sudo apt install gcc-mingw-w64-x86-64"
  ensure_output
  GOOS=windows GOARCH=amd64 \
  CGO_ENABLED=1 \
  CC=x86_64-w64-mingw32-gcc \
  go build -tags '!stub' -o "$OUTPUT_DIR/${APP}.exe" ./cmd/nagare/
  echo "    -> $OUTPUT_DIR/${APP}.exe"
}

build_macos() {
  echo "==> Building for macOS..."
  echo "    NOTE: Requires Xcode Command Line Tools on macOS"
  ensure_output
  GOOS=darwin GOARCH=amd64 \
  CGO_ENABLED=1 \
  go build -tags '!stub' -o "$OUTPUT_DIR/${APP}-darwin" ./cmd/nagare/
  echo "    -> $OUTPUT_DIR/${APP}-darwin"
}

cmd_test() {
  echo "==> Running tests (stub)..."
  go test -tags stub -count=1 -v ./... 2>&1
}

cmd_vet() {
  echo "==> Vetting all code..."
  go vet -tags stub ./... 2>&1
  go vet -tags '!stub' ./... 2>&1
  echo "    -> OK"
}

cmd_clean() {
  echo "==> Cleaning..."
  rm -rf "$OUTPUT_DIR"
  echo "    -> done"
}

case "${1:-help}" in
  linux)        build_linux ;;
  linux-native) build_linux_native ;;
  windows)      build_windows ;;
  macos)        build_macos ;;
  all)
    build_linux
    echo ""
    build_linux_native
    echo ""
    cmd_test
    echo ""
    cmd_vet
    ;;
  test)  cmd_test ;;
  vet)   cmd_vet ;;
  clean) cmd_clean ;;
  help)  usage ;;
  *)     usage ;;
esac
