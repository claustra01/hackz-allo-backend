package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Initialization(c echo.Context) error {

	db := database.Connect()

	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS posts")

	db.AutoMigrate(database.User{})
	db.AutoMigrate(database.Post{})

	return c.String(http.StatusOK, "Initializaton")
}
