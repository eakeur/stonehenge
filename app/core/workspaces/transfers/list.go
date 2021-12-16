package transfers

import (
	"context"
	"stonehenge/app/core/model/transfer"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"time"
)

type Reference struct {
	Id id.ID

	OriginId id.ID

	DestinationId id.ID

	Amount currency.Currency

	EffectiveDate time.Time
}

func (u *workspace) List(ctx context.Context, filter transfer.Filter) ([]Reference, error) {
	list, err := u.tr.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	refs := make([]Reference, 0, len(list))
	for i, a := range list {
		refs[i] = Reference{
			Id:            a.Id,
			OriginId:      a.OriginId,
			DestinationId: a.DestinationId,
			Amount:        a.Amount,
			EffectiveDate: a.EffectiveDate,
		}
	}
	return refs, nil
}
