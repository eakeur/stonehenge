package transaction

import (
	"context"
	"stonehenge/app/gateway/postgres/common"
)

func (t *manager) Rollback(ctx context.Context) {
	tx, err := common.TransactionFrom(ctx)
	if err != nil {
		return
	}

	if err := tx.Rollback(ctx); err != nil {
		return
	}
}
