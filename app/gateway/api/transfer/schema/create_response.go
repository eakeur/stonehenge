package schema

type CreateTransferResponse struct {
	// RemainingBalance is how much budget the logged-in user still has after the transfer occurred
	RemainingBalance float64 `json:"remaining_balance"`
	// TransferID is the generated id that represents this transfer
	TransferID string `json:"transfer_id"`
}
