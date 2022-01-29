package transfer

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"stonehenge/app/core/entities/transfer"
)

type repository struct {
	db *pgxpool.Pool
	logger zerolog.Logger
}

func NewRepository(db *pgxpool.Pool, lg zerolog.Logger) transfer.Repository {
	return &repository{
		db: db,
		logger: lg,
	}
}
