package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hackz-allo/database"
)

type oGetUser struct {
	UserId string
	Name   string
}

func GetUser(c echo.Context) error {

	db := database.Connect()
	token := c.QueryParam("token")

	// ユーザー情報取得
	array := []database.User{}
	db.Find(&array)
	for _, u := range array {
		if u.Id.String() == token {
			obj := new(oGetUser)
			obj.UserId = u.UserId
			obj.Name = u.Name
			return c.JSON(http.StatusOK, obj)
		}
	}
	return c.JSON(http.StatusOK, nil)
}
