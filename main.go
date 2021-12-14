package main

import (
	"fmt"
	"net/http"
	"os"
	"stonehenge/infra/persistence"
)

func main() {

	// Initializes all dependecies to be used by the endpoint tree
	// The intention here is to create one single instance of the database provider and the repositories
	// to be shared within ALL the application

	// Creates an object with access to the database provider
	wsp, dbErr := persistence.NewWorkspace("den1.mysql2.gear.host", "stonehenge", "Zg7J5_sm6fv?", "stonehenge")
	if dbErr != nil {
		fmt.Printf("An error occurred while attempting to create a connection to the Stonehenge database: %v\n", dbErr)
		os.Exit(-1)
	}

	server := NewStonehengeServer(wsp)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, server.MapControllers())
}
