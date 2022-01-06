package transaction

import (
	"context"
)

func NewTransactionMock() Transaction {
	return &repositoryMock{}
}

type repositoryMock struct {
	BeginFunc func(ctx context.Context) (context.Context, error)
	CommitFunc func(ctx context.Context) error
	RollbackFunc func(ctx context.Context) error
}

func (r *repositoryMock) Begin(ctx context.Context) (context.Context, error) {
	return r.BeginFunc(ctx)
}

func (r *repositoryMock) Commit(ctx context.Context) error {
	return r.CommitFunc(ctx)
}

func (r *repositoryMock) Rollback(ctx context.Context) error {
	return r.RollbackFunc(ctx)
}

