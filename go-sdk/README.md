# Galaxy SDK (Cosmos-SDK-like)

Galaxy SDK is a minimal, Cosmos-SDK-inspired framework written in Go. It provides:

- Base app lifecycle (`InitChain`, `BeginBlock`, `DeliverTx`, `EndBlock`)
- Module system with message routing
- Simple in-memory KV store
- Example modules: `auth`, `bank`, `staking`

## Architecture (high level)

- `core/` defines interfaces for apps, modules, messages, transactions, and context.
- `base/` provides a `BaseApp` that wires modules and routes messages by `Msg.Type()`.
- `store/` provides a basic KV store interface and a memory-backed store.
- `modules/` includes example modules and keepers.
- `app/` wires everything together into a usable application.

## Quick start

```go
package main

import (
	"time"

	"github.com/aleksbgs/galaxy-sdk/app"
	"github.com/aleksbgs/galaxy-sdk/core"
	"github.com/aleksbgs/galaxy-sdk/modules/bank"
	"github.com/aleksbgs/galaxy-sdk/types"
)

func main() {
	chain := app.New("localnet")

	// mint some funds to alice
	chain.BankKeeper.Mint(types.Address("alice"), 1000)

	// send tx
	tx := core.SimpleTx{
		Msgs: []core.Msg{
			bank.MsgSend{
				From:   types.Address("alice"),
				To:     types.Address("bob"),
				Amount: 250,
			},
		},
	}

	_ = chain.Tick(time.Now(), tx)
}
```

## What is missing (intentional)

This is a scaffold, not a production chain. Missing pieces include:

- Persistent multi-store (IAVL/RocksDB)
- ABCI server integration with CometBFT
- Crypto/signatures and real auth/fee logic
- Governance, slashing, staking rewards
- Upgrade module, param store, evidence
- CLI, gRPC, REST APIs
- Mempool, block proposer, P2P networking

## Next steps

- Add a multi-store with per-module store keys
- Define protobuf/gRPC services for modules
- Implement auth/fee/signer validation in `DeliverTx`
- Add CometBFT ABCI glue
- Replace `MemStore` with a persistent store

## CometBFT integration

This scaffold includes a build-tagged ABCI adapter:

- `abci/cometbft.go` (requires `-tags cometbft` and the CometBFT module)
- `abci/stub.go` (default build; returns a clear error)

To enable:

1. Add the CometBFT dependency in `go.mod`.
2. Build with `-tags cometbft`.

Example (once dependency is added):

```bash
GOFLAGS="-tags cometbft" go build ./...
```

The adapter uses a JSON tx decoder in `cmd/node` for now; you can replace it with protobuf/amino later.

## Running the ABCI server

This repo now includes a runnable ABCI socket server (build-tagged):

```bash
GOFLAGS="-tags cometbft" go run ./cmd/node --socket unix:///tmp/newchain.sock --chain-id localnet
```

By default it expects JSON-encoded txs with a simple envelope:

```json
{
  "msgs": [
    {
      "type": "bank/send",
      "value": { "From": "alice", "To": "bob", "Amount": 10 }
    }
  ],
  "fee": 0,
  "signers": ["alice"]
}
```

To connect a CometBFT node, initialize and run CometBFT with its `proxy_app` set to the same socket.

## CometBFT bootstrap scripts

There are two helper scripts for initializing and running a local CometBFT node:

```bash
./scripts/cometbft-init.sh
./scripts/cometbft-run.sh
```

Environment overrides:

- `CHAIN_ID` (default `localnet`)
- `MONIKER` (default `node0`)
- `COMETBFT_HOME` (default `~/.cometbft`)
- `PROXY_APP` (default `unix:///tmp/newchain.sock`)

Example:

```bash
CHAIN_ID=localnet MONIKER=node0 PROXY_APP=unix:///tmp/newchain.sock ./scripts/cometbft-init.sh
./scripts/cometbft-run.sh
```

## One-command localnet

This script starts the ABCI server and CometBFT together, with logs in `go-sdk/logs`:

```bash
./scripts/run-localnet.sh
```

## Make target

```bash
make localnet
```

Stop both processes:

```bash
make stop
```

By default the runner writes PID files to `go-sdk/run/` and `make stop` uses those, so it won’t kill other CometBFT processes.

Check status:

```bash
make status
```
