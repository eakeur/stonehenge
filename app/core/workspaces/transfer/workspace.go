package transfer

import (
	"stonehenge/app/core/model/transfer"
)

//go:generate moq -fmt goimports -out usecase_mock.go . UseCase:UseCaseMock

type Workspace interface {
	ListTransfers(filter transfer.Filter) ([]transfer.Transfer, error)
}

type workspace struct {
	tr transfer.Repository
}

func New(tr transfer.Repository) *workspace {
	return &workspace{
		tr: tr,
	}
}
