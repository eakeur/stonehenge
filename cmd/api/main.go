package main

import (
	"context"
	"fmt"
	"log"
	"stonehenge/app"
	"stonehenge/app/config"
	"stonehenge/app/gateway/api"
)

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
