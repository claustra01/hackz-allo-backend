package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeletePost(c echo.Context) error {

	db := database.Connect()

	// クエリ展開
	id := c.QueryParam("id")

	// 投稿削除

	// OKを返す
	return c.String(http.StatusOK, "OK")
}
