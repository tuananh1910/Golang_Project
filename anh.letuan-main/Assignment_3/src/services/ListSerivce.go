package services

import (
	"Assignment_3/src/models/entity"
	"Assignment_3/src/repository"
)

type ListService interface {
	Save(list []entity.List) []entity.List
	FindListById(id uint32) entity.List
}

type listService struct {
	listRepo repository.ListRepository
}

func (l *listService) FindListById(id uint32) entity.List {
	return l.listRepo.FindListById(id)
}

func (l *listService) Save(list []entity.List) []entity.List {
	return l.listRepo.Save(list)
}

func NewListService(listRepo repository.ListRepository) ListService {
	return &listService{listRepo: listRepo}
}
