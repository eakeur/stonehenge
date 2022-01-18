package access

import (
	"context"
	"stonehenge/app/core/types/id"
)

var _ Repository = &RepositoryMock{}

type RepositoryMock struct {
	CreateResult                 Access
	ExtractAccessFromTokenResult Access
	AssignAccessToContextResult  context.Context
	GetAccessFromContextResult   Access
	Error                        error
}

func (f RepositoryMock) Create(_ id.External) (Access, error) {
	return f.CreateResult, f.Error
}

func (f RepositoryMock) ExtractAccessFromToken(_ string) (Access, error) {
	return f.ExtractAccessFromTokenResult, f.Error
}

func (f RepositoryMock) AssignAccessToContext(_ context.Context, _ Access) context.Context {
	return f.AssignAccessToContextResult
}

func (f RepositoryMock) GetAccessFromContext(_ context.Context) (Access, error) {
	return f.GetAccessFromContextResult, f.Error
}
