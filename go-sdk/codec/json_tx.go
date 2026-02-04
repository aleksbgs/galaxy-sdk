package codec

import (
	"encoding/json"
	"fmt"

	"github.com/aleksbgs/galaxy-sdk/core"
)

type rawMsg struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type rawTx struct {
	Msgs    []rawMsg `json:"msgs"`
	Fee     int64    `json:"fee"`
	Signers []string `json:"signers"`
}

// NewJSONTxDecoder builds a decoder for SimpleTx using a msg registry.
func NewJSONTxDecoder(reg *core.MsgRegistry) func([]byte) (core.Tx, error) {
	return func(bz []byte) (core.Tx, error) {
		var raw rawTx
		if err := json.Unmarshal(bz, &raw); err != nil {
			return nil, fmt.Errorf("decode tx: %w", err)
		}
		msgs := make([]core.Msg, 0, len(raw.Msgs))
		for _, rm := range raw.Msgs {
			msg, ok := reg.New(rm.Type)
			if !ok {
				return nil, fmt.Errorf("unknown msg type: %s", rm.Type)
			}
			if err := json.Unmarshal(rm.Value, msg); err != nil {
				return nil, fmt.Errorf("decode msg %s: %w", rm.Type, err)
			}
			msgs = append(msgs, msg)
		}
		return core.SimpleTx{Msgs: msgs, Fee: raw.Fee, Signers: raw.Signers}, nil
	}
}
