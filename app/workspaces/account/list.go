package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/erring"
)

func (u *workspace) List(ctx context.Context, filter account.Filter) ([]account.Account, error) {
	const operation = "Workspaces.Account.List"

	_, err := u.access.GetAccessFromContext(ctx)
	if err != nil {
		return []account.Account{}, erring.Wrap(err, operation)
	}

	list, err := u.accounts.List(ctx, filter)
	if err != nil {
		return []account.Account{}, erring.Wrap(err, operation)
	}

	return list, nil
}
