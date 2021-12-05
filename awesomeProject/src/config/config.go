package config
/*
	go get -u github.com/joho/godotenv
*/
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	PORT = 0
	DBDRIVER = ""
	DBURL = ""
)

func Load() {
	var err error
	err = godotenv.Load()
	if err!= nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err!= nil {
		log.Println(err)
		PORT = 9000
	}

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

}
