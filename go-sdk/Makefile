SHELL := /bin/bash

CHAIN_ID ?= localnet
MONIKER ?= node0
COMETBFT_HOME ?= $(HOME)/.cometbft
PROXY_APP ?= unix:///tmp/newchain.sock
LOG_DIR ?= $(CURDIR)/logs
PID_DIR ?= $(CURDIR)/run

.PHONY: localnet init run stop status

init:
	CHAIN_ID=$(CHAIN_ID) MONIKER=$(MONIKER) COMETBFT_HOME=$(COMETBFT_HOME) PROXY_APP=$(PROXY_APP) \
		./scripts/cometbft-init.sh

run:
	CHAIN_ID=$(CHAIN_ID) MONIKER=$(MONIKER) COMETBFT_HOME=$(COMETBFT_HOME) PROXY_APP=$(PROXY_APP) LOG_DIR=$(LOG_DIR) PID_DIR=$(PID_DIR) \
		./scripts/run-localnet.sh

localnet: init run

stop:
	@if [ -f "$(PID_DIR)/cometbft.pid" ]; then kill "$$(cat $(PID_DIR)/cometbft.pid)" 2>/dev/null || true; fi
	@if [ -f "$(PID_DIR)/abci.pid" ]; then kill "$$(cat $(PID_DIR)/abci.pid)" 2>/dev/null || true; fi
	@rm -f "$(PID_DIR)/cometbft.pid" "$(PID_DIR)/abci.pid"

status:
	@echo "ABCI PID: $$(cat $(PID_DIR)/abci.pid 2>/dev/null || echo '-')"
	@echo "CometBFT PID: $$(cat $(PID_DIR)/cometbft.pid 2>/dev/null || echo '-')"
	@if [ -f "$(PID_DIR)/abci.pid" ]; then ps -p "$$(cat $(PID_DIR)/abci.pid)" >/dev/null 2>&1 && echo "ABCI: running" || echo "ABCI: not running"; else echo "ABCI: not running"; fi
	@if [ -f "$(PID_DIR)/cometbft.pid" ]; then ps -p "$$(cat $(PID_DIR)/cometbft.pid)" >/dev/null 2>&1 && echo "CometBFT: running" || echo "CometBFT: not running"; else echo "CometBFT: not running"; fi
