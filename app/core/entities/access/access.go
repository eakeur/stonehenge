package access

import "stonehenge/app/core/types/id"

// Access holds information about the current logged-in user
type Access struct {
	AccountID   id.External
	AccountName string
	Token       string
}
