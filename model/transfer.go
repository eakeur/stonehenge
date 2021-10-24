package model

import (
	"time"
)

// Transfer hold useful information about transaction between accounts
type Transfer struct {
	// The unique identifier of this entity
	Id string `json:"id"`

	// The id of the origin account
	AccountOriginId string `json:"account_origin_id"`

	// The id of the destination account
	AccountDestinationId string `json:"account_destination_id"`

	// The value to be transacted
	Amount int64 `json:"amount"`

	// The time the transfer has been made
	CreatedAt time.Time `json:"created_at"`
}

// Returns a map of this transfer instance
func (t *Transfer) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                     t.Id,
		"account_origin_id":      t.AccountOriginId,
		"account_destination_id": t.AccountDestinationId,
		"amount":                 t.Amount,
		"created_at":             t.CreatedAt,
	}
}

// Returns an instance of a transfer based on the data passed as parameter
func TransferFromMap(mapInput map[string]interface{}) Transfer {
	return Transfer{
		Id:                   mapInput["id"].(string),
		AccountOriginId:      mapInput["account_origin_id"].(string),
		AccountDestinationId: mapInput["account_destination_id"].(string),
		Amount:               mapInput["amount"].(int64),
		CreatedAt:            mapInput["created_at"].(time.Time),
	}
}
