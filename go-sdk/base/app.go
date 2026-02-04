package base

import (
	"context"
	"fmt"
	"time"

	"github.com/aleksbgs/galaxy-sdk/core"
)

// BaseApp wires modules and routes messages.
type BaseApp struct {
	name     string
	modules  []core.AppModule
	services *core.ServiceRegistry
	chainID  string
	height   int64
}

func NewBaseApp(name, chainID string, modules ...core.AppModule) *BaseApp {
	services := core.NewServiceRegistry()
	for _, mod := range modules {
		mod.RegisterServices(services)
	}
	return &BaseApp{
		name:     name,
		modules:  modules,
		services: services,
		chainID:  chainID,
	}
}

func (b *BaseApp) Name() string { return b.name }

func (b *BaseApp) InitChain(ctx core.Context, appState []byte) error {
	ctx.IsInitChain = true
	for _, mod := range b.modules {
		if err := mod.InitGenesis(ctx, appState); err != nil {
			return fmt.Errorf("init genesis %s: %w", mod.Name(), err)
		}
	}
	return nil
}

func (b *BaseApp) BeginBlock(ctx core.Context) error {
	for _, mod := range b.modules {
		if err := mod.BeginBlock(ctx); err != nil {
			return fmt.Errorf("begin block %s: %w", mod.Name(), err)
		}
	}
	return nil
}

func (b *BaseApp) DeliverTx(ctx core.Context, tx core.Tx) error {
	for _, msg := range tx.GetMsgs() {
		if err := msg.ValidateBasic(); err != nil {
			return err
		}
		server, ok := b.services.Get(msg.Type())
		if !ok {
			return fmt.Errorf("no msg server for %s", msg.Type())
		}
		if err := server.HandleMsg(ctx, msg); err != nil {
			return err
		}
	}
	return nil
}

func (b *BaseApp) EndBlock(ctx core.Context) error {
	for _, mod := range b.modules {
		if err := mod.EndBlock(ctx); err != nil {
			return fmt.Errorf("end block %s: %w", mod.Name(), err)
		}
	}
	return nil
}

// Tick simulates a block for testing.
func (b *BaseApp) Tick(t time.Time, txs ...core.Tx) error {
	b.height++
	ctx := core.NewContext(context.Background(), b.chainID, b.height, t)
	if err := b.BeginBlock(ctx); err != nil {
		return err
	}
	for _, tx := range txs {
		if err := b.DeliverTx(ctx, tx); err != nil {
			return err
		}
	}
	return b.EndBlock(ctx)
}
