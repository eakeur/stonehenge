package account

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
)

func (r *repository) GetByExternalID(ctx context.Context, id id.ExternalID) (account.Account, error) {
	const query string = `select 
		id, 
		external_id, 
		name, 
		document, 
		balance, 
		secret, 
		updated_at, 
		created_at 
	from 
		accounts
	where external_id = $1`

	acc := account.Account{}

	ret := r.db.QueryRow(ctx, query, id)
	err := ret.Scan(
		&acc.ID,
		&acc.ExternalID,
		&acc.Name,
		&acc.Document,
		&acc.Balance,
		&acc.Secret,
		&acc.UpdatedAt,
		&acc.CreatedAt)

	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return account.Account{}, account.ErrNotFound
		}
		return account.Account{}, account.ErrFetching

	}
	return acc, nil
}
