package transactions

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/core/types/currency"
)

type Workspace interface {
	Transfer(ctx context.Context, req TransferRequest)
	Withdraw(ctx context.Context, req WithdrawalRequest) (currency.Currency, error)
	Deposit(ctx context.Context, req DepositRequest) (currency.Currency, error)
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
