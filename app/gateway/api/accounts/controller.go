package accounts

import (
	"stonehenge/app/core/types/logger"
	"stonehenge/app/workspaces/account"
)

type Controller struct {
	workspace account.Workspace
	logger logger.Logger
}

func NewController(workspace account.Workspace, lg logger.Logger) Controller {
	return Controller{
		workspace: workspace,
		logger: lg,
	}
}
