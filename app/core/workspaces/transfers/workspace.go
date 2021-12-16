package transfers

import (
	"context"
	"stonehenge/app/core/model/transfer"
)

//go:generate moq -fmt goimports -out usecase_mock.go . UseCase:UseCaseMock

type Workspace interface {
	List(ctx context.Context, filter transfer.Filter) ([]Reference, error)
}

type workspace struct {
	tr transfer.Repository
}

func New(tr transfer.Repository) *workspace {
	return &workspace{
		tr: tr,
	}
}
