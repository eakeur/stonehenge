package account

import (
	"net/http"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/account"
)

type controller struct {
	builder   rest.ResponseBuilder
	workspace account.Workspace
}

func NewController(workspace account.Workspace, builder rest.ResponseBuilder) Controller {
	return &controller{
		workspace: workspace,
		builder:   builder,
	}
}

type Controller interface {
	Create(r *http.Request) rest.Response
	List(r *http.Request) rest.Response
	GetBalance(r *http.Request) rest.Response
}
