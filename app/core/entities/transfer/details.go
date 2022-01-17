package transfer

import "stonehenge/app/core/types/id"

// Details wraps information about the two parts involved in the transfer
type Details struct {
	OriginExternalID      id.External
	DestinationExternalID id.External
}
