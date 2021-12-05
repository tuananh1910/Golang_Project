package database
/*
	go get -u github.com/jinzhu/gorm
	go get -u github.com/go-sql-driver/mysql
*/
import (
	"awesomeProject/src/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		return nil, err
	}

	return db, nil
}
