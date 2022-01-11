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
	List(context.Context, Filter) ([]Account, error)

	// GetByExternalID gets the account with the ID specified
	GetByExternalID(context.Context, id.ExternalID) (Account, error)

	// GetWithCPF gets the account with the document specified
	GetWithCPF(context.Context, document.Document) (Account, error)

	// GetBalance gets the balance with the ID specified
	GetBalance(context.Context, id.ExternalID) (currency.Currency, error)

	// Create creates a new account and returns its new id
	Create(context.Context, Account) (Account, error)

	// UpdateBalance replaces the balance of the account with the id provided
	UpdateBalance(context.Context, id.ExternalID, currency.Currency) error
}
