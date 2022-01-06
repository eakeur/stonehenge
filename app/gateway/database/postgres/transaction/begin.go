package transaction

import (
	"context"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/gateway/database/postgres/common"
)

func (t *pgxTransaction) Begin(ctx context.Context) (context.Context, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return nil, transaction.ErrBeginTransaction
	}
	return context.WithValue(ctx, common.TXContextKey, tx), nil
}
