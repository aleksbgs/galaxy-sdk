package types

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidMsg   = errors.New("invalid message")
	ErrNotFound     = errors.New("not found")
)

type Address string
