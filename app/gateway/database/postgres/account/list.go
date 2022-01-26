package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/database/postgres/common"
)

func (r *repository) List(ctx context.Context, filter account.Filter) ([]account.Account, error) {
	const operation = "Repositories.Account.List"
	query := `select 
		id, 
		external_id, 
		name, 
		document, 
		balance, 
		secret, 
		updated_at, 
		created_at 
	from 
		accounts`
	args := make([]interface{}, 0)
	if filter.Name != "" {
		r.logger.Trace(ctx, operation, "searching accounts with filter name")
		query = common.AppendCondition(query, "and", "name like ?", 1)
		args = append(args, "%"+filter.Name+"%")
	}

	query += "\n order by id"

	ret, err := r.db.Query(ctx, query, args...)
	if err != nil {
		r.logger.Error(ctx, operation, err.Error())
		return []account.Account{}, account.ErrFetching
	}
	defer ret.Close()

	accounts := make([]account.Account, 0)
	for ret.Next() {
		acc := account.Account{}
		err = ret.Scan(
			&acc.ID,
			&acc.ExternalID,
			&acc.Name,
			&acc.Document,
			&acc.Balance,
			&acc.Secret,
			&acc.UpdatedAt,
			&acc.CreatedAt)
		accounts = append(accounts, acc)
	}

	r.logger.Trace(ctx, operation, "finished process successfully")
	return accounts, nil
}
