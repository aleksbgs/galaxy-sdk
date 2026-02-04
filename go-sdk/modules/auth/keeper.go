package auth

import (
	"encoding/json"

	"github.com/aleksbgs/galaxy-sdk/store"
	"github.com/aleksbgs/galaxy-sdk/types"
)

type Keeper struct {
	store store.KVStore
}

func NewKeeper(store store.KVStore) Keeper {
	return Keeper{store: store}
}

func (k Keeper) GetAccount(addr types.Address) (Account, bool) {
	bz := k.store.Get(accountKey(addr))
	if bz == nil {
		return Account{}, false
	}
	var acct Account
	if err := json.Unmarshal(bz, &acct); err != nil {
		return Account{}, false
	}
	return acct, true
}

func (k Keeper) SetAccount(acct Account) {
	bz, _ := json.Marshal(acct)
	k.store.Set(accountKey(acct.Address), bz)
}

func (k Keeper) EnsureAccount(addr types.Address) Account {
	acct, ok := k.GetAccount(addr)
	if ok {
		return acct
	}
	acct = Account{Address: addr, Sequence: 0}
	k.SetAccount(acct)
	return acct
}

func (k Keeper) IncrementSequence(addr types.Address) {
	acct := k.EnsureAccount(addr)
	acct.Sequence++
	k.SetAccount(acct)
}

func accountKey(addr types.Address) []byte {
	return []byte("acct/" + string(addr))
}
