package accounts

import (
	"context"
	"stonehenge/app/core/model/account"
	"stonehenge/app/core/types/id"
)

type ListRequest struct {
	Context context.Context
	Filter  account.Filter
}

type Reference struct {
	Id   id.ID
	Name string
}

func (u *workspace) List(ctx context.Context, filter account.Filter) ([]Reference, error) {
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
