package database
/*
	go get -u github.com/jinzhu/gorm

*/
import (
	"awesomeProject/src/config"
	"github.com/jinzhu/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		return nil, err
	}

	return db, nil
}
