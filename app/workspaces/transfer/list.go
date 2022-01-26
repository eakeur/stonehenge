package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
)

func (u *workspace) List(ctx context.Context, filter transfer.Filter) ([]Reference, error) {
	const operation = "Workspaces.Transfer.List"
	actor, err := u.tk.GetAccessFromContext(ctx)
	if err != nil {
		u.logger.Error(ctx, operation, err.Error())
		return []Reference{}, err
	}

	if filter.OriginID != actor.AccountID && filter.DestinationID != actor.AccountID {
		u.logger.Error(ctx, operation, "cannot access data that has nothing to do with the logged in user")
		return []Reference{}, account.ErrCannotAccess
	}

	list, err := u.tr.List(ctx, filter)
	if err != nil {
		u.logger.Error(ctx, operation, err.Error())
		return []Reference{}, err
	}
	refs := make([]Reference, len(list))
	for i, a := range list {
		refs[i] = Reference{
			ExternalID:    a.ExternalID,
			OriginID:      a.OriginID,
			DestinationID: a.DestinationID,
			Amount:        a.Amount,
			EffectiveDate: a.EffectiveDate,
		}
	}
	u.logger.Trace(ctx, operation, "finished process successfully")
	return refs, nil
}
