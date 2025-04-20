package main

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"barrytime/go_templ_boilerplate/internal/server"
	"barrytime/go_templ_boilerplate/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

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

	// load sessionStore
	sessionStore, err := server.NewSessionStore(cfg)
	if err != nil {
		log.Fatalf("failed to create session store: %v", err)
	}
	defer sessionStore.Close()

	server, err := server.New(cfg, dataStore, sessionStore)
	if err != nil {
		log.Fatalf("error creating server: %v", err)
	}

	log.Printf("Server listening on port %s\n", cfg.ApiServerPort)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
