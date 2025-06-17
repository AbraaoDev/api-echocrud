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