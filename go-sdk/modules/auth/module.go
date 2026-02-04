package auth

import "github.com/aleksbgs/galaxy-sdk/core"

type Module struct {
	Keeper Keeper
}

func NewModule(keeper Keeper) Module {
	return Module{Keeper: keeper}
}

func (m Module) Name() string { return "auth" }

func (m Module) RegisterServices(_ *core.ServiceRegistry) {
	// auth has no message handlers in this minimal framework
}

func (m Module) InitGenesis(_ core.Context, _ []byte) error { return nil }
func (m Module) BeginBlock(_ core.Context) error            { return nil }
func (m Module) EndBlock(_ core.Context) error              { return nil }
