package core

// Tx is a collection of messages with signatures and fee metadata.
type Tx interface {
	GetMsgs() []Msg
	GetFee() int64
	GetSigners() []string
}
