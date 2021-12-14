package account

import (
	"context"
	"stonehenge/pkg/stonehenge/core/types/currency"
	"stonehenge/pkg/stonehenge/core/types/id"
)

type GetBalanceRequest struct {
	Context context.Context
	Id      string
}

type GetBalanceResponse struct {
	Balance currency.Currency
}

func (u *useCase) GetBalance(request GetBalanceRequest) (*GetBalanceResponse, error) {
	balance, err := u.ac.GetBalance(request.Context, id.ID(request.Id))
	if err != nil {
		return nil, err
	}
	return &GetBalanceResponse{
		Balance: *balance,
	}, nil
}
