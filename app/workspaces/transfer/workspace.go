package transfer

import (
	"context"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/entities/transfer"
)

//go:generate moq -fmt goimports -out usecase_mock.go . UseCase:UseCaseMock

type Workspace interface {
	List(ctx context.Context, filter transfer.Filter) ([]transfer.Transfer, error)
	Create(ctx context.Context, req CreateInput) (CreateOutput, error)
}

type workspace struct {
	accounts  account.Repository
	transfers    transfer.Repository
	transactions transaction.Manager
	access       access.Manager
}

func New(ac account.Repository, tr transfer.Repository, tx transaction.Manager, tk access.Manager) *workspace {
	return &workspace{
		accounts:     ac,
		transfers:    tr,
		transactions: tx,
		access:       tk,
	}
}
