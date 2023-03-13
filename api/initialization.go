package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Initialization(c echo.Context) error {
	return c.String(http.StatusOK, "Initializaton")
}
