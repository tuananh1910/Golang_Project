package entity

type List struct {
	ID uint32 `gorm:"primary_key:auto_increment" json:"id"`

	Humidity float32 `json:"humidity, omitempty"`  // độ ẩm
	Pressure float32 `json:"pressure, omitempty"`  // áp suất (khí quyển)
	WindSpeed float32 `json:"wind_speed, omitempty"`  // tốc độ gió

	TempId uint32 `gorm:"column:temp_id"`
	Temps Temp `gorm:"foreignKey:TempId;references:id" json:"temp, omitempty"`

	WeatherForecastId uint32 `gorm:"column:weather_forecast_id"`

}
