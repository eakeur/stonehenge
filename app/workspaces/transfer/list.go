package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/erring"
)

func (u *workspace) List(ctx context.Context, filter transfer.Filter) ([]transfer.Transfer, error) {
	const operation = "Workspaces.Transfer.List"

	actor, err := u.access.GetAccessFromContext(ctx)
	if err != nil {
		return []transfer.Transfer{}, erring.Wrap(err, operation)
	}

	if filter.OriginID != actor.AccountID && filter.DestinationID != actor.AccountID {
		return []transfer.Transfer{}, erring.Wrap(
			account.ErrCannotAccess,
			operation,
			erring.AdditionalData{Key: "actor", Value: actor.AccountID.String()},
		)
	}

	list, err := u.transfers.List(ctx, filter)
	if err != nil {
		return []transfer.Transfer{}, erring.Wrap(err, operation)
	}

	return list, nil
}
