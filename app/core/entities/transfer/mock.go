package transfer

import (
	"context"
)

type RepositoryMock struct {
	ListFunc     func(context.Context, Filter) ([]Transfer, error)
	ListResult   []Transfer
	CreateFunc   func(ctx context.Context, transfer Transfer) (Transfer, error)
	CreateResult Transfer
	Error        error
	calls        struct {
		List   []listCall
		Create []createCall
	}
}

func (r *RepositoryMock) List(ctx context.Context, filter Filter) ([]Transfer, error) {
	r.calls.List = append(r.calls.List, listCall{
		Ctx:    ctx,
		Filter: filter,
	})
	if r.ListFunc == nil {
		return r.ListResult, r.Error
	}
	return r.ListFunc(ctx, filter)
}

func (r *RepositoryMock) Create(ctx context.Context, transfer Transfer) (Transfer, error) {
	r.calls.Create = append(r.calls.Create, createCall{
		Ctx:      ctx,
		Transfer: transfer,
	})
	if r.CreateFunc == nil {
		return r.CreateResult, r.Error
	}
	return r.CreateFunc(ctx, transfer)
}

type listCall struct {
	Ctx context.Context

	Filter Filter
}

type createCall struct {
	Ctx context.Context

	Transfer Transfer
}
