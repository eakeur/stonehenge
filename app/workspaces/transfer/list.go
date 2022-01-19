package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
)

func (u *workspace) List(ctx context.Context, filter transfer.Filter) ([]Reference, error) {
	actor, err := u.tk.GetAccessFromContext(ctx)
	if err != nil {
		return []Reference{}, err
	}

	if filter.OriginID != actor.AccountID && filter.DestinationID != actor.AccountID {
		return []Reference{}, account.ErrCannotAccess
	}

	list, err := u.tr.List(ctx, filter)
	if err != nil {
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
	return refs, nil
}
