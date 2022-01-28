package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/errors"
)

func (u *workspace) List(ctx context.Context, filter account.Filter) ([]account.Account, error) {
	const operation = "Workspaces.Account.List"
	callParams := errors.AdditionalData{Key: "filter", Value: filter}

	_, err := u.access.GetAccessFromContext(ctx)
	if err != nil {
		return []account.Account{}, errors.Wrap(err, operation, callParams)
	}

	list, err := u.accounts.List(ctx, filter)
	if err != nil {
		return []account.Account{}, errors.Wrap(err, operation, callParams)
	}

	return list, nil
}
