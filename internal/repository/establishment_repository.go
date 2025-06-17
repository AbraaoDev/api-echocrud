package repository

import (
	"echocrud/internal/entities"

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

func (er *EstablishmentRepository) GetAll() ([]entities.Establishment, error) {
    var establishments []entities.Establishment
    err := er.db.Find(&establishments).Error
    return establishments, err
}

func (er *EstablishmentRepository) CreateEstablishment(establishment *entities.Establishment) (uint, error) {
	err := er.db.Create(establishment).Error
	if err != nil {
		return 0, err
	}

	return establishment.ID, nil
}