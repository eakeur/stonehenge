package account

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"stonehenge/app/core/entities/account"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) account.Repository {
	return &repository{
		db: db,
	}
}
