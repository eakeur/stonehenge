package access

import (
	"context"
	"stonehenge/app/core/entities/access"
)

func (f Manager) AssignAccessToContext(ctx context.Context, acc access.Access) context.Context {
	return context.WithValue(ctx, accessContextId, acc)
}
