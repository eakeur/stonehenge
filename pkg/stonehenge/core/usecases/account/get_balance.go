package account

import (
	"context"
	"stonehenge/pkg/stonehenge/core/types/currency"
	"stonehenge/pkg/stonehenge/core/types/id"
)

type GetBalanceResponse struct {
	Balance currency.Currency
}

func (u *useCase) GetBalance(ctx context.Context, id id.ID) (*GetBalanceResponse, error) {
	balance, err := u.ac.GetBalance(ctx, id)
	if err != nil {
		return nil, err
	}
	return &GetBalanceResponse{
		Balance: *balance,
	}, nil
}
