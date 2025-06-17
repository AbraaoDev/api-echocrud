package service

import (
	"echocrud/internal/entities"
	"echocrud/internal/repository"
)


type EstablishmentService struct {
	repository repository.EstablishmentRepository
}

func NewEstablishmentService(repo repository.EstablishmentRepository) EstablishmentService{
	return EstablishmentService{
		repository: repo,
	}
}

func (es *EstablishmentService) GetAll() ([]entities.Establishment, error) {
	return es.repository.GetAll()
}

func (es *EstablishmentService) CreateEstablishment(establishment entities.Establishment) (entities.Establishment, error) {
	establishmentId, err := es.repository.CreateEstablishment(&establishment)
	if(err != nil){
		return entities.Establishment{}, err
	}
	establishment.ID = establishmentId
	return establishment, nil
}