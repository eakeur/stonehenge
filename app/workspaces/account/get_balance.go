package account

import (
	"context"
	"stonehenge/app/core/types/id"
)

func (u *workspace) GetBalance(ctx context.Context, id id.External) (GetBalanceResponse, error) {
	balance, err := u.ac.GetBalance(ctx, id)
	if err != nil {
		return GetBalanceResponse{}, err
	}
	return GetBalanceResponse{Balance: balance}, nil
}
