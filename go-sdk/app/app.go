package app

import (
	"github.com/aleksbgs/galaxy-sdk/base"
	"github.com/aleksbgs/galaxy-sdk/modules/auth"
	"github.com/aleksbgs/galaxy-sdk/modules/bank"
	"github.com/aleksbgs/galaxy-sdk/modules/staking"
	"github.com/aleksbgs/galaxy-sdk/store"
)

// App wires all modules together.
type App struct {
	*base.BaseApp
	AuthKeeper    auth.Keeper
	BankKeeper    bank.Keeper
	StakingKeeper staking.Keeper
}

func New(chainID string) *App {
	// In a real system you'd use IAVL/rocksdb/multistore with separate keys per module.
	mem := store.NewMemStore()

	authKeeper := auth.NewKeeper(mem)
	bankKeeper := bank.NewKeeper(mem, authKeeper)
	stakingKeeper := staking.NewKeeper(mem, bankKeeper)

	authModule := auth.NewModule(authKeeper)
	bankModule := bank.NewModule(bankKeeper)
	stakingModule := staking.NewModule(stakingKeeper)

	baseApp := base.NewBaseApp("GalaxySDK", chainID, authModule, bankModule, stakingModule)

	return &App{
		BaseApp:       baseApp,
		AuthKeeper:    authKeeper,
		BankKeeper:    bankKeeper,
		StakingKeeper: stakingKeeper,
	}
}
