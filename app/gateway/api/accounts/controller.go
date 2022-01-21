package accounts

import (
	"stonehenge/app/workspaces/account"
)

type Controller struct {
	workspace account.Workspace
}

func NewController(workspace account.Workspace) Controller {
	return Controller{
		workspace: workspace,
	}
}
