package common

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type key int

const (
	TXContextKey key = 71
)

// TransactionFrom looks up for a pgx.Tx object in this context and retrieves it
func TransactionFrom(ctx context.Context) (pgx.Tx, bool) {
	v, ok := ctx.Value(TXContextKey).(pgx.Tx)
	return v, ok
}
