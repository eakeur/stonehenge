package transfer

import (
	"context"
)

//go:generate moq -fmt goimports -out repo_mock.go . Repository:RepositoryMock

// Repository is the data access layer for the transfer entity
type Repository interface {
	// List gets all transfers existing
	List(ctx context.Context, filter Filter) ([]Transfer, error)

	// Create creates a new transfer and returns its new id
	Create(ctx context.Context, transfer Transfer) (Transfer, error)
}
