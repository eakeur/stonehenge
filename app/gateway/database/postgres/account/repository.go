package account

import (
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db     *pgxpool.Pool
	logger logger.Logger
}

func NewRepository(db *pgxpool.Pool) account.Repository {
	return &repository{
		db: db,
	}
}
