package account

import (
	"context"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
)

// Repository is the data access layer for the account entity
type Repository interface {

	// List gets all accounts existing that respects a given filter. It may return ErrFetching if an untracked error happens
	List(context.Context, Filter) ([]Account, error)

	// GetByExternalID gets the account with the id.External specified. It may return ErrNotFound
	// if no account with the id.External is found, ErrInvalidID if the id.External provided is corrupt
	// or ErrFetching if an untracked error happens
	GetByExternalID(context.Context, id.External) (Account, error)

	// GetWithCPF gets the account with the document.Document specified. It may return ErrNotFound
	// if no account with that document.Document is found, document.ErrInvalidDocument if the document.Document provided is corrupt
	// or ErrFetching if an untracked error happens
	GetWithCPF(context.Context, document.Document) (Account, error)

	// GetBalance gets the balance with the id.External specified. It may return ErrNotFound
	// if no account with the id.External is found, ErrInvalidID if the id.External provided is corrupt
	// or ErrFetching if an untracked error happens
	GetBalance(context.Context, id.External) (currency.Currency, error)

	// Create creates a new account and returns it with all field fulfilled. It must be called within a transaction.Transaction context
	// and may return ErrAlreadyExist if the data passed in already belongs to another Account, transaction.ErrNoTransaction
	// if there is no transaction in context or ErrCreating if an untracked error happens
	Create(context.Context, Account) (Account, error)

	// UpdateBalance replaces the balance of the account with the id.ID provided with a new one. It must be called within a transaction.Transaction context
	// and may return ErrNotFound if no account with the id.ID is found, transaction.ErrNoTransaction
	// if there is no transaction in context or ErrUpdating if an untracked error happens
	UpdateBalance(context.Context, id.External, currency.Currency) error
}
