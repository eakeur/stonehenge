package transfer

import (
	"net/http"
	"stonehenge/app/gateway/api/rest"
	transferWorker "stonehenge/app/worker/transfer"
	"stonehenge/app/workspaces/transfer"
)

type controller struct {
	workspace transfer.Workspace
	worker    transferWorker.Worker
	builder   rest.ResponseBuilder
}

func NewController(workspace transfer.Workspace, worker transferWorker.Worker,builder rest.ResponseBuilder) Controller {
	return &controller{workspace: workspace, worker: worker, builder: builder}
}

type Controller interface {
	Create(r *http.Request) rest.Response
	List(r *http.Request) rest.Response
}
