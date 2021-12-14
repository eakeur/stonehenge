package transfer

import (
	"stonehenge/pkg/stonehenge/core/model/transfer"
)

//go:generate moq -fmt goimports -out usecase_mock.go . UseCase:UseCaseMock

type UseCase interface {
	ListTransfers(filter transfer.Filter) ([]transfer.Transfer, error)
}

type useCase struct {
	tr transfer.Repository
}

func New(tr transfer.Repository) *useCase {
	return &useCase{
		tr: tr,
	}
}
