package access

import "errors"

var (
	ErrNoAccessInContext = errors.New("there is no access found in the context informed")

	ErrTokenInvalidOrExpired = errors.New("the token informed is invalid or has expired")

	ErrTokenFailedCreation = errors.New("the token could not be fetched with the signing key informed")
)
