package account

import (
	"context"
	"stonehenge/pkg/stonehenge/core/types/currency"
	"stonehenge/pkg/stonehenge/core/types/document"
	"stonehenge/pkg/stonehenge/core/types/id"
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

	// Update replaces the values of all fields of the account with the id provided
	Update(ctx context.Context, id id.ID, account *Account) error

	// UpdateBalance replaces the balance of the account with the id provided
	UpdateBalance(ctx context.Context, id id.ID, balance currency.Currency) error

	// Remove removes an account from the database
	Remove(ctx context.Context, id id.ID) error
}
