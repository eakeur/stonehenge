package transactions

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
)

type WithdrawalRequest struct {
	Context   context.Context
	AccountId id.ID
	Amount    currency.Currency
}

func (u *workspace) Withdraw(ctx context.Context, req WithdrawalRequest) (currency.Currency, error) {
	acc, err := u.ac.Get(ctx, req.AccountId)
	if err != nil {
		return 0, err
	}
	if req.Amount <= 0 {
		return 0, account.ErrAmountInvalid
	}
	if acc.Balance <= 0 || req.Amount > acc.Balance {
		return 0, account.ErrNoMoney
	}
	acc.Balance -= req.Amount
	err = u.ac.UpdateBalance(ctx, req.AccountId, acc.Balance)
	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}