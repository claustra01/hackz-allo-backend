package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hackz-allo/database"
)

func Auth(c echo.Context) error {

	db := database.Connect()
	token := c.QueryParam("token")

	// 認証チェック
	array := []database.User{}
	db.Find(&array)
	for _, u := range array {
		if u.Id.String() == token {
			return c.String(http.StatusOK, "OK")
		}
	}
	return c.String(http.StatusOK, "Failed")
}
