package entity

import "reflect"

type WeatherForecast struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`

	CityID string `gorm:"column:city_id"`
	//ListId uint32 `gorm:"column:listId"`
	//Lists *[]List `json:"lists, omitempty"`

	City City       `gorm:"foreignKey:CityID;references:id" json:"city"`
	List []List `gorm:"foreignKey:WeatherForecastId;references:id" json:"list, omitempty"`

}

func (w WeatherForecast) IsEmpty() bool {
	return reflect.DeepEqual(w,WeatherForecast{})
}