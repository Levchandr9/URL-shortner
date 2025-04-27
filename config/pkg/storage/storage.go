package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not fount")
	ErrURLExist    = errors.New("url exists")
)
