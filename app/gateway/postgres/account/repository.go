package account

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"stonehenge/app/core/entities/account"
)

type repository struct {
	db     *pgxpool.Pool
	logger zerolog.Logger
}

func NewRepository(db *pgxpool.Pool, lg zerolog.Logger) account.Repository {
	return &repository{
		db: db,
		logger: lg,
	}
}
