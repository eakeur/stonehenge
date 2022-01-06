package account

import (
	"context"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
)

type GetBalanceResponse struct {
	Balance currency.Currency
}

func (u *workspace) GetBalance(ctx context.Context, id id.ExternalID) (GetBalanceResponse, error) {
	response := GetBalanceResponse{}
	balance, err := u.ac.GetBalance(ctx, id)
	if err != nil {
		return response, err
	}
	response.Balance = balance
	return response, nil
}
