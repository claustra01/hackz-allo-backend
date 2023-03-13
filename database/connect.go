package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	dsn := "postgres://" + os.Getenv("AZURE_PG_USER") + ":" + os.Getenv("AZURE_PG_PASSWORD") + "@hackzallopostgres.postgres.database.azure.com/postgres?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("===== DB Connected =====")
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	return db
}
