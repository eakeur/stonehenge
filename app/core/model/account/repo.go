package account

import (
	"context"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
)

//go:generate moq -fmt goimports -out repo_mock.go . Repository:RepositoryMock

// Repository is the data access layer for the account entity
type Repository interface {
	// List gets all accounts existing
	List(ctx context.Context, filter Filter) ([]Account, error)

	// Get gets the account with the ID specified
	Get(ctx context.Context, id id.ID) (*Account, error)

	// GetBalance gets the balance with the ID specified
	GetBalance(ctx context.Context, id id.ID) (*currency.Currency, error)

	// Create creates a new account and returns its new id
	Create(ctx context.Context, account *Account) (*id.ID, error)

	// CheckExistence checks if a document is registered to any account and throws an error if it is
	CheckExistence(ctx context.Context, document document.Document) error

	// UpdateBalance replaces the balance of the account with the id provided
	UpdateBalance(ctx context.Context, id id.ID, balance currency.Currency) error

	// StartOperation creates a transaction in this context
	StartOperation(ctx context.Context) (context.Context, error)

	// CommitOperation finishes successfully a transaction in this context or rollbacks it
	// in case of failing
	CommitOperation(ctx context.Context) error

	// RollbackOperation finishes unsuccessfully a transaction in this context
	RollbackOperation(ctx context.Context)
}
