package schema

type CreateResponse struct {
	RemainingBalance float64 `json:"remaining_balance"`
	TransferID       string `json:"transfer_id"`
}
