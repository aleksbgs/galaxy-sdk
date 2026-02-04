#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
LOG_DIR=${LOG_DIR:-"$ROOT_DIR/logs"}
PID_DIR=${PID_DIR:-"$ROOT_DIR/run"}

CHAIN_ID=${CHAIN_ID:-localnet}
MONIKER=${MONIKER:-node0}
COMETBFT_HOME=${COMETBFT_HOME:-"$HOME/.cometbft"}
PROXY_APP=${PROXY_APP:-unix:///tmp/newchain.sock}

mkdir -p "$LOG_DIR" "$PID_DIR"

# Ensure CometBFT is initialized and configured
CHAIN_ID="$CHAIN_ID" MONIKER="$MONIKER" COMETBFT_HOME="$COMETBFT_HOME" PROXY_APP="$PROXY_APP" \
  "$ROOT_DIR/scripts/cometbft-init.sh"

# Start ABCI server
(
  cd "$ROOT_DIR"
  GOFLAGS="-tags cometbft" \
    go run ./cmd/node --socket "$PROXY_APP" --chain-id "$CHAIN_ID" \
    2>&1 | tee "$LOG_DIR/abci.log"
) &
ABCI_PID=$!
echo "$ABCI_PID" > "$PID_DIR/abci.pid"

# Start CometBFT
(
  COMETBFT_HOME="$COMETBFT_HOME" \
    cometbft start --home "$COMETBFT_HOME" \
    2>&1 | tee "$LOG_DIR/cometbft.log"
) &
COMET_PID=$!
echo "$COMET_PID" > "$PID_DIR/cometbft.pid"

echo "ABCI PID: $ABCI_PID"
echo "CometBFT PID: $COMET_PID"
echo "Logs: $LOG_DIR/abci.log and $LOG_DIR/cometbft.log"
echo "PID files: $PID_DIR/abci.pid and $PID_DIR/cometbft.pid"

echo "Press Ctrl+C to stop both."

cleanup() {
  kill "$COMET_PID" "$ABCI_PID" 2>/dev/null || true
  rm -f "$PID_DIR/abci.pid" "$PID_DIR/cometbft.pid"
}
trap cleanup INT TERM

wait
