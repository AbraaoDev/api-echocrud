package service

import (
	"echocrud/internal/entity"
	"echocrud/internal/repository"
)


type StoreService struct {
	storeRepository         repository.StoreRepository
	establishmentRepository repository.EstablishmentRepository
}

func NewStoreService(storeRepo repository.StoreRepository, establishmentRepo repository.EstablishmentRepository) StoreService {
	return StoreService{
		storeRepository:         storeRepo,
		establishmentRepository: establishmentRepo,
	}
}

func (ss *StoreService) CreateStore(store entity.Store) (*entity.Store, error) {
	establishment, err := ss.establishmentRepository.GetEstablishmentById(store.EstablishmentID)
	if err != nil {
		return nil, err 
	}
	if establishment == nil {
		return nil, ErrEstablishmentNotFound
	}

	storeId, err := ss.storeRepository.CreateStore(&store)
	if err != nil {
		return nil, err
	}
	store.ID = storeId
	return &store, nil
}

func (ss *StoreService) GetStoreByID(establishmentId, storeId uint) (*entity.Store, error) {
	store, err := ss.storeRepository.GetStoreByIDAndEstablishmentID(storeId, establishmentId)
	if err != nil {
		return nil, err
	}
	if store == nil {
		return nil, ErrStoreNotFound 
	}
	return store, nil
}

func (ss *StoreService) GetAllStoresByEstablishment(id_establishment uint) ([]entity.Store, error) {
	establishment, err := ss.establishmentRepository.GetEstablishmentById(id_establishment)
	if err != nil {
		return nil, err
	}
	if establishment == nil {
		return nil, ErrEstablishmentNotFound
	}
	
	return ss.storeRepository.GetAllStoresByEstablishment(id_establishment)
}

func (ss *StoreService) UpdateStore(establishmentId, storeId uint, storeData entity.Store) (*entity.Store, error) {
	existingStore, err := ss.storeRepository.GetStoreByIDAndEstablishmentID(storeId, establishmentId)
	if err != nil {
		return nil, err
	}
	if existingStore == nil {
		return nil, ErrStoreNotFound
	}

	err = ss.storeRepository.UpdateStore(storeId, &storeData)
	if err != nil {
		return nil, err
	}

	return ss.storeRepository.GetStoreByID(storeId)
}

func (ss *StoreService) DeleteStore(establishmentId, storeId uint) error {
	existingStore, err := ss.storeRepository.GetStoreByIDAndEstablishmentID(storeId, establishmentId)
	if err != nil {
		return err
	}
	if existingStore == nil {
		return ErrStoreNotFound
	}

	return ss.storeRepository.DeleteStore(storeId)
}