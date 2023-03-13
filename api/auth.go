package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hackz-allo/database"
)

func Auth(c echo.Context) error {

	db := database.Connect()

	return c.String(http.StatusOK, "obj")
}
