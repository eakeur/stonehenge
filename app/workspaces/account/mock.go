package account

import (
	"context"
	"stonehenge/app/core/entities/account"
	"stonehenge/app/core/types/id"
)

var _ Workspace = workspaceMock{}


type workspaceMock struct {
	CreateResult CreateOutput
	GetBalanceResult GetBalanceResponse
	ListResult []account.Account
	Error error
}

func (w workspaceMock) Create(_ context.Context, _ CreateInput) (CreateOutput, error) {
	return w.CreateResult, w.Error
}

func (w workspaceMock) GetBalance(_ context.Context, _ id.External) (GetBalanceResponse, error) {
	return w.GetBalanceResult, w.Error
}

func (w workspaceMock) List(_ context.Context, _ account.Filter) ([]account.Account, error) {
	return w.ListResult, w.Error
}
