package cli

import "errors"

var (
	ErrInternal       = errors.New("internal error")
	ErrInvalidRequest = errors.New("invalid request")
)
