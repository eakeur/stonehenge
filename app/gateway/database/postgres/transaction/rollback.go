package transaction

import (
	"context"
	"stonehenge/app/gateway/database/postgres/common"
)

func (t *pgxTransaction) Rollback(ctx context.Context) {
	tx, ok := common.TransactionFrom(ctx)
	if !ok {
		return
	}

	if err := tx.Rollback(ctx); err != nil {
		return
	}
}
