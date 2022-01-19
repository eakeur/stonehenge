package id

import "errors"

var (
	// ErrInvalidID throws when the identification provided is an invalid form of ID
	ErrInvalidID = errors.New("the account id provided is invalid. please, try again with a valid one")
)
