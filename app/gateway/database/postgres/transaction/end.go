package transaction

import (
	"context"
	"stonehenge/app/gateway/database/postgres/common"
)

func (t manager) End(ctx context.Context) {
	tx, err := common.TransactionFrom(ctx)
	if err != nil {
		return
	}
	err = tx.Commit(ctx)
	if err != nil {
		err = tx.Rollback(ctx)
		if err != nil {
			return
		}
	}
}
