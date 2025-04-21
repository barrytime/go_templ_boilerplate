package handler

import (
	"barrytime/go_templ_boilerplate/internal/model"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) SessionAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get session from Redis
		session, err := h.Session.Get(c.Request(), h.Cfg.SessionName)
		if err != nil {
			fmt.Println(err)
			return errAPIResponse(c, http.StatusInternalServerError, "failed to get session: ")
		}

		// Check for authenticated user
		authUser, ok := session.Values["user"].(*model.User)
		if !ok || authUser == nil || authUser.ID.String() == "" {
			return errAPIResponse(c, http.StatusUnauthorized, "unauthorized")
		}

		// Validate user still exists in DB
		user, err := h.AuthStore.GetUserByID(c.Request().Context(), authUser.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				// User deleted or invalid; clear session
				session.Values["user"] = nil
				session.Save(c.Request(), c.Response())
				return errAPIResponse(c, http.StatusUnauthorized, "invalid session")
			}
			return errAPIResponse(c, http.StatusInternalServerError, "failed to validate user: "+err.Error())
		}

		// Set user in context
		c.Set("user", user)

		return next(c)
	}
}

// views middleware

func (h *AuthHandler) ViewSessionAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := h.Session.Get(c.Request(), h.Cfg.SessionName)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		authUser, ok := session.Values["user"].(*model.User)
		if !ok || authUser == nil || authUser.ID.String() == "" {
			c.Set("returnURL", c.Request().URL.Path) // Store original URL
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		user, err := h.AuthStore.GetUserByID(c.Request().Context(), authUser.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				session.Values["user"] = nil
				session.Save(c.Request(), c.Response())
				return c.Redirect(http.StatusSeeOther, "/login")
			}
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		c.Set("user", user)

		return next(c)
	}
}

func (h *AuthHandler) RedirectIfAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := h.Session.Get(c.Request(), h.Cfg.SessionName)
		if err == nil {
			if authUser, ok := session.Values["user"].(*model.User); ok && authUser != nil && authUser.ID.String() != "" {
				return c.Redirect(http.StatusSeeOther, "/dashboard")
			}
		}
		return next(c)
	}
}
