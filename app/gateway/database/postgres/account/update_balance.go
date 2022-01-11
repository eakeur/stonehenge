package account

import (
	"context"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) UpdateBalance(ctx context.Context, id id.ExternalID, balance currency.Currency) error {
	db, err := common.TransactionFrom(ctx)
	if err != nil {
		return err
	}

	const script string = `
		update
			accounts
		set
			balance = $1
		where
			id = $2
	`
	_, err = db.Exec(ctx, script, balance, id)
	if err != nil {
		return err
	}

	return nil
}
