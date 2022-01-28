package transaction

import (
	"context"
)


type RepositoryMock struct {
	BeginFunc    func(ctx context.Context) context.Context
	BeginResult context.Context
	CommitFunc   func(ctx context.Context) error
	RollbackFunc func(ctx context.Context)
	EndFunc		 func(ctx context.Context)
	Error error
}

func (r *RepositoryMock) Begin(ctx context.Context) context.Context{
	if r.BeginFunc == nil {
		return r.BeginResult
	}
	return r.BeginFunc(ctx)
}

func (r *RepositoryMock) Commit(ctx context.Context) error {
	if r.CommitFunc == nil {
		return r.Error
	}
	return r.CommitFunc(ctx)
}

func (r *RepositoryMock) Rollback(ctx context.Context) {
	if r.RollbackFunc != nil {
		r.RollbackFunc(ctx)
	}
}

func (r *RepositoryMock) End(ctx context.Context){
	if r.EndFunc != nil {
		r.EndFunc(ctx)
	}
}
