package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"Assignment_3/src/models/entity"
	"Assignment_3/src/response"
	"Assignment_3/src/services"
)

type WeatherForecastController interface {
	FindWeatherForecastByNameCity(ctx *gin.Context)
	GetAllWeatherForecast(ctx *gin.Context)
	Setting(ctx *gin.Context)
}

type weatherForecastController struct {
	weatherForecastService services.WeatherForecastService
	cityService            services.CityService
	listService            services.ListService
}

func (w *weatherForecastController) Setting(ctx *gin.Context) {
	panic("implement me")
}

// call open api and save data from open api to DB
func (w *weatherForecastController) GetAllWeatherForecast(ctx *gin.Context) {
	var err error
	var city string
	city = "hanoi"
	//city = ctx.Query("name_city")

	url := "https://community-open-weather-map.p.rapidapi.com/climate/month?q=" + city

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "ff6e0bdda3mshe20900192136e43p14e5cajsnb9a65a10ad76")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		res := response.BuildErrResponse("Please enter city again", "City Not Found", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	fmt.Println(res)
	fmt.Println(string(body))

	var weatherForecasts = entity.WeatherForecast{}
	err = json.Unmarshal(body, &weatherForecasts)

	if err != nil {
		res := response.BuildResponse(
			true, "Fail", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

	} else {
		w.weatherForecastService.GetAllWeatherForecast(weatherForecasts)
	}

}

// get data weather forecast follow name city (param name_city in url)
func (w *weatherForecastController) FindWeatherForecastByNameCity(ctx *gin.Context) {
	var nameCity string
	nameCity = ctx.Query("name_city")

	if nameCity == "" {
		res := response.BuildErrResponse("Please enter city again", "City Not Found", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	weatherForecasts := w.weatherForecastService.FindWeatherForecastByNameCity(nameCity)
	if weatherForecasts.IsEmpty() {
		w.GetAllWeatherForecast(ctx) // if city not exists in DB
		w.FindWeatherForecastByNameCity(ctx)
	}

	res := response.BuildResponse(true, "OK", weatherForecasts)
	ctx.JSON(http.StatusOK, res)

}

func NewWeatherForecastController(wFS services.WeatherForecastService) WeatherForecastController {
	return &weatherForecastController{weatherForecastService: wFS}
}
