package dto

import "time"

type TransferDTO struct {
	// The unique identifier of this entity
	Id string `json:"id"`

	// The id of the origin account
	AccountOriginId string `json:"account_origin_id"`

	// The id of the destination account
	AccountDestinationId string `json:"account_destination_id"`

	// The value to be transacted
	Amount int64 `json:"amount"`

	// The transaction effective date
	Date time.Time `json:"date"`
}
