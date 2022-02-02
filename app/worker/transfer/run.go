package transfer

import (
	"stonehenge/app/core/types/erring"
	"time"
)

func (w worker) Run() {
	const operation = "Worker.Transfers.Create"
	for {
		time.Sleep(time.Duration(w.timeout) * time.Millisecond)
		select {
		case req := <-w.queue:
			res, err := w.workspace.Create(req.ctx, req.request)
			if err != nil {
				err = erring.Wrap(err, operation)
			}
			req.output <- result{Output: res, Error: err}
		case op := <-w.stop:
			if op == CloseCommand {
				break
			}
		}
	}
}
