package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) UpdateBalance(ctx context.Context, id id.External, balance currency.Currency) error {
	const operation = "Repositories.Account.UpdateBalance"
	db, err := common.TransactionFrom(ctx)
	if err != nil {
		r.logger.Error(ctx, operation, err.Error())
		return err
	}

	const script string = `
		update
			accounts
		set
			balance = $1
		where
			external_id = $2
	`
	res, err := db.Exec(ctx, script, balance, id)
	if err != nil {
		r.logger.Error(ctx, operation, err.Error())
		return account.ErrUpdating
	}

	if res.RowsAffected() != 1 {
		r.logger.Error(ctx, operation, "no row was changed")
		return account.ErrNotFound
	}
	r.logger.Trace(ctx, operation, "finished process successfully")
	return nil
}
