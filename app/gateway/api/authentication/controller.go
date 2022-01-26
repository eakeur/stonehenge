package authentication

import (
	"stonehenge/app/workspaces/authentication"
)

type Controller struct {
	workspace authentication.Workspace
}

func NewController(workspace authentication.Workspace) Controller {
	return Controller{
		workspace: workspace,
	}
}
