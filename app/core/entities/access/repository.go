package access

import (
	"context"
	"stonehenge/app/core/types/id"
)

// Repository is an interface with useful actions to create and manage access token and objects
type Repository interface {
	// Create creates an Access object containing the account's external ID and its correspondent access token.
	Create(id.External) (Access, error)

	// ExtractAccessFromToken extracts an access object from the given token. It may return an error if the token
	// is invalid or expired
	ExtractAccessFromToken(string) (Access, error)

	// AssignAccessToContext assigns an Access object to a context so that a logged-in user can be known in other
	// parts of the application.
	AssignAccessToContext(context.Context, Access) context.Context

	// GetAccessFromContext returns an Access instance if there is any logged-in user in this context,
	// otherwise returns an error
	GetAccessFromContext(context.Context) (Access, error)
}
