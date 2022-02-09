package transfer

import (
	"context"
	"stonehenge/app/core/entities/transfer"
)

var _ Workspace = WorkspaceMock{}

type WorkspaceMock struct {
	ListResult   []transfer.Transfer
	CreateResult CreateOutput
	Error        error
}

func (w WorkspaceMock) List(_ context.Context, _ transfer.Filter) ([]transfer.Transfer, error) {
	return w.ListResult, w.Error
}

func (w WorkspaceMock) Create(_ context.Context, _ CreateInput) (CreateOutput, error) {
	return w.CreateResult, w.Error
}
