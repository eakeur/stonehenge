package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
)

// initialBalance expressed in cents
const initialBalance currency.Currency = 5000

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {

	acc := account.Account{
		Name:     req.Name,
		Secret:   req.Secret,
		Document: req.Document,
		Balance:  initialBalance,
	}

	if err := acc.Validate(); err != nil {
		return CreateOutput{}, err
	}

	ctx, err := u.tx.Begin(ctx)
	if err != nil {
		return CreateOutput{}, account.ErrCreating
	}
	defer u.tx.Rollback(ctx)

	acc, err = u.ac.Create(ctx, acc)
	if err != nil {
		return CreateOutput{}, err
	}

	err = u.tx.Commit(ctx)
	if err != nil {
		return CreateOutput{}, account.ErrCreating
	}

	tok, err := u.tk.Create(acc.ExternalID)
	if err != nil {
		return CreateOutput{}, err
	}

	return CreateOutput{
		AccountID: acc.ExternalID,
		CreatedAt: acc.CreatedAt,
		Access:    tok,
	}, nil
}
