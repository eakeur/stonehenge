package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
	"stonehenge/app/core/types/password"
	"time"
)

// initialBalance expressed in cents
const initialBalance currency.Currency = 5000

type CreateInput struct {
	Document document.Document
	Secret   password.Password
	Name     string
}

type CreateOutput struct {
	AccountID id.ExternalID
	CreatedAt time.Time
}

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {

	response := CreateOutput{}

	// Checks document's consistency
	if err := req.Document.Validate(); err != nil {
		return response, err
	}

	//Checks document uniqueness
	res, err := u.ac.GetWithCPF(ctx, req.Document)
	if err != nil && err != account.ErrNotFound {
		return response, account.ErrCreating
	}
	if res.Document != "" {
		return response, account.ErrAlreadyExist
	}

	acc := account.Account{
		Name:     req.Name,
		Secret:   req.Secret,
		Document: req.Document,
		Balance:  initialBalance,
	}

	ctx, err = u.tx.Begin(ctx)
	if err != nil {
		return response, account.ErrCreating
	}
	defer u.tx.Rollback(ctx)

	accountId, err := u.ac.Create(ctx, &acc)
	if err != nil {
		return response, account.ErrCreating
	}
	err = u.tx.Commit(ctx)
	if err != nil {
		return response, account.ErrCreating
	}

	response.AccountID = accountId
	response.CreatedAt = acc.CreatedAt

	return response, nil
}
