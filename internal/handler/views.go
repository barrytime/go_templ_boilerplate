package handler

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"barrytime/go_templ_boilerplate/internal/store"
	"barrytime/go_templ_boilerplate/internal/views"

	"github.com/labstack/echo/v4"
)

type ViewHandler struct {
	Cfg       *config.Config
	AuthStore *store.AuthStore
}

func (h *ViewHandler) HomeViewHandler(c echo.Context) error {
	return render(c, views.Home(h.Cfg.Env))
}
