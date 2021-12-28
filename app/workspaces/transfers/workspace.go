package transfers

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/model/transfer"
)

//go:generate moq -fmt goimports -out usecase_mock.go . UseCase:UseCaseMock

type Workspace interface {
	List(ctx context.Context, filter transfer.Filter) ([]Reference, error)
	Create(ctx context.Context, req CreateInput) (CreateOutput, error)
}

type workspace struct {
	ac account.Repository
	tr transfer.Repository
}

func New(ac account.Repository, tr transfer.Repository) *workspace {
	return &workspace{
		ac: ac,
		tr: tr,
	}
}
