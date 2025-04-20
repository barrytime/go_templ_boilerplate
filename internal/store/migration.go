package store

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MigrationStore struct {
	db *sqlx.DB
}

func NewMigrationStore(db *sql.DB) *MigrationStore {
	return &MigrationStore{
		db: sqlx.NewDb(db, "postgres"),
	}
}

func (s *MigrationStore) ExecuteFromFS(ctx context.Context, sqlFile string) error {
	sqlString, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	if _, err := s.db.ExecContext(ctx, string(sqlString[:])); err != nil {
		return err
	}

	return nil
}

func (s *MigrationStore) MigrateUp(ctx context.Context, cfg *config.Config) error {
	return s.ExecuteFromFS(ctx, fmt.Sprintf("%s/init.up.sql", cfg.SqlPath))
}

func (s *MigrationStore) MigrateDown(ctx context.Context, cfg *config.Config) error {
	return s.ExecuteFromFS(ctx, fmt.Sprintf("%s/init.down.sql", cfg.SqlPath))
}
