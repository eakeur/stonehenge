package account

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/erring"
)

func (r *repository) GetWithCPF(ctx context.Context, doc document.Document) (account.Account, error) {
	const operation = "Repositories.Account.GetWithCPF"
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
	where 
		document = $1`

	ret := r.db.QueryRow(ctx, query, doc)
	acc := account.Account{}
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
			return account.Account{}, erring.Wrap(account.ErrNotFound, operation)
		}

		return account.Account{}, erring.Wrap(err, operation)

	}
	return acc, nil
}
