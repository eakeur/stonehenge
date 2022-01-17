package id

import "github.com/google/uuid"

var Zero = External(uuid.Nil)

type External uuid.UUID

func NewExternal() External {
	return External(uuid.New())
}

func ExternalFrom(id string) External {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return Zero
	}
	return External(parsed)
}
