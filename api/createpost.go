package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hackz-allo/database"
)

func CreatePost(c echo.Context) error {

	db := database.Connect()
	obj := new(database.Response)

	return c.JSON(http.StatusOK, obj)
}
