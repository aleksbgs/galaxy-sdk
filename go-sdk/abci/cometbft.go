//go:build cometbft

package abci

import (
	"context"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/aleksbgs/galaxy-sdk/core"
)

// CometBFTApp adapts the SDK App to CometBFT ABCI++.
type CometBFTApp struct {
	app       core.App
	txDecoder TxDecoder
}

func NewCometBFTApp(app core.App, decoder TxDecoder) (*CometBFTApp, error) {
	return &CometBFTApp{app: app, txDecoder: decoder}, nil
}

func (c *CometBFTApp) Info(_ context.Context, _ *abci.RequestInfo) (*abci.ResponseInfo, error) {
	return &abci.ResponseInfo{AppVersion: "0.0.1"}, nil
}

func (c *CometBFTApp) Query(_ context.Context, _ *abci.RequestQuery) (*abci.ResponseQuery, error) {
	return &abci.ResponseQuery{Code: 1, Log: "query not implemented"}, nil
}

func (c *CometBFTApp) CheckTx(_ context.Context, req *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	tx, err := c.decodeTx(req.Tx)
	if err != nil {
		return &abci.ResponseCheckTx{Code: 1, Log: err.Error()}, nil
	}
	for _, msg := range tx.GetMsgs() {
		if err := msg.ValidateBasic(); err != nil {
			return &abci.ResponseCheckTx{Code: 1, Log: err.Error()}, nil
		}
	}
	return &abci.ResponseCheckTx{Code: 0}, nil
}

func (c *CometBFTApp) InitChain(_ context.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	ctx := core.NewContext(context.Background(), req.ChainId, 0, req.Time)
	ctx.IsInitChain = true
	if err := c.app.InitChain(ctx, req.AppStateBytes); err != nil {
		return nil, err
	}
	return &abci.ResponseInitChain{}, nil
}

func (c *CometBFTApp) PrepareProposal(_ context.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	// Accept and return all txs as-is for now.
	return &abci.ResponsePrepareProposal{Txs: req.Txs}, nil
}

func (c *CometBFTApp) ProcessProposal(_ context.Context, _ *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}

func (c *CometBFTApp) FinalizeBlock(_ context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	ctx := core.NewContext(context.Background(), req.ChainId, req.Height, req.Time)
	if err := c.app.BeginBlock(ctx); err != nil {
		return nil, err
	}

	results := make([]*abci.ExecTxResult, len(req.Txs))
	for i, raw := range req.Txs {
		tx, err := c.decodeTx(raw)
		if err != nil {
			results[i] = &abci.ExecTxResult{Code: 1, Log: err.Error()}
			continue
		}
		if err := c.app.DeliverTx(ctx, tx); err != nil {
			results[i] = &abci.ExecTxResult{Code: 1, Log: err.Error()}
			continue
		}
		results[i] = &abci.ExecTxResult{Code: 0}
	}

	if err := c.app.EndBlock(ctx); err != nil {
		return nil, err
	}

	return &abci.ResponseFinalizeBlock{TxResults: results}, nil
}

func (c *CometBFTApp) Commit(_ context.Context, _ *abci.RequestCommit) (*abci.ResponseCommit, error) {
	return &abci.ResponseCommit{}, nil
}

func (c *CometBFTApp) ListSnapshots(_ context.Context, _ *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	return &abci.ResponseListSnapshots{}, nil
}

func (c *CometBFTApp) OfferSnapshot(_ context.Context, _ *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	return &abci.ResponseOfferSnapshot{Result: abci.ResponseOfferSnapshot_REJECT}, nil
}

func (c *CometBFTApp) LoadSnapshotChunk(_ context.Context, _ *abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	return &abci.ResponseLoadSnapshotChunk{}, nil
}

func (c *CometBFTApp) ApplySnapshotChunk(_ context.Context, _ *abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	return &abci.ResponseApplySnapshotChunk{Result: abci.ResponseApplySnapshotChunk_REJECT}, nil
}

func (c *CometBFTApp) ExtendVote(_ context.Context, _ *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	return &abci.ResponseExtendVote{}, nil
}

func (c *CometBFTApp) VerifyVoteExtension(_ context.Context, _ *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
	return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
}

func (c *CometBFTApp) decodeTx(raw []byte) (core.Tx, error) {
	if c.txDecoder == nil {
		return nil, fmt.Errorf("tx decoder not configured")
	}
	return c.txDecoder(raw)
}

var _ abci.Application = (*CometBFTApp)(nil)
