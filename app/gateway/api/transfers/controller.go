package transfers

import (
	"net/http"
	"stonehenge/app/core/types/logger"
	"stonehenge/app/gateway/api/rest"
	"stonehenge/app/workspaces/transfer"
)

type controller struct {
	workspace transfer.Workspace
	logger logger.Logger
}

func NewController(workspace transfer.Workspace) Controller {
	return controller{workspace: workspace}
}

type Controller interface {
	Create(r *http.Request) rest.Response
	List(r *http.Request) rest.Response
}
