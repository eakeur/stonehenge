package account

import (
	"context"
	"stonehenge/pkg/stonehenge/core/model/account"
	"stonehenge/pkg/stonehenge/core/types/currency"
)

//go:generate moq -fmt goimports -out usecase_mock.go . UseCase:UseCaseMock

type UseCase interface {
	// Create creates a new account and returns its new id
	Create(ctx context.Context, req CreateInput) (*CreateOutput, error)

	// GetBalance gets the account balance with the ID specified
	GetBalance(ctx context.Context, id string) (*currency.Currency, error)

	// List gets all accounts existing that satisfies the passed filter
	List(ctx context.Context, filter account.Filter) ([]Reference, error)
}

type useCase struct {
	ac account.Repository
}

func New(ac account.Repository) *useCase {
	return &useCase{
		ac: ac,
	}
}
