package transfer

import "time"

// Filter stores information that refines the transfer list, bringing up only what is needed
type Filter struct {
	// OriginID filters transfers made by this account only
	OriginID string

	// DestinationID filters transfers made to this account only
	DestinationID string

	// InitialDate filters transfers made at this time or later
	InitialDate time.Time

	// FinalDate filters transfers made at this time or earlier
	FinalDate time.Time
}
