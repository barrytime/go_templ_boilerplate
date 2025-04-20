package store_test

import (
	"barrytime/go_templ_boilerplate/internal/fixtures"
	"barrytime/go_templ_boilerplate/internal/store"
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestMigrationStore(t *testing.T) {
	err := godotenv.Load("../../.env")
	require.NoError(t, err)

	env := fixtures.NewTestEnv(t)
	defer env.DB.Close()
	ctx := context.Background()

	// load db
	dataStore := store.New(env.DB)

	// cleanup func
	defer func() {
		err := dataStore.Migrations.MigrateDown(ctx, env.Config)
		require.NoError(t, err, "failed to migrate down during cleanup")
	}()
	err = dataStore.Migrations.MigrateUp(ctx, env.Config)
	require.NoError(t, err)
}
