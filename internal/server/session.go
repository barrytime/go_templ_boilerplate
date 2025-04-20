package server

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"barrytime/go_templ_boilerplate/internal/model"
	"encoding/gob"

	"github.com/boj/redistore"
)

func NewSessionStore(cfg *config.Config) (*redistore.RediStore, error) {
	gob.Register(&model.User{}) // Enable session storage for custom structs
	redisStore, err := redistore.NewRediStore(10, "tcp", cfg.RedisURL(), "", "", []byte(cfg.SessionSecret))
	if err != nil {
		return redisStore, err
	}

	// Set session options
	redisStore.SetMaxAge(86400 * 7)     // 7 days session expiration
	redisStore.SetKeyPrefix("session:") // Prefix keys to avoid conflicts
	redisStore.Options.HttpOnly = true  // Prevent JS access to session cookie
	redisStore.Options.Secure = false   // Set to true in production (for HTTPS)

	return redisStore, nil
}
