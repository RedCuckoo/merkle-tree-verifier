package main

import "errors"

var (
	ErrInternal       = errors.New("internal error")
	ErrInvalidRequest = errors.New("invalid request")
)
