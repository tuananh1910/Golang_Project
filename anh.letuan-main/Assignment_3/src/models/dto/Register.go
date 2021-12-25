package dto

type RegisterDTO struct {
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
	LastName string `json:"last_name" form:"last_name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
