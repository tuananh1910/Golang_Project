package repository

import (
	"gorm.io/gorm"

	"Assignment_3/src/models/entity"
)

type CityRepository interface {
	Save(city entity.City) entity.City
}

type cityRepository struct {
	connect *gorm.DB
}

func (c *cityRepository) Save(city entity.City) entity.City {
	c.connect.Save(&city)
	return city
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityRepository{connect: db}
}
