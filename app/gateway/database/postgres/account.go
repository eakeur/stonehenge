package postgres

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"

	"github.com/jackc/pgx/v4/pgxpool"
)

type accountRepo struct {
	tx Transaction
	db *pgxpool.Pool
}

func (t *accountRepo) GetWithCPF(ctx context.Context, document document.Document) (account.Account, error) {
	const query string = "select * from accounts where document = $1"
	ret := t.db.QueryRow(ctx, query, document)
	acc := account.Account{}
	acc, err := parseAccount(ret, acc)
	if err != nil {
		return acc, account.ErrNotFound
	}
	return acc, nil
}

func (t *accountRepo) StartOperation(ctx context.Context) (context.Context, error) {
	return t.tx.Begin(ctx)
}

func (t *accountRepo) CommitOperation(ctx context.Context) error {
	if err := t.tx.Commit(ctx); err != nil {
		t.RollbackOperation(ctx)
		return err
	}
	return nil
}

func (t *accountRepo) RollbackOperation(ctx context.Context) {
	if err := t.tx.Rollback(ctx); err != nil {
		return
	}
	return
}

func (t *accountRepo) List(ctx context.Context, filter account.Filter) ([]account.Account, error) {
	query := "select * from accounts"
	args := make([]interface{}, 0)
	if filter.Name != "" {
		query = AppendCondition(query, "and", "name like ?")
		args = append(args, "%"+filter.Name+"%")
	}

	ret, err := t.db.Query(ctx, query, args...)
	if err != nil {
		return nil, account.ErrNotFound
	}
	defer ret.Close()
	accounts := make([]account.Account, 0)

	for ret.Next() {
		acc := account.Account{}
		acc, err := parseAccount(ret, acc)
		if err != nil {
			continue
		}
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

func (t *accountRepo) Get(ctx context.Context, id id.ExternalID) (account.Account, error) {
	const query string = "select * from accounts where external_id = $1"
	acc := account.Account{}
	ret := t.db.QueryRow(ctx, query, id)
	acc, err := parseAccount(ret, acc)
	if err != nil {
		return acc, account.ErrNotFound
	}
	return acc, nil
}

func (t *accountRepo) GetBalance(ctx context.Context, id id.ExternalID) (currency.Currency, error) {
	const query string = "select balance from accounts where external_id = $1"
	ret := t.db.QueryRow(ctx, query, id)
	var balance currency.Currency
	if err := ret.Scan(&balance); err != nil {
		return 0, account.ErrNotFound
	}
	return balance, nil
}

func (t *accountRepo) Create(ctx context.Context, acc *account.Account) (id.ExternalID, error) {
	db, found := t.tx.From(ctx)
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

func (t *accountRepo) UpdateBalance(ctx context.Context, id id.ExternalID, balance currency.Currency) error {
	db, found := t.tx.From(ctx)
	if !found {
		return account.ErrCreating
	}

	const script string = `
		update
			accounts
		set
			balance = $1
		where
			id = $2
	`
	_, err := db.Exec(ctx, script, balance, id)
	if err != nil {
		return account.ErrCreating
	}

	return nil
}

func parseAccount(row Scanner, acc account.Account) (account.Account, error) {
	err := row.Scan(&acc.ID, &acc.Name, &acc.Document, &acc.Balance, &acc.Secret, &acc.UpdatedAt, &acc.CreatedAt)
	if err != nil {
		return acc, err
	}
	return acc, nil
}

func NewAccountRepo(db *pgxpool.Pool, tx Transaction) account.Repository {
	return &accountRepo{
		tx, db,
	}
}
