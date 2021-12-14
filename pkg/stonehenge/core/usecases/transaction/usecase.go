package transaction

import (
	"context"
	"stonehenge/pkg/stonehenge/core/model/account"
	"stonehenge/pkg/stonehenge/core/model/transfer"
	"stonehenge/pkg/stonehenge/core/types/currency"
)

//go:generate moq -fmt goimports -out usecase_mock.go . UseCase:UseCaseMock

type UseCase interface {
	Transfer(ctx context.Context, req TransferRequest)
	Withdraw(ctx context.Context, req WithdrawalRequest) (currency.Currency, error)
	Deposit(ctx context.Context, req DepositRequest) (currency.Currency, error)
}
type useCase struct {
	ac account.Repository
	tr transfer.Repository
}

func New(ac account.Repository, tr transfer.Repository) *useCase {
	return &useCase{
		ac: ac,
		tr: tr,
	}
}
