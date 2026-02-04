package abci

import "github.com/aleksbgs/galaxy-sdk/core"

// TxDecoder decodes raw tx bytes into a core.Tx.
type TxDecoder func([]byte) (core.Tx, error)
