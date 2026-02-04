#!/usr/bin/env bash
set -euo pipefail

CHAIN_ID=${CHAIN_ID:-localnet}
MONIKER=${MONIKER:-node0}
HOME_DIR=${COMETBFT_HOME:-$HOME/.cometbft}
PROXY_APP=${PROXY_APP:-unix:///tmp/newchain.sock}

if ! command -v cometbft >/dev/null 2>&1; then
  echo "cometbft binary not found in PATH" >&2
  exit 1
fi

mkdir -p "$HOME_DIR"

if [ ! -f "$HOME_DIR/config/config.toml" ]; then
  cometbft init --home "$HOME_DIR" --chain-id "$CHAIN_ID" --moniker "$MONIKER" >/dev/null
fi

CONFIG_FILE="$HOME_DIR/config/config.toml"

# Update proxy_app and other common settings
if command -v python3 >/dev/null 2>&1; then
  python3 - <<PY
from pathlib import Path
path = Path("$CONFIG_FILE")
text = path.read_text()
text = text.replace('proxy_app = "tcp://127.0.0.1:26658"', f'proxy_app = "{PROXY_APP}"')
text = text.replace('moniker = ""', f'moniker = "{MONIKER}"')
path.write_text(text)
PY
else
  # Fallback to sed
  sed -i.bak "s#proxy_app = \"tcp://127.0.0.1:26658\"#proxy_app = \"$PROXY_APP\"#" "$CONFIG_FILE"
  sed -i.bak "s#moniker = \"\"#moniker = \"$MONIKER\"#" "$CONFIG_FILE"
fi

echo "CometBFT initialized at $HOME_DIR"
