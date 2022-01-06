package transfer

import (
	"stonehenge/app/core/types/id"
	"time"
)

// Filter stores information that refines the transfer list, bringing up only what is needed
type Filter struct {
	// OriginID filters transfers made by this account only
	OriginID id.ExternalID

	// DestinationID filters transfers made to this account only
	DestinationID id.ExternalID

	// InitialDate filters transfers made at this time or later
	InitialDate time.Time

	// FinalDate filters transfers made at this time or earlier
	FinalDate time.Time
}
