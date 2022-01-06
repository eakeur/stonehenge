package account

import (
	"context"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
)


type RepositoryMock struct {
	ListFunc       func(context.Context, Filter) ([]Account, error)
	GetFunc        func(ctx context.Context, id id.ExternalID) (Account, error)
	GetWithCPFFunc func(ctx context.Context, document document.Document) (Account, error)
	GetBalanceFunc func(ctx context.Context, id id.ExternalID) (currency.Currency, error)
	CreateFunc     func(ctx context.Context, account *Account) (id.ExternalID, error)
	calls struct{
		List []listCall
		Get []getCall
		GetWithCPF []getWithCPFCall
		GetBalance []getCall
		Create []createCall
		UpdateBalance []updateBalanceCall
	}
}

func (r *RepositoryMock) List(ctx context.Context, filter Filter) ([]Account, error) {
	r.calls.List = append(r.calls.List, listCall{
		Ctx:    ctx,
		Filter: filter,
	})
	return r.ListFunc(ctx, filter)
}

func (r *RepositoryMock) Get(ctx context.Context, id id.ExternalID) (Account, error) {
	r.calls.Get = append(r.calls.Get, getCall{
		Ctx: ctx,
		ID:  id,
	})
	return r.GetFunc(ctx, id)
}

func (r *RepositoryMock) GetWithCPF(ctx context.Context, document document.Document) (Account, error) {
	r.calls.GetWithCPF = append(r.calls.GetWithCPF, getWithCPFCall{
		Ctx: ctx,
		Document:  document,
	})
	return r.GetWithCPFFunc(ctx, document)
}

func (r *RepositoryMock) GetBalance(ctx context.Context, id id.ExternalID) (currency.Currency, error) {
	r.calls.GetBalance = append(r.calls.GetBalance, getCall{
		Ctx: ctx,
		ID:  id,
	})
	return r.GetBalanceFunc(ctx, id)
}

func (r *RepositoryMock) Create(ctx context.Context, account *Account) (id.ExternalID, error) {
	r.calls.Create = append(r.calls.Create, createCall{
		Ctx:    ctx,
		Account: account,
	})
	return r.CreateFunc(ctx, account)
}

func (r *RepositoryMock) UpdateBalance(ctx context.Context, id id.ExternalID, balance currency.Currency) error {
	r.calls.UpdateBalance = append(r.calls.UpdateBalance, updateBalanceCall{
		Ctx:    ctx,
		ID: id,
		Balance: balance,
	})
	return r.UpdateBalance(ctx, id, balance)
}



type listCall struct {
	Ctx context.Context

	Filter Filter
}


type getCall struct {
	Ctx context.Context

	ID id.ExternalID
}

type getWithCPFCall struct {
	Ctx context.Context

	Document document.Document
}


type createCall struct {
	Ctx context.Context

	Account *Account
}

type updateBalanceCall struct {
	Ctx context.Context

	ID id.ExternalID

	Balance currency.Currency
}