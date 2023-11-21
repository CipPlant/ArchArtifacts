package model

import "errors"

var (
	ErrNoResults   = errors.New("no such results")
	ErrNoSuchPhoto = errors.New("no such photo")
)
