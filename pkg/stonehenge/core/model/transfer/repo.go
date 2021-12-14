package transfer

import (
	"context"
	"stonehenge/pkg/stonehenge/core/types/id"
)

//go:generate moq -fmt goimports -out repo_mock.go . Repository:RepositoryMock

// Repository is the data access layer for the transfer entity
type Repository interface {
	// List gets all transfers existing
	List(ctx context.Context, filter Filter) ([]Transfer, error)

	// Get gets the transfer with the ID specified
	Get(ctx context.Context, id id.ID) (*Transfer, error)

	// Create creates a new transfer and returns its new id
	Create(ctx context.Context, account *Transfer) (*id.ID, error)
}
