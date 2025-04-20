package handler

import (
	"barrytime/go_templ_boilerplate/internal/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// standard response
func apiDataResponse(c echo.Context, status int, data map[string]interface{}) error {
	return c.JSON(status, echo.Map{"error": false, "data": data})
}

// standard response
func apiMessageResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, echo.Map{"error": false, "message": message})
}

// standard error
func errAPIResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, echo.Map{"error": true, "message": message})
}

// decode request body
func decode[T model.Validator](r *http.Request) (T, error) {
	var t T
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		return t, fmt.Errorf("failed to decode request body: %w", err)
	}
	if err := t.Validate(); err != nil {
		return t, fmt.Errorf("request validation failed: %w", err)
	}
	return t, nil
}

// render
func render(c echo.Context, t templ.Component) error {
	return t.Render(c.Request().Context(), c.Response().Writer)
}
