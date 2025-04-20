package store

import (
	"barrytime/go_templ_boilerplate/internal/model"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthStore struct {
	db *sqlx.DB
}

func NewAuthStore(db *sql.DB) *AuthStore {
	return &AuthStore{
		db: sqlx.NewDb(db, "postgres"),
	}
}

func (s *AuthStore) CreateUser(ctx context.Context, userRequest *model.NewUserRequest) (*model.User, error) {
	var user model.User
	const query = `INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING *`

	bytes, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	hashedPassword := base64.StdEncoding.EncodeToString(bytes)

	if err := s.db.GetContext(ctx, &user, query, userRequest.FirstName, userRequest.LastName, userRequest.Email, hashedPassword); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

func (s *AuthStore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	const query = `SELECT * FROM users WHERE email = $1`
	if err := s.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

func (s *AuthStore) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User
	const query = `SELECT * FROM users WHERE id = $1`
	if err := s.db.GetContext(ctx, &user, query, id); err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (u *AuthStore) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	const query = `SELECT * FROM users`
	if err := u.db.SelectContext(ctx, &users, query); err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return users, nil

}
