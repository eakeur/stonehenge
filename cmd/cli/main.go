package main

import (
	"context"
	"log"
	"stonehenge/app/gateway/api/server"
	"stonehenge/app/gateway/database/postgres"
	"stonehenge/app/gateway/database/postgres/transaction"
)

func main() {

	ctx := context.Background()

	connection, err := postgres.NewConnection(ctx, "postgres://postgres:postgres@localhost:5432/stonehenge?sslmode=disable", "/home/igor/go/src/stonehenge/app/gateway/database/postgres/migrations", nil, 5)
	if err != nil {
		log.Fatalln(err)
	}

	connection.Query(ctx, "select * from accounts")

	adapter := transaction.NewTransactionAdapter(connection)

	repos := server.NewPostgresRepositoryWrapper(connection, adapter)

	server.NewWorkspaceWrapper(repos)
}
