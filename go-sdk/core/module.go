package core

// AppModule defines a module's lifecycle and message handling.
type AppModule interface {
	Name() string
	RegisterServices(svc *ServiceRegistry)
	InitGenesis(ctx Context, data []byte) error
	BeginBlock(ctx Context) error
	EndBlock(ctx Context) error
}

// MsgServer handles messages for a module.
type MsgServer interface {
	Route() string
	HandleMsg(ctx Context, msg Msg) error
}
