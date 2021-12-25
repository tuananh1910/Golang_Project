package repository

import (
	"gorm.io/gorm"

	"Assignment_3/src/models/entity"
)

type SupportRepository interface {
	Save(support entity.Support) entity.Support
}

type supportRepository struct {
	connect *gorm.DB
}

func (s *supportRepository) Save(support entity.Support) entity.Support {
	s.connect.Save(support)
	return support
}

func NewSupportRepository(connect *gorm.DB) SupportRepository {
	return &supportRepository{connect: connect}
}
