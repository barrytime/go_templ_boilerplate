package server

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"barrytime/go_templ_boilerplate/internal/handler"
	"barrytime/go_templ_boilerplate/internal/store"
	"net/http"
	"time"

	"github.com/boj/redistore"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config       *config.Config
	store        *store.Store
	sessionStore *redistore.RediStore
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// e.Use(middleware.CSRF())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.Recover())

	e.Static("/public", "public")

	api := e.Group("/api")
	api.GET("/health", handler.HealthHandler)

	authHandler := handler.AuthHandler{
		AuthStore: s.store.AuthStore,
		Session:   s.sessionStore,
		Cfg:       s.config,
	}

	api.POST("/auth/login", authHandler.Login)
	api.GET("/auth/logout", authHandler.Logout)
	api.POST("/auth/register", authHandler.RegisterUser)
	private := api.Group("")
	private.GET("/me", authHandler.PrivateHandler)

	viewHandler := handler.ViewHandler{
		Cfg: s.config,
	}

	e.GET("/", viewHandler.HomeViewHandler)

	// hot reload
	if s.config.Env != "prod" {
		e.GET("/ws", HandleHotReload)
		NotifyClients()

	}

	return e
}

func New(cfg *config.Config, dbStore *store.Store, sessionStore *redistore.RediStore) (*http.Server, error) {

	serverAddr, err := cfg.AddressString()
	if err != nil {
		return nil, err
	}

	newServer := &Server{
		config:       cfg,
		store:        dbStore,
		sessionStore: sessionStore,
	}

	server := &http.Server{
		Addr:         serverAddr,
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
