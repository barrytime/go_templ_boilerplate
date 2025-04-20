package store

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	Migrations *MigrationStore
	AuthStore  *AuthStore
}

func New(db *sql.DB) *Store {
	return &Store{
		Migrations: NewMigrationStore(db),
		AuthStore:  NewAuthStore(db),
	}
}

func NewPostgresDB(conf *config.Config) (*sql.DB, error) {
	dsn := conf.DatabaseURL()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
