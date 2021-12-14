package transaction

import (
	"context"
	"stonehenge/pkg/stonehenge/core/model/account"
	"stonehenge/pkg/stonehenge/core/types/currency"
	"stonehenge/pkg/stonehenge/core/types/id"
)

type WithdrawalRequest struct {
	Context   context.Context
	AccountId id.ID
	Amount    currency.Currency
}

func (u *useCase) Withdraw(ctx context.Context, req WithdrawalRequest) (currency.Currency, error) {
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
	err = u.ac.Update(ctx, req.AccountId, acc)
	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}
