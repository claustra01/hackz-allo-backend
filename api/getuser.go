package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {

	type json struct {
		UserId string
		Name   string
	}

	db := database.Connect()
	token := c.QueryParam("token")

	// ユーザー情報取得
	array := []database.User{}
	db.Find(&array)
	for _, u := range array {
		if u.Id.String() == token {
			obj := new(json)
			obj.UserId = u.UserId
			obj.Name = u.Name
			return c.JSON(http.StatusOK, obj)
		}
	}
	return c.JSON(http.StatusOK, nil)
}
