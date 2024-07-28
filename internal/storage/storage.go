package storage

import "errors"

var (
	ErrURLNotFound = errors.New("URL not found")
	ERRUrlExists   = errors.New("URL already exists")
)
