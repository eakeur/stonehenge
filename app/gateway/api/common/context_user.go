package common

import (
	"context"
	"stonehenge/app/core/types/id"
)

type key int

const (
	accountContextId key = 41
)

// FetchContextUser fetches the logged-in user from the context and returns their id
func FetchContextUser(ctx context.Context) id.ExternalID {
	acc, found := ctx.Value(accountContextId).(id.ExternalID)
	if !found {
		return ""
	}
	return acc
}

// AssignUserToContext assigns the id of the logged-in user to a context
func AssignUserToContext(ctx context.Context, id id.ExternalID) context.Context {
	return context.WithValue(ctx, accountContextId, id)
}
