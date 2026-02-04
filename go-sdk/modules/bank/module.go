package bank

import "github.com/aleksbgs/galaxy-sdk/core"

type Module struct {
	Keeper    Keeper
	msgServer MsgServer
}

func NewModule(keeper Keeper) Module {
	return Module{Keeper: keeper, msgServer: NewMsgServer(keeper)}
}

func (m Module) Name() string { return "bank" }

func (m Module) RegisterServices(reg *core.ServiceRegistry) {
	reg.Register(m.msgServer)
}

func (m Module) InitGenesis(_ core.Context, _ []byte) error { return nil }
func (m Module) BeginBlock(_ core.Context) error            { return nil }
func (m Module) EndBlock(_ core.Context) error              { return nil }
