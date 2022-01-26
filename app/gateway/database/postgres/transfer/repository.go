package transfer

import (
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
	logger logger.Logger
}

func NewRepository(db *pgxpool.Pool, lg logger.Logger) transfer.Repository {
	return &repository{
		db: db,
		logger: lg,
	}
}
