package transfers

import (
	"stonehenge/app/core/workspaces/transfers"
)

type Controller struct {
	workspace transfers.Workspace
}

func New(workspace transfers.Workspace) Controller {
	return Controller{
		workspace: workspace,
	}
}
