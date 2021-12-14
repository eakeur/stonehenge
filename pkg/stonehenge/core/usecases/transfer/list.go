package transfer

import (
	"context"
	"stonehenge/pkg/stonehenge/core/model/transfer"
	"stonehenge/pkg/stonehenge/core/types/currency"
	"stonehenge/pkg/stonehenge/core/types/id"
	"time"
)

type ListRequest struct {
	Context context.Context
	Filter  transfer.Filter
}

type Reference struct {
	Id id.ID

	OriginId id.ID

	DestinationId id.ID

	Amount currency.Currency

	EffectiveDate time.Time
}

func (u *useCase) List(request ListRequest) ([]Reference, error) {
	list, err := u.tr.List(request.Context, request.Filter)
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
