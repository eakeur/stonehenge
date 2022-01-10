package transaction

import (
	"context"
)

type RepositoryMock struct {
	BeginFunc    func(ctx context.Context) (context.Context, error)
	CommitFunc   func(ctx context.Context) error
	RollbackFunc func(ctx context.Context) error
}

func (r *RepositoryMock) Begin(ctx context.Context) (context.Context, error) {
	return r.BeginFunc(ctx)
}

func (r *RepositoryMock) Commit(ctx context.Context) error {
	return r.CommitFunc(ctx)
}

func (r *RepositoryMock) Rollback(ctx context.Context) error {
	return r.RollbackFunc(ctx)
}
