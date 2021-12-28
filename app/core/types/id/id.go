package id

import "github.com/google/uuid"

type ID string

func New() ID {
	return ID(uuid.New().String())
}

func From(id string) ID {
	return ID(id)
}
