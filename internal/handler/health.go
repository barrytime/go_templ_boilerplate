package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthHandler(c echo.Context) error {
	return apiMessageResponse(c, http.StatusOK, "OK")
}
