package server

import (
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	account_repo "stonehenge/app/gateway/database/postgres/account"
	transfer_repo "stonehenge/app/gateway/database/postgres/transfer"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RepositoryWrapper struct {
	Account  account.Repository
	Transfer transfer.Repository
}

func NewPostgresRepositoryWrapper(db *pgxpool.Pool) *RepositoryWrapper {
	return &RepositoryWrapper{
		Account:  account_repo.NewRepository(db),
		Transfer: transfer_repo.NewRepository(db),
	}
}
