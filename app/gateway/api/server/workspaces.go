package server

import (
	"stonehenge/app/workspaces/accounts"
	"stonehenge/app/workspaces/transfers"
)

type WorkspaceWrapper struct {
	Accounts  accounts.Workspace
	Transfers transfers.Workspace
}

func NewWorkspaceWrapper(wrapper *RepositoryWrapper) *WorkspaceWrapper {
	return &WorkspaceWrapper{
		Accounts:  accounts.New(wrapper.Account),
		Transfers: transfers.New(wrapper.Account, wrapper.Transfer),
	}
}
