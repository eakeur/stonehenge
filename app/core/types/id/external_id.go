package id

import "github.com/google/uuid"

var ZeroValue = ExternalID(uuid.MustParse("00000000-0000-0000-0000-000000000000"))

type ExternalID uuid.UUID

func New() ExternalID {
	return ExternalID(uuid.New())
}

func From(id string) ExternalID {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return ZeroValue
	}
	return ExternalID(parsed)
}
