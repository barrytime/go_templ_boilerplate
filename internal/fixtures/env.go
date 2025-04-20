package fixtures

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"barrytime/go_templ_boilerplate/internal/store"
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestEnv struct {
	Config *config.Config
	DB     *sql.DB
	ctx    context.Context
}

func NewTestEnv(t *testing.T) *TestEnv {
	os.Setenv("ENV", string(config.EnvTest))
	conf, err := config.New()
	require.NoError(t, err)

	db, err := store.NewPostgresDB(conf)
	require.NoError(t, err)

	return &TestEnv{
		Config: conf,
		DB:     db,
	}
}
