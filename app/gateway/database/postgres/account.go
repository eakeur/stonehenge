package postgres

import (
	"context"
	"github.com/google/uuid"
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
		// TODO implement error
		return nil, err
	}
	defer ret.Close()
	accounts := make([]account.Account, 0)

	for ret.Next() {
		acc, err := parseAccount(ret)
		if err != nil {
			// TODO implement error
			return nil, err
		}
		accounts = append(accounts, *acc)
	}
	return accounts, nil
}

func (t *accountRepo) Get(ctx context.Context, id id.ID) (*account.Account, error) {
	const query string = "select * from accounts where id = $1"
	ret := t.db.QueryRow(ctx, query, id)
	acc, err := parseAccount(ret)
	if err != nil {
		// TODO implement error
		return nil, err
	}
	return acc, nil
}

func (t *accountRepo) GetBalance(ctx context.Context, id id.ID) (*currency.Currency, error) {
	const query string = "select balance from accounts where id = $1"
	ret := t.db.QueryRow(ctx, query, id)
	var balance currency.Currency
	if err := ret.Scan(&balance); err != nil {
		// TODO implement error
		return nil, err
	}
	return &balance, nil
}

func (t *accountRepo) Create(ctx context.Context, account *account.Account) (*id.ID, error) {
	db, found := t.tx.From(ctx)
	if !found {
		// TODO implement missing transactions
		return nil, nil
	}

	const script string = `
		insert into
			accounts (id, document, secret, name, balance)
		values 
			($1, $2, $3, $4, $5)
		returning 
			created_at, updated_at
	`

	account.Id = id.ID(uuid.New().String())
	row := db.QueryRow(ctx, script, account.Id, account.Document, account.Secret, account.Name, account.Balance)
	err := row.Scan(
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		//TODO implement me
		return nil, err
	}

	return &account.Id, nil
}

func (t *accountRepo) CheckExistence(ctx context.Context, document document.Document) error {
	const query string = "select count(*) as quantity from accounts where document = $1"
	ret := t.db.QueryRow(ctx, query, document)
	var quantity int
	if err := ret.Scan(&quantity); err != nil {
		// TODO implement error
		return err
	}
	if quantity > 0 {
		return account.ErrExists
	}
	return nil
}

func (t *accountRepo) UpdateBalance(ctx context.Context, id id.ID, balance currency.Currency) error {
	db, found := t.tx.From(ctx)
	if !found {
		// TODO implement missing transactions
		return nil
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
		//TODO implement me
		return nil
	}

	return nil
}

func parseAccount(row Scanner) (*account.Account, error) {
	acc := new(account.Account)
	err := row.Scan(&acc.Id, &acc.Name, &acc.Document, &acc.Balance, &acc.Secret, &acc.UpdatedAt, &acc.CreatedAt)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func NewAccountRepo(db *pgxpool.Pool, tx Transaction) account.Repository {
	return &accountRepo{
		tx, db,
	}
}