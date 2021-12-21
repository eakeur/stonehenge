package accounts

import "errors"

var (

	// ErrTokenGeneration throws when an unexpected error occurs at the token generation of a logged-in user
	ErrTokenGeneration = errors.New("an error occurred while generating your access token")
)
