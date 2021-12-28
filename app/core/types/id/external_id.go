package id

import "github.com/google/uuid"

type ExternalID string

func New() ExternalID {
	return ExternalID(uuid.NewString())
}

func From(id string) ExternalID {
	return ExternalID(id)
}

