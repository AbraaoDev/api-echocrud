package service

import (
	"echocrud/internal/entity"
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

func (es *EstablishmentService) GetAll() ([]entity.Establishment, error) {
	return es.repository.GetAll()
}

func (es *EstablishmentService) CreateEstablishment(establishment entity.Establishment) (entity.Establishment, error) {
	establishmentId, err := es.repository.CreateEstablishment(&establishment)
	if(err != nil){
		return entity.Establishment{}, err
	}
	establishment.ID = establishmentId
	return establishment, nil
}


func (es *EstablishmentService) GetEstablishmentById(id_establishment uint) (*entity.Establishment, error) {
	establishment, err := es.repository.GetEstablishmentById(id_establishment)
	if(err != nil){
		return nil, err
	}

	return establishment, err
}
