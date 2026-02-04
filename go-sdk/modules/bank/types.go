package bank

import (
	"errors"

	"github.com/aleksbgs/galaxy-sdk/core"
	"github.com/aleksbgs/galaxy-sdk/types"
)

// MsgSend moves tokens between accounts.
type MsgSend struct {
	From   types.Address
	To     types.Address
	Amount int64
}

func (m MsgSend) Type() string { return "bank/send" }

func (m MsgSend) ValidateBasic() error {
	if m.From == "" || m.To == "" {
		return errors.New("missing address")
	}
	if m.Amount <= 0 {
		return errors.New("amount must be positive")
	}
	return nil
}

var _ core.Msg = MsgSend{}
