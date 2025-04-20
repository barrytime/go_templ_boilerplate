package main

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"barrytime/go_templ_boilerplate/internal/store"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx := context.Background()

	// load config
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error initializing configuration: %v", err)
	}

	// load db
	db, err := store.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("error initializing db: %v", err)
	}
	defer db.Close()
	dataStore := store.New(db)

	if os.Args[1] == "up" {
		log.Println("Populate db")

		if err := dataStore.Migrations.MigrateUp(ctx, cfg); err != nil {
			log.Fatalf("error migrating up: %v", err)

		}
	} else if os.Args[1] == "down" {
		log.Println("Drop db")
		if err := dataStore.Migrations.MigrateDown(ctx, cfg); err != nil {
			log.Fatalf("error migrating down: %v", err)
		}

	} else {
		log.Println("error\nUsage:\ncmd [arg(up|down)]")
	}
}
