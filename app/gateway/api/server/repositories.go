package server

import (
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/model/transfer"
	account_repo "stonehenge/app/gateway/database/postgres/account"
	"stonehenge/app/gateway/database/postgres/transaction"
	transfer_repo "stonehenge/app/gateway/database/postgres/transfer"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RepositoryWrapper struct {
	Account  account.Repository
	Transfer transfer.Repository
}

func NewPostgresRepositoryWrapper(db *pgxpool.Pool, tx transaction.Transaction) *RepositoryWrapper {
	return &RepositoryWrapper{
		Account:  account_repo.NewRepository(db, tx),
		Transfer: transfer_repo.NewRepository(db, tx),
	}
}
