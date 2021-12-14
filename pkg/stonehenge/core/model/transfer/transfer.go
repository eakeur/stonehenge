package transfer

import (
	"stonehenge/pkg/stonehenge/core/types/audits"
	"stonehenge/pkg/stonehenge/core/types/currency"
	"stonehenge/pkg/stonehenge/core/types/id"
	"time"
)

// Transfer holds useful information about a transaction between accounts
type Transfer struct {
	// ID is the unique identifier of this entity
	Id id.ID `json:"id"`

	// OriginId is the id of the account that will send money in this transfer
	OriginId id.ID `json:"account_origin_id"`

	// DestinationId is the id of the account that will receive money in this transfer
	DestinationId id.ID `json:"account_destination_id"`

	// Amount is the value to be transferred
	Amount currency.Currency `json:"amount"`

	// EffectiveDate is the date that the transaction happened
	EffectiveDate time.Time `json:"effective_date"`

	audits.Audit
}

// ToMap returns a map of this transfer instance
func (t *Transfer) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                     t.Id,
		"account_origin_id":      t.OriginId,
		"account_destination_id": t.DestinationId,
		"amount":                 t.Amount,
		"updated_at":             t.UpdatedAt,
		"created_at":             t.CreatedAt,
	}
}

// ToArray returns an array of this transfer instance
func (t *Transfer) ToArray() []interface{} {
	return []interface{}{
		t.Id,
		t.OriginId,
		t.DestinationId,
		t.Amount,
		t.UpdatedAt,
		t.CreatedAt,
	}
}
