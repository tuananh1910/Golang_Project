package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"

	"Assignment_3/src/models/entity"
)

func SetupDBConnection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Fail load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	//dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3307)/%s?charset=utf8mb4&parseTime=True",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connect DB fail")
	}
	db.AutoMigrate(
		&entity.Account{},
		&entity.WeatherForecast{},
		&entity.Temp{},
		&entity.City{},
		&entity.List{})
	return db
}
