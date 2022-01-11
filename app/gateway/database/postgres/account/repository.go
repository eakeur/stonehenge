package account

import (
	"stonehenge/app/core/entities/account"
	"stonehenge/app/gateway/database/postgres/common"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) account.Repository {
	return &repository{
		db: db,
	}
}

func parse(row common.Scanner, acc account.Account) (account.Account, error) {
	err := row.Scan(&acc.ID, &acc.ExternalID, &acc.Name, &acc.Document, &acc.Balance, &acc.Secret, &acc.UpdatedAt, &acc.CreatedAt)
	if err != nil {
		return acc, err
	}
	return acc, nil
}
