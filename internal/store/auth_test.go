package store_test

import (
	"barrytime/go_templ_boilerplate/internal/fixtures"
	"barrytime/go_templ_boilerplate/internal/model"
	"barrytime/go_templ_boilerplate/internal/store"

	"context"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func testUser(t *testing.T, expected *model.NewUserRequest, actual *model.User) {
	now := time.Now()
	require.Equal(t, expected.Email, actual.Email)
	require.Equal(t, expected.FirstName, actual.FirstName)
	require.Equal(t, expected.LastName, actual.LastName)
	require.WithinDuration(t, now, actual.CreatedAt, time.Second)
}

func TestAuthStore(t *testing.T) {
	err := godotenv.Load("../../.env")
	require.NoError(t, err)

	env := fixtures.NewTestEnv(t)
	defer env.DB.Close()
	ctx := context.Background()

	// load db
	dataStore := store.New(env.DB)
	defer func() {
		err := dataStore.Migrations.MigrateDown(ctx, env.Config)
		require.NoError(t, err, "failed to migrate down during cleanup")
	}()
	err = dataStore.Migrations.MigrateUp(ctx, env.Config)
	require.NoError(t, err)

	// create user
	newUserRequest1 := &model.NewUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "jon@gmail.com",
		Password:  "password123",
	}

	user1, err := dataStore.AuthStore.CreateUser(ctx, newUserRequest1)
	require.NoError(t, err)
	require.NotNil(t, user1)
	testUser(t, newUserRequest1, user1)

	// compare password
	err = user1.ComparePassword(newUserRequest1.Password)
	require.NoError(t, err)
	// compare wrong password -> error
	err = user1.ComparePassword("wrong password")
	require.Error(t, err)
	require.NotEqual(t, user1.Password, newUserRequest1.Password)

	// get user by id
	user1a, err := dataStore.AuthStore.GetUserByID(ctx, user1.ID)
	require.NoError(t, err)
	require.NotNil(t, user1a)
	testUser(t, newUserRequest1, user1a)

	// get user by email
	user1b, err := dataStore.AuthStore.GetUserByEmail(ctx, user1.Email)
	require.NoError(t, err)
	require.NotNil(t, user1b)
	testUser(t, newUserRequest1, user1b)

	// create user with same email -> error
	nullUser, err := dataStore.AuthStore.CreateUser(ctx, newUserRequest1)
	require.Error(t, err)
	require.Nil(t, nullUser)

	// get all users
	users, err := dataStore.AuthStore.GetAllUsers(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.Len(t, users, 1)

}
