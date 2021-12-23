package transfer

import (
	"stonehenge/app/core/types/audits"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/id"
	"time"
)

// Transfer holds useful information about a transaction between accounts
type Transfer struct {
	// ID is the unique identifier of this entity
	ID id.ID `json:"id"`

	// OriginID is the id of the account that will send money in this transfer
	OriginID id.ID `json:"account_origin_id"`

	// DestinationID is the id of the account that will receive money in this transfer
	DestinationID id.ID `json:"account_destination_id"`

	// Amount is the value to be transferred
	Amount currency.Currency `json:"amount"`

	// EffectiveDate is the date that the transaction happened
	EffectiveDate time.Time `json:"effective_date"`

	audits.Audit
}
