package transfer

import (
	"context"
	"stonehenge/app/core/entities/access"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/logger"
)

//go:generate moq -fmt goimports -out usecase_mock.go . UseCase:UseCaseMock

type Workspace interface {
	List(ctx context.Context, filter transfer.Filter) ([]Reference, error)
	Create(ctx context.Context, req CreateInput) (CreateOutput, error)
}

type workspace struct {
	ac account.Repository
	tr transfer.Repository
	tx transaction.Manager
	tk access.Manager
	logger logger.Logger
}

func New(ac account.Repository, tr transfer.Repository, tx transaction.Manager, tk access.Manager, lg logger.Logger) *workspace {
	return &workspace{
		ac: ac,
		tr: tr,
		tx: tx,
		tk: tk,
		logger: lg,
	}
}
