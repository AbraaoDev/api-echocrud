package repository

import (
	"echocrud/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return StoreRepository{
		db: db,
	}
}

func (sr *StoreRepository) CreateStore(store *entity.Store) (uint, error) {
	err := sr.db.Create(store).Error
	if err != nil {
		return 0, err
	}
	return store.ID, nil
}

func (sr *StoreRepository) GetStoreByID(id_store uint) (*entity.Store, error) {
	var store entity.Store
	err := sr.db.First(&store, id_store).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, err
	}
	return &store, nil
}

func (sr *StoreRepository) GetAllStoresByEstablishment(id_establishment uint) ([]entity.Store, error) {
	var stores []entity.Store
	err := sr.db.Where("establishment_id = ?", id_establishment).Find(&stores).Error
	return stores, err
}

func (sr *StoreRepository) UpdateStore(id_store uint, store *entity.Store) error {
	return sr.db.Model(&entity.Store{}).Where("id = ?", id_store).Updates(store).Error
}

func (sr *StoreRepository) DeleteStore(id_store uint) error {
	return sr.db.Delete(&entity.Store{}, id_store).Error
}