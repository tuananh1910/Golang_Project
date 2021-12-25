package entity

type City struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string ` json:"name, omitempty"`
	Country string `json:"country, omitempty"`

}
