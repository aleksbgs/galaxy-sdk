package staking

import (
	"encoding/json"
	"fmt"

	"github.com/aleksbgs/galaxy-sdk/modules/bank"
	"github.com/aleksbgs/galaxy-sdk/store"
	"github.com/aleksbgs/galaxy-sdk/types"
)

type Keeper struct {
	store      store.KVStore
	bankKeeper bank.Keeper
}

func NewKeeper(store store.KVStore, bankKeeper bank.Keeper) Keeper {
	return Keeper{store: store, bankKeeper: bankKeeper}
}

func (k Keeper) GetDelegation(delegator types.Address, validator string) int64 {
	bz := k.store.Get(delegationKey(delegator, validator))
	if bz == nil {
		return 0
	}
	var amt int64
	_ = json.Unmarshal(bz, &amt)
	return amt
}

func (k Keeper) SetDelegation(delegator types.Address, validator string, amount int64) {
	bz, _ := json.Marshal(amount)
	k.store.Set(delegationKey(delegator, validator), bz)
}

func (k Keeper) Delegate(delegator types.Address, validator string, amount int64) error {
	if err := k.bankKeeper.Send(delegator, types.Address("bonded_pool"), amount); err != nil {
		return fmt.Errorf("bond failed: %w", err)
	}
	current := k.GetDelegation(delegator, validator)
	k.SetDelegation(delegator, validator, current+amount)
	return nil
}

func delegationKey(delegator types.Address, validator string) []byte {
	return []byte("del/" + string(delegator) + "/" + validator)
}
