package handler

import (
	"barrytime/go_templ_boilerplate/internal/config"
	"barrytime/go_templ_boilerplate/internal/model"
	"barrytime/go_templ_boilerplate/internal/store"
	"net/http"

	"github.com/boj/redistore"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthStore *store.AuthStore
	Session   *redistore.RediStore
	Cfg       *config.Config
}

func (h *AuthHandler) Login(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return errAPIResponse(c, http.StatusMethodNotAllowed, "invalid request method")
	}

	req, err := decode[*model.LoginRequest](c.Request())
	if err != nil {
		return errAPIResponse(c, http.StatusBadRequest, err.Error())
	}
	defer c.Request().Body.Close()

	existingUser, err := h.AuthStore.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return errAPIResponse(c, http.StatusUnauthorized, "unauthorized")
	}

	if err := existingUser.ComparePassword(req.Password); err != nil {
		return errAPIResponse(c, http.StatusUnauthorized, "unauthorized")
	}

	session, _ := h.Session.Get(c.Request(), h.Cfg.SessionName)
	session.Values["user"] = existingUser
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return errAPIResponse(c, http.StatusInternalServerError, "failed to save session")
	}

	return apiMessageResponse(c, http.StatusAccepted, "login success")
}

func (h *AuthHandler) Logout(c echo.Context) error {
	session, _ := h.Session.Get(c.Request(), h.Cfg.SessionName)
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
	return apiMessageResponse(c, http.StatusAccepted, "logout success")
}

func (h *AuthHandler) RegisterUser(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return errAPIResponse(c, http.StatusMethodNotAllowed, "invalid request method")
	}

	req, err := decode[*model.NewUserRequest](c.Request())
	if err != nil {
		return errAPIResponse(c, http.StatusBadRequest, err.Error())
	}
	defer c.Request().Body.Close()

	user, err := h.AuthStore.CreateUser(c.Request().Context(), req)
	if err != nil {
		return errAPIResponse(c, http.StatusBadRequest, "error creating user")
	}

	session, _ := h.Session.Get(c.Request(), h.Cfg.SessionName)
	session.Values["user"] = user
	if err := session.Save(c.Request(), c.Response()); err != nil {
		return errAPIResponse(c, http.StatusInternalServerError, "failed to save session")
	}

	return apiDataResponse(c, http.StatusCreated, echo.Map{
		"user": user,
	})
}

func (h *AuthHandler) PrivateHandler(c echo.Context) error {
	user := c.Get("user").(*model.User)
	return apiDataResponse(c, http.StatusOK, echo.Map{
		"user": user,
	})
}
