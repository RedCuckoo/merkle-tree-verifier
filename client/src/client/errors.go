package client

import "errors"

var (
	ErrTreeCreated = errors.New(
		"merkle tree already created or \"generate\" has been already called, run \"reset\" to start over",
	)
	ErrGeneratedNotCalled = errors.New("no files, call \"generate\" first")
	ErrServerRootMismatch = errors.New(
		"merkle root tree returned from the server differs from calculated",
	)
)
