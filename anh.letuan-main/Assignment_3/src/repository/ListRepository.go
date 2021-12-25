package repository

import (
	"gorm.io/gorm"

	"Assignment_3/src/models/entity"
)

type ListRepository interface {
	Save(list []entity.List) []entity.List
	FindListById(id uint32) entity.List
}

type listRepository struct {
	connect *gorm.DB
}

func (l *listRepository) FindListById(id uint32) entity.List {
	var list entity.List
	l.connect.Where("id=?", id).Find(&list)
	return list
}

func (l *listRepository) Save(list []entity.List) []entity.List {
	l.Save(list)
	return list
}

func NewListRepository(db *gorm.DB) ListRepository {
	return &listRepository{connect: db}
}
