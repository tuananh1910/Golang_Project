package repository

import (
	"gorm.io/gorm"

	"Assignment_3/src/models/entity"
)

type WeatherForecastRepo interface {
	FindWeatherForecastByNameCity(nameCity string) entity.WeatherForecast
	SaveAll(weatherForecast entity.WeatherForecast)
}

type weatherForecastConnectDB struct {
	connect *gorm.DB
}

func (w *weatherForecastConnectDB) SaveAll(weatherForecast entity.WeatherForecast) {
	w.connect.Save(&weatherForecast)
}

func (w *weatherForecastConnectDB) FindWeatherForecastByNameCity(nameCity string) entity.WeatherForecast {

	var weatherForecast entity.WeatherForecast
	var city entity.City
	var temp entity.Temp
	var list []entity.List

	w.connect.Where("city_name=?", nameCity).Find(&weatherForecast)

	w.connect.Where("name =?", nameCity).Find(&city)
	weatherForecast.City = city

	w.connect.Where("weather_forecast_id = ?", weatherForecast.ID).Find(&list)
	weatherForecast.List = list

	var listTemp []entity.List
	for _, valueList := range list {

		w.connect.Where("id=?", valueList.TempId).Find(&temp)

		valueList.Temps = temp

		listTemp = append(listTemp, valueList)

	}
	weatherForecast.List = listTemp

	return weatherForecast
}

func NewWeatherForecastRepo(db *gorm.DB) WeatherForecastRepo {
	return &weatherForecastConnectDB{connect: db}
}
