package account

import (
	"context"
	"stonehenge/pkg/stonehenge/core/model/account"
	"stonehenge/pkg/stonehenge/core/types/id"
)

type ListRequest struct {
	Context context.Context
	Filter  account.Filter
}

type Reference struct {
	Id   id.ID
	Name string
}

func (u *useCase) List(ctx context.Context, filter account.Filter) ([]Reference, error) {
	list, err := u.ac.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	refs := make([]Reference, 0, len(list))
	for i, a := range list {
		refs[i] = Reference{
			Id:   a.Id,
			Name: a.Name,
		}
	}
	return refs, nil
}
