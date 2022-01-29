package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/erring"
)

// initialBalance expressed in cents
const initialBalance currency.Currency = 5000

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {
	const operation = "Workspaces.Account.Create"

	acc := account.Account{
		Name:     req.Name,
		Secret:   req.Secret,
		Document: req.Document,
		Balance:  initialBalance,
	}

	if err := acc.Validate(); err != nil {
		return CreateOutput{}, erring.Wrap(err, operation)
	}

	ctx = u.transactions.Begin(ctx)
	defer u.transactions.End(ctx)

	acc, err := u.accounts.Create(ctx, acc)
	if err != nil {
		return CreateOutput{}, erring.Wrap(err, operation)
	}

	tok, err := u.access.Create(acc.ExternalID)
	if err != nil {
		return CreateOutput{}, erring.Wrap(err, operation)
	}

	return CreateOutput{
		AccountID: acc.ExternalID,
		CreatedAt: acc.CreatedAt,
		Access:    tok,
	}, nil
}
