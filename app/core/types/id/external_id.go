package id

import "github.com/google/uuid"

type ExternalID string

func New() ExternalID {
	return ExternalID(uuid.New().String())
}

func From(id string) ExternalID {
	return ExternalID(id)
}

