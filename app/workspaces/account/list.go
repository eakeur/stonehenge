package account

import (
	"context"
	"stonehenge/app/core/entities/account"
)

func (u *workspace) List(ctx context.Context, filter account.Filter) ([]Reference, error) {
	list, err := u.ac.List(ctx, filter)
	if err != nil {
		return []Reference{}, err
	}
	refs := make([]Reference, len(list))
	for i, a := range list {
		refs[i] = Reference{
			ExternalID: a.ExternalID,
			Name:       a.Name,
		}
	}
	return refs, nil
}
