package main

import (
	"context"
	"fmt"
	"os"
	"stonehenge/providers"
)

func main() {

	// Initializes all dependecies to be used by the endpoint tree
	// The intention here is to create one single instance of the database provider and the repositories
	// to be shared within ALL the application

	// Creates an object with access to the database provider
	provider, dbErr := providers.ConnectToDatabase(context.Background())
	if dbErr != nil {
		fmt.Printf("An error occurred while attempting to create a connection to the Stonehenge database: %v\n", dbErr)
		os.Exit(-1)
	}

	// Initializes all the repositories and injects the database provider reference
	providers.InjectDependenciesInRepositories(provider)

}
