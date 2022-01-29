package transfer

import (
	"net/http"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/transfer"
)

type controller struct {
	workspace transfer.Workspace
	builder   rest.ResponseBuilder
}

func NewController(workspace transfer.Workspace, builder rest.ResponseBuilder) Controller {
	return &controller{workspace: workspace, builder: builder}
}

type Controller interface {
	Create(r *http.Request) rest.Response
	List(r *http.Request) rest.Response
}
