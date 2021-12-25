package services

import (
	"log"

	"github.com/mashingan/smapping"

	"Assignment_3/src/models/dto"
	"Assignment_3/src/models/entity"
	"Assignment_3/src/repository"
)

type SupportService interface {
	Save(support dto.Support) entity.Support
}

type supportService struct {
	supportRepo repository.SupportRepository
}

func (s *supportService) Save(supportDTO dto.Support) entity.Support {
	support := entity.Support{}

	err := smapping.FillStruct(&support, smapping.MapFields(&supportDTO))

	if err != nil {
		log.Println("failed map")
	}

	res := s.supportRepo.Save(support)
	return res
}

func NewSupportService(supportRepository repository.SupportRepository) SupportService {
	return &supportService{supportRepo: supportRepository}
}
