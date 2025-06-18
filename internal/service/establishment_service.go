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

	if establishment == nil {
		return nil, ErrEstablishmentNotFound
	}

	return establishment, err
}

func (es *EstablishmentService) DeleteEstablishment(id uint) error {
	// checar se existe o estabelecimento
	existingEstablishment, err := es.repository.GetEstablishmentById(id)
	if err != nil {
		return err
	}
	
	if existingEstablishment == nil {
		return ErrEstablishmentNotFound
	}

	hasStores, err := es.repository.HasStores(id)
	if err != nil {
		return err
	}

	if hasStores {
		return ErrEstablishmentHasStores  
	}

	return es.repository.DeleteEstablishment(id)
}


func (es *EstablishmentService) UpdateEstablishment(id uint, establishmentData entity.Establishment) (*entity.Establishment, error) {
	existingEstablishment, err := es.repository.GetEstablishmentById(id)
	if err != nil {
		return nil, err 
	}

	if existingEstablishment == nil {
		return nil, ErrEstablishmentNotFound 
	}

	err = es.repository.UpdateEstablishment(id, &establishmentData)
	if err != nil {
		return nil, err 
	}

	updatedEstablishment, err := es.repository.GetEstablishmentById(id)
	if err != nil {
		return nil, err 
	}

	return updatedEstablishment, nil
}
