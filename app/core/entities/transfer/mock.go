package transfer

import (
	"context"
	"stonehenge/app/core/types/id"
)

func NewRepositoryMock() Repository {
	return &repositoryMock{}
}

type repositoryMock struct {
	ListFunc   func(context.Context, Filter) ([]Transfer, error)
	GetFunc    func(ctx context.Context, id id.ExternalID) (Transfer, error)
	CreateFunc func(ctx context.Context, transfer *Transfer) (id.ExternalID, error)
	calls      struct {
		List []listCall

		Get []getCall

		Create []createCall
	}
}

func (r *repositoryMock) List(ctx context.Context, filter Filter) ([]Transfer, error) {
	r.calls.List = append(r.calls.List, listCall{
		Ctx:    ctx,
		Filter: filter,
	})
	return r.ListFunc(ctx, filter)
}

func (r *repositoryMock) Get(ctx context.Context, id id.ExternalID) (Transfer, error) {
	r.calls.Get = append(r.calls.Get, getCall{
		Ctx: ctx,
		ID:  id,
	})
	return r.GetFunc(ctx, id)
}

func (r *repositoryMock) Create(ctx context.Context, transfer *Transfer) (id.ExternalID, error) {
	r.calls.Create = append(r.calls.Create, createCall{
		Ctx:      ctx,
		Transfer: transfer,
	})
	return r.CreateFunc(ctx, transfer)
}

type listCall struct {
	Ctx context.Context

	Filter Filter
}

type getCall struct {
	Ctx context.Context

	ID id.ExternalID
}

type createCall struct {
	Ctx context.Context

	Transfer *Transfer
}
