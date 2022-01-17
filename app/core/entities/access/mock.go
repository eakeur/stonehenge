package access

import (
	"context"
	"stonehenge/app/core/types/id"
)

var _ Factory = &FactoryMock{}

type FactoryMock struct {
	CreateResult                 Access
	ExtractAccessFromTokenResult Access
	AssignAccessToContextResult  context.Context
	GetAccessFromContextResult   Access
	Error                        error
}

func (f FactoryMock) Create(_ id.External) (Access, error) {
	return f.CreateResult, f.Error
}

func (f FactoryMock) ExtractAccessFromToken(_ string) (Access, error) {
	return f.ExtractAccessFromTokenResult, f.Error
}

func (f FactoryMock) AssignAccessToContext(_ context.Context, _ Access) context.Context {
	return f.AssignAccessToContextResult
}

func (f FactoryMock) GetAccessFromContext(_ context.Context) (Access, error) {
	return f.GetAccessFromContextResult, f.Error
}
