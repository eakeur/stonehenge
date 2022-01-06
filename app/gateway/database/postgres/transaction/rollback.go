package transaction

import (
	"context"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/gateway/database/postgres/common"
)

func (t *pgxTransaction) Rollback(ctx context.Context) error {
	tx, ok := common.TransactionFrom(ctx)
	if !ok {
		return transaction.ErrNoTransaction
	}

	if err := tx.Rollback(ctx); err != nil {
		return err
	}

	return nil
}
