package core

import (
	"context"
	"time"
)

// Context carries request-scoped data across modules.
type Context struct {
	context.Context
	BlockHeight int64
	BlockTime   time.Time
	ChainID     string
	IsCheckTx   bool
	IsInitChain bool
}

func NewContext(base context.Context, chainID string, height int64, t time.Time) Context {
	return Context{
		Context:     base,
		ChainID:     chainID,
		BlockHeight: height,
		BlockTime:   t,
	}
}
