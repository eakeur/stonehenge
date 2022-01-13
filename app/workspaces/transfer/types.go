package transfer

import (
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"time"
)

type Reference struct {
	ExternalID    id.ExternalID
	OriginID      id.ID
	DestinationID id.ID
	Amount        currency.Currency
	EffectiveDate time.Time
}

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
