package main

import (
	"context"
	"log"
	"stonehenge/app/gateway/api/server"
	"stonehenge/app/gateway/database/postgres"
)

func main() {

	ctx := context.Background()

	connection, err := postgres.NewConnection(ctx, "postgres://postgres:postgres@localhost:5432/stonehenge?sslmode=disable", "/home/igor/go/src/stonehenge/app/gateway/database/postgres/migrations", nil, 5)
	if err != nil {
		log.Fatalln(err)
	}

	adapter := postgres.NewTransactionAdapter(connection)

	repos := server.NewPostgresRepositoryWrapper(connection, adapter)

	workspaces := server.NewWorkspaceWrapper(repos)

	server.New(workspaces)
}