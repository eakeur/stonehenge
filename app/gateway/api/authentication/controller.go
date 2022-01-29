package authentication

import (
	"stonehenge/app/core/types/logger"
	"stonehenge/app/workspaces/authentication"
)

type Controller struct {
	workspace authentication.Workspace
	logger logger.Logger
}

func NewController(workspace authentication.Workspace, lg logger.Logger) Controller {
	return Controller{
		workspace: workspace,
		logger: lg,
	}
}
