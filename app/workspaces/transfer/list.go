package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/errors"
)

func (u *workspace) List(ctx context.Context, filter transfer.Filter) ([]transfer.Transfer, error) {
	const operation = "Workspaces.Transfer.List"
	callParams := errors.AdditionalData{Key: "filter", Value: filter}

	actor, err := u.access.GetAccessFromContext(ctx)
	if err != nil {
		return []transfer.Transfer{}, errors.Wrap(err, operation, callParams)
	}

	if filter.OriginID != actor.AccountID && filter.DestinationID != actor.AccountID {
		return []transfer.Transfer{}, errors.Wrap(
			account.ErrCannotAccess,
			operation,
			callParams,
			errors.AdditionalData{Key: "actor", Value: actor.AccountID.String()},
		)
	}

	list, err := u.transfers.List(ctx, filter)
	if err != nil {
		return []transfer.Transfer{}, errors.Wrap(err, operation, callParams)
	}

	return list, nil
}
