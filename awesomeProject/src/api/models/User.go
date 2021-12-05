package models

type User struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"size:20;not null;unique" json:"username"`
	Email string `gorm:"size:50;not null; unique" json:"email"`
	Password string `gorm:"size:60;not null" json:"password"`

}
