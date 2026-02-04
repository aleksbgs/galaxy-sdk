package auth

import "github.com/aleksbgs/galaxy-sdk/types"

type Account struct {
	Address  types.Address `json:"address"`
	Sequence uint64        `json:"sequence"`
}
