package access

import (
	"context"
	"stonehenge/app/core/entities/access"
)

func (f Factory) AssignAccessToContext(ctx context.Context, acc access.Access) context.Context {
	return context.WithValue(ctx, accessContextId, acc)
}
