package schema

import (
	"time"
)

type ListResponse struct {
	ExternalID    string `json:"external_id"`
	OriginID      string `json:"origin_id"`
	DestinationID string `json:"destination_id"`
	Amount        float64 `json:"amount"`
	EffectiveDate time.Time `json:"effective_date"`
}
