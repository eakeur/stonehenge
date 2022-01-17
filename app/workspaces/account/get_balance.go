package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
)

func (u *workspace) GetBalance(ctx context.Context, id id.External) (GetBalanceResponse, error) {
	actor, err := u.tk.GetAccessFromContext(ctx)
	if err != nil {
		return GetBalanceResponse{}, err
	}

	if id != actor.AccountID {
		return GetBalanceResponse{}, account.ErrCannotAccess
	}
	balance, err := u.ac.GetBalance(ctx, id)
	if err != nil {
		return GetBalanceResponse{}, err
	}
	return GetBalanceResponse{Balance: balance}, nil
}
