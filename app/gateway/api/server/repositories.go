package server

import (
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/gateway/database/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RepositoryWrapper struct {
	Account  account.Repository
	Transfer transfer.Repository
}

func NewPostgresRepositoryWrapper(db *pgxpool.Pool, tx postgres.Transaction) *RepositoryWrapper {
	return &RepositoryWrapper{
		Account:  postgres.NewAccountRepo(db, tx),
		Transfer: postgres.NewTransferRepo(db, tx),
	}
}
