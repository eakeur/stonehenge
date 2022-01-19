package access

import (
	"context"
	"stonehenge/app/core/entities/access"
)

func (f Manager) GetAccessFromContext(ctx context.Context) (access.Access, error) {
	acc, found := ctx.Value(accessContextId).(access.Access)
	if !found {
		return access.Access{}, access.ErrNoAccessInContext
	}
	return acc, nil
}
