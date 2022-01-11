package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) Create(ctx context.Context, acc account.Account) (account.Account, error) {
	db, found := common.TransactionFrom(ctx)
	if !found {
		return acc, account.ErrCreating
	}

	const script string = `
		insert into
			accounts (id, document, secret, name, balance)
		values 
			($1, $2, $3, $4, $5)
		returning 
			id, external_id, created_at, updated_at
	`

	row := db.QueryRow(ctx, script, acc.ID, acc.Document, acc.Secret, acc.Name, acc.Balance)
	err := row.Scan(
		&acc.ID,
		&acc.ExternalID,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		return acc, account.ErrCreating
	}

	return acc, nil
}
