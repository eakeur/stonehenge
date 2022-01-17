package transfer

import (
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"time"
)

type Reference struct {
	ExternalID    id.External
	OriginID      id.ID
	DestinationID id.ID
	Amount        currency.Currency
	EffectiveDate time.Time
}

type CreateInput struct {
	DestID id.External
	Amount currency.Currency
}

type CreateOutput struct {
	RemainingBalance currency.Currency
	TransferId       id.External
	CreatedAt        time.Time
}
