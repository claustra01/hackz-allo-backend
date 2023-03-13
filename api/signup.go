package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Json struct {
	Status  bool
	Message string
}

func SignUp(c echo.Context) error {

	db := database.Connect()
	obj := new(Json)

	// クエリ展開
	id := c.FormValue("user_id")
	name := c.FormValue("name")
	password := c.FormValue("password")

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
		obj.Status = false
		obj.Message = "This id is used."
		return c.JSON(http.StatusOK, obj)
	}

	// ユーザー登録
	user := new(database.User)
	user.UserId = id
	user.Name = name
	user.Password = password
	db.Create(&user)

	obj.Status = true
	obj.Message = id + " is registered!"
	return c.JSON(http.StatusOK, obj)
}
