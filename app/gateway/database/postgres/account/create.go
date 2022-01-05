package account

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/types/id"
)

func (r *repository) Create(ctx context.Context, acc *account.Account) (id.ExternalID, error) {
	db, found := r.tx.From(ctx)
	if !found {
		return id.New(), account.ErrCreating
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
		return id.New(), account.ErrCreating
	}

	return acc.ExternalID, nil
}
