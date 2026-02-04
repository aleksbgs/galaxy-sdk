package staking

import (
	"github.com/aleksbgs/galaxy-sdk/core"
	"github.com/aleksbgs/galaxy-sdk/types"
)

type MsgServer struct {
	keeper Keeper
}

func NewMsgServer(keeper Keeper) MsgServer {
	return MsgServer{keeper: keeper}
}

func (s MsgServer) Route() string { return "staking/delegate" }

func (s MsgServer) HandleMsg(_ core.Context, msg core.Msg) error {
	delegate, ok := msg.(MsgDelegate)
	if !ok {
		return types.ErrInvalidMsg
	}
	return s.keeper.Delegate(delegate.Delegator, delegate.Validator, delegate.Amount)
}
