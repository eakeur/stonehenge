package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
)

func (u *workspace) GetBalance(ctx context.Context, id id.External) (GetBalanceResponse, error) {
	const operation = "Workspaces.Account.GetBalance"
	actor, err := u.tk.GetAccessFromContext(ctx)
	if err != nil {
		u.logger.Error(ctx, operation, err.Error())
		return GetBalanceResponse{}, err
	}

	if id != actor.AccountID {
		u.logger.Error(ctx, operation, err.Error())
		return GetBalanceResponse{}, account.ErrCannotAccess
	}

	balance, err := u.ac.GetBalance(ctx, id)
	if err != nil {
		u.logger.Error(ctx, operation, err.Error())
		return GetBalanceResponse{}, err
	}

	u.logger.Trace(ctx, operation, "finished process successfully")
	return GetBalanceResponse{Balance: balance}, nil
}
