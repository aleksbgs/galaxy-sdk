#!/usr/bin/env bash
set -euo pipefail

HOME_DIR=${COMETBFT_HOME:-$HOME/.cometbft}

if ! command -v cometbft >/dev/null 2>&1; then
  echo "cometbft binary not found in PATH" >&2
  exit 1
fi

cometbft start --home "$HOME_DIR"
