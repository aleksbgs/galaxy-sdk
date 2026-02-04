package bank

import (
	"encoding/json"
	"fmt"

	"github.com/aleksbgs/galaxy-sdk/modules/auth"
	"github.com/aleksbgs/galaxy-sdk/store"
	"github.com/aleksbgs/galaxy-sdk/types"
)

type Keeper struct {
	store      store.KVStore
	authKeeper auth.Keeper
}

func NewKeeper(store store.KVStore, authKeeper auth.Keeper) Keeper {
	return Keeper{store: store, authKeeper: authKeeper}
}

func (k Keeper) GetBalance(addr types.Address) int64 {
	bz := k.store.Get(balanceKey(addr))
	if bz == nil {
		return 0
	}
	var bal int64
	_ = json.Unmarshal(bz, &bal)
	return bal
}

func (k Keeper) SetBalance(addr types.Address, amount int64) {
	bz, _ := json.Marshal(amount)
	k.store.Set(balanceKey(addr), bz)
}

func (k Keeper) Send(from, to types.Address, amount int64) error {
	k.authKeeper.EnsureAccount(from)
	k.authKeeper.EnsureAccount(to)
	fromBal := k.GetBalance(from)
	if fromBal < amount {
		return fmt.Errorf("insufficient funds")
	}
	k.SetBalance(from, fromBal-amount)
	k.SetBalance(to, k.GetBalance(to)+amount)
	k.authKeeper.IncrementSequence(from)
	return nil
}

func (k Keeper) Mint(to types.Address, amount int64) {
	k.authKeeper.EnsureAccount(to)
	k.SetBalance(to, k.GetBalance(to)+amount)
}

func balanceKey(addr types.Address) []byte {
	return []byte("bal/" + string(addr))
}
