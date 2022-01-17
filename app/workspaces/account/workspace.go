package account

import (
	"context"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/types/id"
)

type Workspace interface {
	// Create creates a new account and returns its new id
	Create(ctx context.Context, req CreateInput) (CreateOutput, error)

	// GetBalance gets the account balance with the ID specified
	GetBalance(ctx context.Context, id id.External) (GetBalanceResponse, error)

	// List gets all accounts existing that satisfies the passed filter
	List(ctx context.Context, filter account.Filter) ([]Reference, error)

	// Authenticate verifies a user credential and returns the account id if it's all ok
	Authenticate(ctx context.Context, req AuthenticationRequest) (access.Access, error)
}

type workspace struct {
	ac account.Repository
	tx transaction.Transaction
	tk access.Factory
}

func New(ac account.Repository, tx transaction.Transaction, tk access.Factory) *workspace {
	return &workspace{
		ac: ac,
		tx: tx,
		tk: tk,
	}
}
