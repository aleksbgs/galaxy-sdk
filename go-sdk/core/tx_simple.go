package core

// SimpleTx is a minimal Tx implementation for prototyping.
type SimpleTx struct {
	Msgs    []Msg
	Fee     int64
	Signers []string
}

func (s SimpleTx) GetMsgs() []Msg       { return s.Msgs }
func (s SimpleTx) GetFee() int64        { return s.Fee }
func (s SimpleTx) GetSigners() []string { return s.Signers }
