package core

// Msg is a transaction message routed to a module.
type Msg interface {
	Type() string
	ValidateBasic() error
}
