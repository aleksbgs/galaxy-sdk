package staking

import (
	"errors"

	"github.com/aleksbgs/galaxy-sdk/core"
	"github.com/aleksbgs/galaxy-sdk/types"
)

// MsgDelegate bonds tokens to a validator.
type MsgDelegate struct {
	Delegator types.Address
	Validator string
	Amount    int64
}

func (m MsgDelegate) Type() string { return "staking/delegate" }

func (m MsgDelegate) ValidateBasic() error {
	if m.Delegator == "" || m.Validator == "" {
		return errors.New("missing delegator or validator")
	}
	if m.Amount <= 0 {
		return errors.New("amount must be positive")
	}
	return nil
}

var _ core.Msg = MsgDelegate{}
