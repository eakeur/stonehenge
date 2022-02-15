package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/erring"
	"time"
)

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {
	const operation = "Workspaces.Transfer.Create"

	actor, err := u.access.GetAccessFromContext(ctx)
	if err != nil {
		return CreateOutput{}, erring.Wrap(err, operation)
	}

	ctx = u.transactions.Begin(ctx)
	defer u.transactions.End(ctx)

	// Fetches the origin account and checks for errors
	origin, err := u.accounts.GetByExternalID(ctx, actor.AccountID)
	if err != nil {
		return CreateOutput{}, erring.Wrap(transfer.ErrNonexistentOrigin, operation)
	}

	// Fetches the origin account and checks for errors
	dest, err := u.accounts.GetByExternalID(ctx, req.DestID)
	if err != nil {
		return CreateOutput{}, transfer.ErrNonexistentDestination
	}

	t := transfer.Transfer{Amount: req.Amount, OriginID: origin.ID, DestinationID: dest.ID}
	if err := t.Validate(); err != nil {
		return CreateOutput{}, erring.Wrap(
			err,
			operation,
			erring.AdditionalData{Key: "actor", Value: actor.AccountID.String()},
		)
	}

	// Validates if the balance of the origin is zero and if it's sufficient to accomplish the operation
	if origin.Balance <= 0 || req.Amount > origin.Balance {
		return CreateOutput{}, erring.Wrap(account.ErrNoMoney, operation, erring.AdditionalData{Key: "origin_balance", Value: origin.Balance})
	}

	// Updates the balance of the origin account after transaction
	remaining := origin.Balance - req.Amount
	err = u.accounts.UpdateBalance(ctx, actor.AccountID, remaining)
	if err != nil {
		return CreateOutput{}, erring.Wrap(err, operation)
	}

	// Updates the balance of the destination account after transaction
	err = u.accounts.UpdateBalance(ctx, req.DestID, dest.Balance+req.Amount)
	if err != nil {
		return CreateOutput{}, erring.Wrap(err, operation)
	}

	// Creates a transfer register on storage
	t.EffectiveDate = time.Now()
	t, err = u.transfers.Create(ctx, t)
	if err != nil {
		return CreateOutput{}, erring.Wrap(err, operation)
	}

	return CreateOutput{
		RemainingBalance: remaining,
		TransferID:       t.ExternalID,
		CreatedAt:        t.CreatedAt,
	}, nil
}
