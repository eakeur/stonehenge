package account

import (
	"context"
	"stonehenge/app/core/types/currency"
	"stonehenge/app/core/types/document"
	"stonehenge/app/core/types/id"
)

type RepositoryMock struct {
	ListFunc              func(context.Context, Filter) ([]Account, error)
	ListResult            []Account
	GetByExternalIDFunc   func(ctx context.Context, id id.External) (Account, error)
	GetByExternalIDResult Account
	GetWithCPFFunc        func(ctx context.Context, document document.Document) (Account, error)
	GetWithCPFResult      Account
	GetBalanceFunc        func(ctx context.Context, id id.External) (currency.Currency, error)
	GetBalanceResult      currency.Currency
	CreateFunc            func(ctx context.Context, account Account) (Account, error)
	CreateResult          Account
	UpdateBalanceFunc     func(ctx context.Context, id id.External, balance currency.Currency) error
	Error                 error
	calls                 struct {
		List          []listCall
		Get           []getCall
		GetWithCPF    []getWithCPFCall
		GetBalance    []getCall
		Create        []createCall
		UpdateBalance []updateBalanceCall
	}
}

func (r *RepositoryMock) List(ctx context.Context, filter Filter) ([]Account, error) {
	r.calls.List = append(r.calls.List, listCall{
		Ctx:    ctx,
		Filter: filter,
	})
	if r.ListFunc == nil {
		return r.ListResult, r.Error
	}
	return r.ListFunc(ctx, filter)
}

func (r *RepositoryMock) GetByExternalID(ctx context.Context, id id.External) (Account, error) {
	r.calls.Get = append(r.calls.Get, getCall{
		Ctx: ctx,
		ID:  id,
	})
	if r.GetByExternalIDFunc == nil {
		return r.GetByExternalIDResult, r.Error
	}
	return r.GetByExternalIDFunc(ctx, id)
}

func (r *RepositoryMock) GetWithCPF(ctx context.Context, document document.Document) (Account, error) {
	r.calls.GetWithCPF = append(r.calls.GetWithCPF, getWithCPFCall{
		Ctx:      ctx,
		Document: document,
	})
	if r.GetWithCPFFunc == nil {
		return r.GetWithCPFResult, r.Error
	}
	return r.GetWithCPFFunc(ctx, document)
}

func (r *RepositoryMock) GetBalance(ctx context.Context, id id.External) (currency.Currency, error) {
	r.calls.GetBalance = append(r.calls.GetBalance, getCall{
		Ctx: ctx,
		ID:  id,
	})
	if r.GetBalanceFunc == nil {
		return r.GetBalanceResult, r.Error
	}
	return r.GetBalanceFunc(ctx, id)
}

func (r *RepositoryMock) Create(ctx context.Context, account Account) (Account, error) {
	r.calls.Create = append(r.calls.Create, createCall{
		Ctx:     ctx,
		Account: account,
	})
	if r.CreateFunc == nil {
		return r.CreateResult, r.Error
	}
	return r.CreateFunc(ctx, account)
}

func (r *RepositoryMock) UpdateBalance(ctx context.Context, id id.External, balance currency.Currency) error {
	r.calls.UpdateBalance = append(r.calls.UpdateBalance, updateBalanceCall{
		Ctx:     ctx,
		ID:      id,
		Balance: balance,
	})
	if r.UpdateBalanceFunc == nil {
		return r.Error
	}
	return r.UpdateBalanceFunc(ctx, id, balance)
}

type listCall struct {
	Ctx context.Context

	Filter Filter
}

type getCall struct {
	Ctx context.Context

	ID id.External
}

type getWithCPFCall struct {
	Ctx context.Context

	Document document.Document
}

type createCall struct {
	Ctx context.Context

	Account Account
}

type updateBalanceCall struct {
	Ctx context.Context

	ID id.External

	Balance currency.Currency
}
