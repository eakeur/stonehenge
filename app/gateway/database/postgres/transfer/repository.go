package transfer

import (
	"stonehenge/app/core/entities/transfer"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) transfer.Repository {
	return &repository{
		db: db,
	}
}
