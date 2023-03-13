package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {

	db := database.Connect()
	obj := new(database.Response)

	// クエリ展開
	id := c.QueryParam("user_id")
	name := c.QueryParam("name")
	password := c.QueryParam("password")

	// ユーザー名重複チェック
	array := []database.User{}
	db.Find(&array)
	dup := false
	for _, u := range array {
		if u.UserId == id {
			dup = true
		}
	}
	if dup {
		obj.Result = "Failed"
		obj.Message = "ID: " + id + " is used."
		return c.JSON(http.StatusOK, obj)
	}

	// ユーザー登録
	user := new(database.User)
	user.UserId = id
	user.Name = name
	user.Password = password
	db.Create(&user)

	obj.Result = "OK"
	obj.Message = id + " is registered!"
	return c.JSON(http.StatusOK, obj)
}
