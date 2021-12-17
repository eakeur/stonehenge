package accounts

import (
	"stonehenge/app/core/workspaces/accounts"
)

type Controller struct {
	workspace accounts.Workspace
}

func New(workspace accounts.Workspace) Controller {
	return Controller{
		workspace: workspace,
	}
}
