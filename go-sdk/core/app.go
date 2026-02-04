package core

// App defines the high-level blockchain application interface.
type App interface {
	Name() string
	InitChain(ctx Context, appState []byte) error
	BeginBlock(ctx Context) error
	DeliverTx(ctx Context, tx Tx) error
	EndBlock(ctx Context) error
}
