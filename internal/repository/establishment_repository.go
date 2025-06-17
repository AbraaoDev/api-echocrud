package repository

import (
	"echocrud/internal/entity"

	"gorm.io/gorm"
)

type EstablishmentRepository struct {
	db *gorm.DB
}

func NewEstablishmentRepository(db *gorm.DB) EstablishmentRepository {
	return EstablishmentRepository{
		db: db,
	}
}

func (er *EstablishmentRepository) GetAll() ([]entity.Establishment, error) {
    var establishments []entity.Establishment
    err := er.db.Find(&establishments).Error
    return establishments, err
}

func (er *EstablishmentRepository) GetEstablishmentById(id_establishment uint) (*entity.Establishment, error) {
	var establishment entity.Establishment
	err := er.db.Find(&establishment, id_establishment).Error
	if err != nil {
		return nil, err
	}

	return &establishment, nil
}

func (er *EstablishmentRepository) CreateEstablishment(establishment *entity.Establishment) (uint, error) {
	err := er.db.Create(establishment).Error
	if err != nil {
		return 0, err
	}

	return establishment.ID, nil
}