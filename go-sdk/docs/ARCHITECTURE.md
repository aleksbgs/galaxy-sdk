# Architecture Notes

## Design goals

- Keep modules isolated with a small shared interface surface.
- Route messages by `Msg.Type()` to module-owned message servers.
- Separate state access via keepers.
- Keep the base app thin and composable.

## Core flows

1. `InitChain` calls each module's `InitGenesis`.
2. `BeginBlock` calls each module's `BeginBlock`.
3. `DeliverTx` validates `Msg`s and routes them to registered `MsgServer` handlers.
4. `EndBlock` calls each module's `EndBlock`.

## Suggested future enhancements

- Multi-store with per-module stores and versioned commits.
- Context that carries gas meter, events, and logger.
- Message router with middleware (auth, fee, ante handlers).
- Module manager with ordering/constraints for begin/end block.
- ABCI interface to plug into CometBFT.
