package transfer

import (
	"context"
	"stonehenge/app/workspaces/transfer"
)

type result struct {
	Output transfer.CreateOutput
	Error  error
}

type request struct {
	ctx     context.Context
	request transfer.CreateInput
	output  chan result
}
