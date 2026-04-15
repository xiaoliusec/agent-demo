#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")" && pwd)"
BUILD_FRONTEND=1

if [[ "${1:-}" == "--skip-frontend" ]]; then
  BUILD_FRONTEND=0
fi

cd "$ROOT_DIR"

echo "[1/2] Building frontend..."
if [[ "$BUILD_FRONTEND" -eq 1 ]]; then
  (
    cd frontend
    npm run build
  )
else
  echo "Skipped frontend build (--skip-frontend)"
fi

echo "[2/2] Building backend binary..."
go build -tags production \
  -ldflags '-extldflags "-framework UniformTypeIdentifiers"' \
  -o testflow ./cmd/testflow

echo "Build complete: $ROOT_DIR/testflow"
