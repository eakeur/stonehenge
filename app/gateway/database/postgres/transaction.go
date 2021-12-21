package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type key int

const (
	TXContextKey key = 71
)

// Transaction is an object that wraps the main actions of a standard pgx.TX object
type Transaction interface {

	// From looks up for a pgx.Tx object in this context and retrieves it
	From(ctx context.Context) (pgx.Tx, bool)

	// Begin starts a transaction and stores an object to it in this context
	Begin(ctx context.Context) (context.Context, error)

	// Commit commits a transaction in this context
	Commit(ctx context.Context) error

	// Rollback rollbacks a transaction in this context
	Rollback(ctx context.Context) error
}

type transaction struct {
	db *pgxpool.Pool
}

func (t *transaction) From(ctx context.Context) (pgx.Tx, bool) {
	v, ok := ctx.Value(TXContextKey).(pgx.Tx)
	return v, ok
}

func (t *transaction) Begin(ctx context.Context) (context.Context, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return nil, ErrBeginTransaction
	}
	return context.WithValue(ctx, TXContextKey, tx), nil
}

func (t *transaction) Commit(ctx context.Context) error {
	tx, ok := t.From(ctx)
	if !ok {
		return ErrNoTransaction
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (t *transaction) Rollback(ctx context.Context) error {
	tx, ok := t.From(ctx)
	if !ok {
		return ErrNoTransaction
	}

	if err := tx.Rollback(ctx); err != nil {
		return err
	}

	return nil
}

// NewTransactionAdapter creates a transaction adapter object
func NewTransactionAdapter(db *pgxpool.Pool) Transaction {
	return &transaction{
		db: db,
	}
}
