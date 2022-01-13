package account

import (
	"context"
	"stonehenge/app/core/entities/account"
)

func (u *workspace) List(ctx context.Context, filter account.Filter) ([]Reference, error) {
	list, err := u.ac.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return toReference(list), nil
}

func toReference(accounts []account.Account) []Reference {
	refs := make([]Reference, len(accounts))
	for i, a := range accounts {
		refs[i] = Reference{
			ExternalID: a.ExternalID,
			Name:       a.Name,
		}
	}
	return refs
}
