package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewPost(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
