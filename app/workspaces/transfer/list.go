package transfer

import (
	"context"
	"stonehenge/app/core/entities/transfer"
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
	refs := make([]Reference, len(list))
	for i, a := range list {
		refs[i] = Reference{
			Id:            a.ID,
			OriginId:      a.OriginID,
			DestinationId: a.DestinationID,
			Amount:        a.Amount,
			EffectiveDate: a.EffectiveDate,
		}
	}
	return refs, nil
}