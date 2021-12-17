package accounts

import (
	"stonehenge/app/workspaces/accounts"
)

type Controller struct {
	workspace accounts.Workspace
}

func New(workspace accounts.Workspace) Controller {
	return Controller{
		workspace: workspace,
	}
}
