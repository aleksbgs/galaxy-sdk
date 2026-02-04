package bank

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

func (s MsgServer) Route() string { return "bank/send" }

func (s MsgServer) HandleMsg(_ core.Context, msg core.Msg) error {
	send, ok := msg.(MsgSend)
	if !ok {
		return types.ErrInvalidMsg
	}
	return s.keeper.Send(send.From, send.To, send.Amount)
}
