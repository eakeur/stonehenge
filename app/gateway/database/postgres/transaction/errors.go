package transaction

import "errors"

var (

	// ErrNoTransaction throws when no transaction can be found in a given context
	ErrNoTransaction = errors.New("could not find any transaction in this context")

	// ErrBeginTransaction occurs when it fails to begin a transaction in a given context
	ErrBeginTransaction = errors.New("could not begin transaction")
)
