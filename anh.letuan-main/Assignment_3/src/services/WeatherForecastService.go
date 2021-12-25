package services

import (
	"log"

	"github.com/mashingan/smapping"

	"Assignment_3/src/models/entity"
	"Assignment_3/src/repository"
)

type WeatherForecastService interface {
	FindWeatherForecastByNameCity(nameCity string) entity.WeatherForecast
	GetAllWeatherForecast(weatherForecast entity.WeatherForecast)
}

type weatherForecastService struct {
	weatherForecastRepo repository.WeatherForecastRepo
}

func (w *weatherForecastService) GetAllWeatherForecast(weatherForecastDTO entity.WeatherForecast) {
	//"save all service"
	weatherForecast := entity.WeatherForecast{}

	err := smapping.FillStruct(&weatherForecast, smapping.MapFields(&weatherForecastDTO))

	if err != nil {
		log.Println("failed map")
	}
	w.weatherForecastRepo.SaveAll(weatherForecast)

}

func (w *weatherForecastService) FindWeatherForecastByNameCity(nameCity string) entity.WeatherForecast {
	return w.weatherForecastRepo.FindWeatherForecastByNameCity(nameCity)
}

func NewWeatherForecastService(weatherForecastRepo repository.WeatherForecastRepo) WeatherForecastService {
	return &weatherForecastService{weatherForecastRepo: weatherForecastRepo}
}
