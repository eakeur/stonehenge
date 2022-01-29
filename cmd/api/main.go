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
	cfg := config.Config{
		Database: config.DatabaseConfigurations{
			User:           "postgres",
			Password:       "postgres",
			Host:           "localhost",
			Port:           "5432",
			Name:           "stonehenge",
			SSLMode:        "disable",
			MigrationsPath: "/home/igor/go/src/stonehenge/app/gateway/postgres/migrations",
		},
		Access: config.AccessConfigurations{
			ExpirationTime: "15",
			SigningKey:     "EB4CKU35",
		},
		Server: config.ServerConfigurations{
			ListenPort: "8080",
			Hostname:   "localhost",
		},
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
