package transfer

import (
	"context"
	"github.com/rs/zerolog"
	"stonehenge/app/workspaces/transfer"
)

const CloseCommand = "close"

type Worker interface {
	AddToQueue(ctx context.Context, input transfer.CreateInput) chan result
	Run()
	Close()
}

type worker struct {
	timeout   int
	logger    zerolog.Logger
	workspace transfer.Workspace
	queue     chan request
	stop      chan string
}

func NewWorker(timeout int, workspace transfer.Workspace, logger zerolog.Logger) Worker {
	w := worker{
		timeout:   timeout,
		logger:    logger,
		workspace: workspace,
		queue:     make(chan request),
		stop:      make(chan string),
	}

	return w
}
