package entity

type Account struct {
	ID        uint32 `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`

	Support Support `gorm:"foreignKey:IdUser; references: id" json:"support"`
}


