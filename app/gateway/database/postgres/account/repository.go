package account

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/gateway/database/postgres/common"
	"stonehenge/app/gateway/database/postgres/transaction"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	tx transaction.Transaction
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool, tx transaction.Transaction) account.Repository {
	return &repository{
		tx: tx, db: db,
	}
}

func (r *repository) StartOperation(ctx context.Context) (context.Context, error) {
	return r.tx.Begin(ctx)
}

func (r *repository) CommitOperation(ctx context.Context) error {
	if err := r.tx.Commit(ctx); err != nil {
		r.RollbackOperation(ctx)
		return err
	}
	return nil
}

func (r *repository) RollbackOperation(ctx context.Context) {
	if err := r.tx.Rollback(ctx); err != nil {
		return
	}
}

func parse(row common.Scanner, acc account.Account) (account.Account, error) {
	err := row.Scan(&acc.ID, &acc.Name, &acc.Document, &acc.Balance, &acc.Secret, &acc.UpdatedAt, &acc.CreatedAt)
	if err != nil {
		return acc, err
	}
	return acc, nil
}
