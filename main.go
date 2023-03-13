package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hackz-allo/api"
)

func gormConnect() *gorm.DB {
	dsn := "host=db user=postgres password=password dbname=database port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error(Init): ", err)
	} else {
		fmt.Println("DB Connected!")
	}
	return db
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	var a []int
	db := gormConnect()
	db.Find(&a)

	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CLIENT_URL_LOCAL"), os.Getenv("CLIENT_URL_REMOTE")},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
	}))

	// ルートを設定
	e.GET("/initialization", api.Initialization)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
