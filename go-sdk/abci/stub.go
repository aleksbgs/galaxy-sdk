//go:build !cometbft

package abci

import (
	"errors"

	"github.com/aleksbgs/galaxy-sdk/core"
)

// CometBFTApp is a stub when not built with the cometbft tag.
type CometBFTApp struct{}

// NewCometBFTApp returns an error unless built with the "cometbft" tag.
func NewCometBFTApp(_ core.App, _ TxDecoder) (*CometBFTApp, error) {
	return nil, errors.New("cometbft integration disabled: build with -tags cometbft")
}
