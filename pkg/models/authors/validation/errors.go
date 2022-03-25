package validation

import "errors"

var (
	ErrInvalidID      = errors.New("Invalid author id")
	ErrInvalidName    = errors.New("Invalid author name")
	ErrInvalidSurname = errors.New("Invalid author surname")
)
