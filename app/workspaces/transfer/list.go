package transfer

import (
	"context"
	"stonehenge/app/core/entities/transfer"
)

func (u *workspace) List(ctx context.Context, filter transfer.Filter) ([]Reference, error) {
	list, err := u.tr.List(ctx, filter)
	if err != nil {
		return nil, err
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
