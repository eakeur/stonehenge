package schema

import (
	"encoding/json"
	"io"
)

type CreateTransferRequest struct {
	// DestinationID is the ID of the transfer that wil receive the money
	DestinationID string `json:"account_destination_id" example:"123e4567-e89b-12d3-a456-426614174000"`

	// Amount is the quantity that is going to be attempted to transfer
	Amount float64 `json:"amount" example:"15.99"`
}

func NewCreateTransferRequest(body io.ReadCloser) (CreateTransferRequest, error) {
	defer body.Close()
	req := CreateTransferRequest{}
	err := json.NewDecoder(body).Decode(&req)
	if err != nil {
		return CreateTransferRequest{}, err
	}
	return req, nil
}
