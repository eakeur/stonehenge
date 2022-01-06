package transfers

import (
	"stonehenge/app/workspaces/transfer"
)

type Controller struct {
	workspace transfer.Workspace
}

func New(workspace transfer.Workspace) Controller {
	return Controller{
		workspace: workspace,
	}
}
