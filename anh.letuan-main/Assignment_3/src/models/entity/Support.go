package entity

type Support struct {
	ID uint32 `gorm:"primary_key: auto_increment" json:"id"`
	Description string `json:"description"`

	IdUser uint32 `gorm:"column:id_user" json:"id_user"`
}
