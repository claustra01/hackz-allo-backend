package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hackz-allo/database"
)

func Initialization(c echo.Context) error {

	db := database.Connect()

	db.Exec("DROP TABLE IF EXISTS users")

	db.AutoMigrate(database.User{})

	return c.String(http.StatusOK, "Initializaton")
}
