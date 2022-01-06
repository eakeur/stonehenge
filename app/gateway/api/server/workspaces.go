package server

import (
	"stonehenge/app/core/entities/transaction"
	"stonehenge/app/workspaces/account"
	"stonehenge/app/workspaces/transfer"
)

type WorkspaceWrapper struct {
	Accounts  account.Workspace
	Transfers transfer.Workspace
}

func NewWorkspaceWrapper(wrapper *RepositoryWrapper, helper transaction.Transaction) *WorkspaceWrapper {
	return &WorkspaceWrapper{
		Accounts:  account.New(wrapper.Account, helper),
		Transfers: transfer.New(wrapper.Account, wrapper.Transfer, helper),
	}
}
