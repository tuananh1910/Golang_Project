package entity

// temperature
type Temp struct {
	ID uint32 `gorm:"primary_key:auto_increment"json:"id"`

	Average float32 `json:"average, omitempty"`
	AverageMax float32 `json:"average_max, omitempty"`
	AverageMin float32 `json:"average_min, omitempty"`
	RecordMax float32 `json:"record_max, omitempty"`
	RecordMin float32 `json:"record_min, omitempty"`
	//WeatherForecast  WeatherForecast `gorm:"foreignkey:weatherForecastId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"weatherForecast"`

}
