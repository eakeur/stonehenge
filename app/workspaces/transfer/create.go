package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/errors"
	"time"
)

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {
	const operation = "Workspaces.Transfer.Create"
	callParams := errors.AdditionalData{Key: "request", Value: req}

	actor, err := u.tk.GetAccessFromContext(ctx)
	if err != nil {
		return CreateOutput{}, errors.Wrap(err, operation, callParams)
	}

	ctx = u.tx.Begin(ctx)
	defer u.tx.End(ctx)

	// Fetches the origin account and checks for errors
	origin, err := u.ac.GetByExternalID(ctx, actor.AccountID)
	if err != nil {
		return CreateOutput{}, errors.Wrap(transfer.ErrNonexistentOrigin, operation, callParams)
	}

	// Fetches the origin account and checks for errors
	dest, err := u.ac.GetByExternalID(ctx, req.DestID)
	if err != nil {
		return CreateOutput{}, transfer.ErrNonexistentDestination
	}

	t := transfer.Transfer{Amount: req.Amount, OriginID: origin.ID, DestinationID: dest.ID}
	if err := t.Validate(); err != nil {
		return CreateOutput{}, errors.Wrap(
			err,
			operation,
			callParams,
			errors.AdditionalData{Key: "actor", Value: actor.AccountID.String()},
		)
	}

	// Validates if the balance of the origin is zero and if it's sufficient to accomplish the operation
	if origin.Balance <= 0 || req.Amount > origin.Balance {
		return CreateOutput{}, errors.Wrap(account.ErrNoMoney, operation, callParams, errors.AdditionalData{Key: "origin_balance", Value: origin.Balance})
	}

	// Updates the balance of the origin account after transaction
	remaining := origin.Balance - req.Amount
	err = u.ac.UpdateBalance(ctx, actor.AccountID, remaining)
	if err != nil {
		return CreateOutput{}, errors.Wrap(account.ErrNoMoney, operation, callParams)
	}

	// Updates the balance of the destination account after transaction
	err = u.ac.UpdateBalance(ctx, req.DestID, dest.Balance+req.Amount)
	if err != nil {
		return CreateOutput{}, errors.Wrap(account.ErrNoMoney, operation, callParams)
	}

	// Creates a transfer register on storage
	t.EffectiveDate = time.Now()
	t, err = u.tr.Create(ctx, t)
	if err != nil {
		return CreateOutput{}, errors.Wrap(account.ErrNoMoney, operation, callParams)
	}

	return CreateOutput{
		RemainingBalance: remaining,
		TransferId:       t.ExternalID,
		CreatedAt:        t.CreatedAt,
	}, nil
}
