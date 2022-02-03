package main

import (
	"context"
	"fmt"
	"log"
	"stonehenge/app"
	"stonehenge/app/config"
	"stonehenge/app/gateway/api"
)

// @title Stonehenge API
// @version 1.0
// @contact.name Igor Reis (@eakeur)
// @contact.email igor.reisleandro@gmail.com
// @securityDefinitions.apikey AuthKey
// @in header
// @name Authorization
func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfigurations()
	if err != nil {
		log.Fatalf("Could not load environment: %v", err)
	}

	application, err := app.NewApplication(ctx, cfg)
	if err != nil {
		log.Fatalf("Could not set up application: %v", err)
	}

	stonehenge := api.NewServer(application)
	addr := fmt.Sprintf("%s:%s", cfg.Server.Hostname, cfg.Server.ListenPort)
	log.Printf("Listening on http://%v", addr)
	err = stonehenge.Serve(addr)
	if err != nil {
		log.Fatalf("Server shut down: %v", err)
	}
}
