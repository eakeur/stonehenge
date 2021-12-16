package transactions

import (
	"context"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
)

type DepositRequest struct {
	AccountId id.ID
	Amount    currency.Currency
}

func (u *workspace) Deposit(ctx context.Context, req DepositRequest) (currency.Currency, error) {
	acc, err := u.ac.Get(ctx, req.AccountId)
	if err != nil {
		return 0, err
	}
	acc.Balance = acc.Balance + req.Amount
	ctx, err = u.ac.StartOperation(ctx)
	if err != nil {
		//TODO create could not start operation error
		return 0, err
	}
	err = u.ac.UpdateBalance(ctx, req.AccountId, acc.Balance)
	err = u.ac.FinishOperation(ctx)
	if err != nil {
		//TODO create could not finish operation error
		return 0, err
	}

	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}
