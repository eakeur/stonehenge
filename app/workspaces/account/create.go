package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/errors"
)

// initialBalance expressed in cents
const initialBalance currency.Currency = 5000

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {
	const operation = "Workspaces.Account.Create"
	callParams := errors.AdditionalData{Key: "request", Value: req}

	acc := account.Account{
		Name:     req.Name,
		Secret:   req.Secret,
		Document: req.Document,
		Balance:  initialBalance,
	}

	if err := acc.Validate(); err != nil {
		return CreateOutput{}, errors.Wrap(err, operation, callParams)
	}

	ctx = u.tx.Begin(ctx)
	defer u.tx.End(ctx)

	acc, err := u.ac.Create(ctx, acc)
	if err != nil {
		return CreateOutput{}, errors.Wrap(err, operation, callParams)
	}

	tok, err := u.tk.Create(acc.ExternalID)
	if err != nil {
		return CreateOutput{}, errors.Wrap(err, operation, callParams)
	}

	return CreateOutput{
		AccountID: acc.ExternalID,
		CreatedAt: acc.CreatedAt,
		Access:    tok,
	}, nil
}
