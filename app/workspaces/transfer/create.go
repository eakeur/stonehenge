package transfer

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/entities/transfer"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"time"
)

type CreateInput struct {
	OriginID id.ExternalID
	DestID   id.ExternalID
	Amount   currency.Currency
}

type CreateOutput struct {
	RemainingBalance currency.Currency
	TransferId       id.ExternalID
	CreatedAt        time.Time
}

func (u *workspace) Create(ctx context.Context, req CreateInput) (CreateOutput, error) {
	t := &transfer.Transfer{
		Amount: req.Amount,
	}

	response := CreateOutput{}

	// Checks if the amount is valid (bigger than 0)
	if t.Amount <= 0 {
		return response, transfer.ErrAmountInvalid
	}

	// Checks if the origin and destination accounts are the same
	if t.DestinationID == t.OriginID {
		return response, transfer.ErrSameAccount
	}

	// Fetches the origin account and checks for errors
	origin, err := u.ac.Get(ctx, req.OriginID)
	if err != nil {
		return response, account.ErrNotFound
	}

	// Validates if the balance of the origin is zero and if it's sufficient to accomplish the operation
	if origin.Balance <= 0 || req.Amount > origin.Balance {
		return response, account.ErrNoMoney
	}

	// Fetches the origin account and checks for errors
	dest, err := u.ac.Get(ctx, req.DestID)
	if err != nil {
		return response, account.ErrNotFound
	}

	ctx, err = u.tx.Begin(ctx)
	if err != nil {
		return response, transfer.ErrRegistering
	}

	t.OriginID = origin.ID
	t.DestinationID = dest.ID

	// Updates the balance of the origin account after transaction
	remaining := origin.Balance - req.Amount
	err = u.ac.UpdateBalance(ctx, req.OriginID, remaining)
	if err != nil {
		u.tx.Rollback(ctx)
		return response, transfer.ErrRegistering
	}

	// Updates the balance of the destination account after transaction
	err = u.ac.UpdateBalance(ctx, req.DestID, dest.Balance+req.Amount)
	if err != nil {
		u.tx.Rollback(ctx)
		return response, transfer.ErrRegistering
	}

	// Creates a transfer register on storage
	t.EffectiveDate = time.Now()
	transferId, err := u.tr.Create(ctx, t)
	if err != nil {
		u.tx.Rollback(ctx)
		return response, transfer.ErrRegistering
	}

	err = u.tx.Commit(ctx)
	if err != nil {
		u.tx.Rollback(ctx)
		return response, transfer.ErrRegistering
	}

	response.RemainingBalance = remaining
	response.TransferId = transferId
	response.CreatedAt = t.CreatedAt
	return response, nil
}
