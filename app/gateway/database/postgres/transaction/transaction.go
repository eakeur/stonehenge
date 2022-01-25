package transaction

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"stonehenge/app/core/entities/transaction"
)

type manager struct {
	db *pgxpool.Pool
}

// NewManager creates a transaction adapter object
func NewManager(db *pgxpool.Pool) transaction.Manager {
	return &manager{
		db: db,
	}
}
