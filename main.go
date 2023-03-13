package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"hackz-allo/api"
)

func gormConnect() *gorm.DB {
	dsn := "postgres://" + os.Getenv("AZURE_PG_USER") + ":" + os.Getenv("AZURE_PG_PASSWORD") + "@hackzallopostgres.postgres.database.azure.com/postgres?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {

	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// DB初期設定
	db := gormConnect()
	db.Logger = db.Logger.LogMode(logger.Info)

	// インスタンス作成
	e := echo.New()

	// ミドルウェア設定
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

	// DB初期化 開発用
	e.GET("/api/initialization", api.Initialization)

	// 認証
	e.GET("/api/login", api.LogIn)
	e.GET("/api/signup", api.SignUp)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
