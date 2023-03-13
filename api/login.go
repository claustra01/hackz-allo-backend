package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hackz-allo/database"
)

func LogIn(c echo.Context) error {

	db := database.Connect()
	obj := new(database.Response)

	// クエリ展開
	id := c.FormValue("user_id")
	password := c.FormValue("password")

	// ログイン判定
	array := []database.User{}
	db.Find(&array)
	for _, u := range array {
		if u.UserId == id {
			if u.Password == password {
				// 成功
				obj.Status = true
				obj.Message = "Login successful!"
				return c.JSON(http.StatusOK, obj)
			} else {
				// パスワードが違う時
				obj.Status = false
				obj.Message = "Password is incorrect."
				return c.JSON(http.StatusOK, obj)
			}
		}
	}

	// ユーザーが見つからない時
	obj.Status = false
	obj.Message = "User not found."
	return c.JSON(http.StatusOK, obj)
}
