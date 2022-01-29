package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/erring"
	"stonehenge/app/core/types/id"
)

func (u *workspace) GetBalance(ctx context.Context, id id.External) (GetBalanceResponse, error) {
	const operation = "Workspaces.Account.GetBalance"
	callParams := erring.AdditionalData{Key: "id", Value: id.String()}

	actor, err := u.access.GetAccessFromContext(ctx)
	if err != nil {
		return GetBalanceResponse{}, erring.Wrap(err, operation, callParams)
	}

	if id != actor.AccountID {
		return GetBalanceResponse{}, erring.Wrap(
			account.ErrCannotAccess,
			operation,
			callParams,
			erring.AdditionalData{Key: "actor", Value: actor.AccountID.String()},
		)
	}

	balance, err := u.accounts.GetBalance(ctx, id)
	if err != nil {
		return GetBalanceResponse{}, erring.Wrap(err, operation, callParams)
	}

	return GetBalanceResponse{Balance: balance}, nil
}
