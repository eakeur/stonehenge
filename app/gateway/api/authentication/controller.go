package authentication

import (
	"net/http"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/authentication"
)

type controller struct {
	workspace authentication.Workspace
	builder   rest.ResponseBuilder
}

func NewController(workspace authentication.Workspace, builder rest.ResponseBuilder) Controller {
	return &controller{
		workspace: workspace,
		builder:   builder,
	}
}

type Controller interface {
	Authenticate(r *http.Request) rest.Response
}
