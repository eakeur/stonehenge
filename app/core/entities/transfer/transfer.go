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
	ID id.ID

	// ExternalID is the public identifier of this entity
	ExternalID id.External

	// OriginID is the id of the account that will send money in this transfer
	OriginID id.ID

	// DestinationID is the id of the account that will receive money in this transfer
	DestinationID id.ID

	// Amount is the value to be transferred
	Amount currency.Currency

	// EffectiveDate is the date that the transaction happened
	EffectiveDate time.Time

	// Details is a wrapper for identifying the two parts involved in this transfer externally
	Details Details

	audits.Audit
}
