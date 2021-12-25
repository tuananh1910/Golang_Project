package services

import (
	"Assignment_3/src/models/entity"
	"Assignment_3/src/repository"
)

type CityService interface {
	Save(city entity.City) entity.City
}

type cityService struct {
	cityRepo repository.CityRepository
}

func (c cityService) Save(city entity.City) entity.City {
	return c.cityRepo.Save(city)

}

func NewCityService(cityRepo repository.CityRepository) CityService {
	return &cityService{cityRepo: cityRepo}
}
