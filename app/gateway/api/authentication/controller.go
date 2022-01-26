package authentication

import (
	"stonehenge/app/core/types/logger"
	"stonehenge/app/workspaces/authentication"
)

type Controller struct {
	workspace authentication.Workspace
	logger logger.Logger
}

func NewController(workspace authentication.Workspace) Controller {
	return Controller{
		workspace: workspace,
	}
}
