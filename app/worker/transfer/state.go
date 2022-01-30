package transfer

import (
	"context"
	transferworkspace "stonehenge/app/workspaces/transfer"
)

func (w worker) AddToQueue(ctx context.Context, input transferworkspace.CreateInput) chan result {
	ch := make(chan result)
	w.queue <- request{
		ctx:     ctx,
		request: input,
		output:  ch,
	}
	return ch
}

func (w worker) Close() {
	w.stop <- "close"
}
