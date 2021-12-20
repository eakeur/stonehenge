package accounts

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/types/id"
)

type Workspace interface {
	// Create creates a new account and returns its new id
	Create(ctx context.Context, req CreateInput) (*CreateOutput, error)

	// GetBalance gets the account balance with the ID specified
	GetBalance(ctx context.Context, id id.ID) (*GetBalanceResponse, error)

	// List gets all accounts existing that satisfies the passed filter
	List(ctx context.Context, filter account.Filter) ([]Reference, error)

	// Authenticate verifies a user credential and returns the account id if it's all ok
	Authenticate(ctx context.Context, req AuthenticationRequest) (*id.ID, error)
}

type workspace struct {
	ac account.Repository
}

func New(ac account.Repository) *workspace {
	return &workspace{
		ac: ac,
	}
}
