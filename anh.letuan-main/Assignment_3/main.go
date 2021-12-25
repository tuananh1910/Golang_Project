package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Assignment_3/src/config"
	"Assignment_3/src/controllers"
	"Assignment_3/src/middlewares"
	"Assignment_3/src/repository"
	"Assignment_3/src/services"
)

var (
	db *gorm.DB = config.SetupDBConnection()

	accountRepository repository.AccountRepository = repository.NewAccountRepo(db)
	jwtService        services.JWTService          = services.NewJWTService()
	authService       services.AuthService         = services.NewAuthService(accountRepository)
	authController    controllers.AuthController   = controllers.NewAuthController(authService, jwtService)

	weatherForecastRepo    repository.WeatherForecastRepo        = repository.NewWeatherForecastRepo(db)
	weatherForecastService services.WeatherForecastService       = services.NewWeatherForecastService(weatherForecastRepo)
	weatherController      controllers.WeatherForecastController = controllers.NewWeatherForecastController(weatherForecastService)

	listRepository repository.ListRepository  = repository.NewListRepository(db)
	listService    services.ListService       = services.NewListService(listRepository)
	listController controllers.ListController = controllers.NewListController(listService)

	supportRepo       repository.SupportRepository  = repository.NewSupportRepository(db)
	supportService    services.SupportService       = services.NewSupportService(supportRepo)
	supportController controllers.SupportController = controllers.NewSupportController(supportService)
)

func main() {

	route := gin.Default()

	authRoutes := route.Group("api/auth")
	{
		authRoutes.POST("/login", authController.LogIn)
		authRoutes.POST("/register", authController.Register)
	}

	homeRoutes := route.Group("api/weather-forecast")
	{
		homeRoutes.GET("/Sai-Gon", weatherController.FindWeatherForecastByNameCity)
	}

	weatherForecastRotues := route.Group("api/weather-forecast", middlewares.AuthorizeJWT(jwtService))
	{
		weatherForecastRotues.GET("/", weatherController.FindWeatherForecastByNameCity)
		weatherForecastRotues.GET("/details-weather-forecast-a day", listController.FindListById)

		weatherForecastRotues.POST("/support", supportController.Save)

		// thiet lap lai don vi nhiet do
		weatherForecastRotues.POST("/settings", weatherController.Setting)

	}

	_ = route.Run(":8080")

}
