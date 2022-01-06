package transaction

import "context"

// Transaction is an object that encapsulates many storage accesses into one transaction
type Transaction interface {
	// Begin starts a transaction and stores an object to it in this context
	Begin(ctx context.Context) (context.Context, error)

	// Commit commits a transaction in this context
	Commit(ctx context.Context) error

	// Rollback rollbacks a transaction in this context
	Rollback(ctx context.Context) error
}
