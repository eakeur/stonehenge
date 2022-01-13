package common

import (
	"context"
	"stonehenge/app/core/entities/transaction"

	"github.com/jackc/pgx/v4"
)

type key int

const (
	TXContextKey key = 71
)

// TransactionFrom looks up for a pgx.Tx object in this context and retrieves it
func TransactionFrom(ctx context.Context) (pgx.Tx, error) {
	v, ok := ctx.Value(TXContextKey).(pgx.Tx)
	if !ok {
		return v, transaction.ErrNoTransaction
	}
	return v, nil
}
