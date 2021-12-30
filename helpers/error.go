package helpers

import "errors"

var (
	ErrNotFound      = errors.New("data not found")
	ErrDbServer      = errors.New("db server error")
	ErrEmptyUrlParam = errors.New("Empty url param")
	ErrIdNotFound    = errors.New("Id not found")
)
