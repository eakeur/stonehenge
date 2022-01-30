package schema

import (
	"time"
)

type ListResponse struct {
	ExternalID    string `json:"id"`
	OriginID      string `json:"account_origin_id"`
	DestinationID string `json:"account_destination_id"`
	Amount        float64 `json:"amount"`
	EffectiveDate time.Time `json:"effective_date"`
}
