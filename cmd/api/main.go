package main

import (
	"context"
	"log"
	"stonehenge/app/gateway/access"
	"stonehenge/app/gateway/api/server"
	"stonehenge/app/gateway/database/postgres"
	"stonehenge/app/gateway/database/postgres/transaction"
	"time"
)

func main() {

	ctx := context.Background()

	connection, err := postgres.NewConnection(ctx, "postgres://postgres:postgres@localhost:5432/stonehenge?sslmode=disable", "/home/igor/go/src/stonehenge/app/gateway/database/postgres/migrations", nil, 5)
	if err != nil {
		log.Fatalln(err)
	}

	helper := transaction.NewTransaction(connection)
	repos := server.NewPostgresRepositoryWrapper(connection)
	tokenFac := access.NewManager(10*time.Minute, []byte("EDF"))
	workspaces := server.NewWorkspaceWrapper(repos, helper, tokenFac)
	server.New(workspaces, tokenFac)
}
