package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTimeLine(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
