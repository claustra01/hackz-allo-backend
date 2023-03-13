package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type oSignUp struct {
	Id       string `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func SignUp(c echo.Context) error {

	db := database.Connect()
	obj := new(Json)

	// クエリ展開
	o := new(oSignUp)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id := o.Id
	name := o.Name
	password := o.Password

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
	uuidObj, _ := uuid.NewUUID()
	user := new(database.User)
	user.Id = uuidObj
	user.UserId = id
	user.Name = name
	user.Password = password
	db.Create(&user)

	obj.Result = "OK"
	obj.Message = "ID: " + id + " is registered!"
	return c.JSON(http.StatusOK, obj)
}
