package account

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) Create(ctx context.Context, acc account.Account) (account.Account, error) {
	db, err := common.TransactionFrom(ctx)
	if err != nil {
		return account.Account{}, err
	}

	const script string = `
		insert into
			accounts (document, secret, name, balance)
		values 
			($1, $2, $3, $4)
		returning 
			id, external_id, created_at, updated_at
	`

	row := db.QueryRow(ctx, script, acc.Document, acc.Secret, acc.Name, acc.Balance)
	err = row.Scan(
		&acc.ID,
		&acc.ExternalID,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return account.Account{}, account.ErrAlreadyExist
		}

		return account.Account{}, account.ErrCreating
	}

	return acc, nil
}
