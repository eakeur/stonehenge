package accounts

import (
	"context"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
)

type GetBalanceResponse struct {
	Balance currency.Currency
}

func (u *workspace) GetBalance(ctx context.Context, id id.ID) (*GetBalanceResponse, error) {
	balance, err := u.ac.GetBalance(ctx, id)
	if err != nil {
		return nil, err
	}
	return &GetBalanceResponse{
		Balance: *balance,
	}, nil
}
