package transaction

import (
	"context"
	"stonehenge/app/gateway/database/postgres/common"
)

func (t *pgxTransaction) Rollback(ctx context.Context) {
	tx, err := common.TransactionFrom(ctx)
	if err != nil {
		return
	}

	if err := tx.Rollback(ctx); err != nil {
		return
	}
}
