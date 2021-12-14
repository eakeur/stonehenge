package transaction

import (
	"context"
	"stonehenge/pkg/stonehenge/core/types/currency"
	"stonehenge/pkg/stonehenge/core/types/id"
)

type DepositRequest struct {
	AccountId id.ID
	Amount    currency.Currency
}

func (u *useCase) Deposit(ctx context.Context, req DepositRequest) (currency.Currency, error) {
	acc, err := u.ac.Get(ctx, req.AccountId)
	if err != nil {
		return 0, err
	}
	acc.Balance = acc.Balance + req.Amount
	err = u.ac.Update(ctx, req.AccountId, acc)
	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}
