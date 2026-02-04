package core

import "fmt"

// MsgRegistry maps msg type strings to constructors.
type MsgRegistry struct {
	constructors map[string]func() Msg
}

func NewMsgRegistry() *MsgRegistry {
	return &MsgRegistry{constructors: make(map[string]func() Msg)}
}

func (r *MsgRegistry) Register(msgType string, ctor func() Msg) {
	if msgType == "" || ctor == nil {
		panic("invalid msg registry registration")
	}
	if _, exists := r.constructors[msgType]; exists {
		panic(fmt.Sprintf("msg type already registered: %s", msgType))
	}
	r.constructors[msgType] = ctor
}

func (r *MsgRegistry) New(msgType string) (Msg, bool) {
	ctor, ok := r.constructors[msgType]
	if !ok {
		return nil, false
	}
	return ctor(), true
}
