package repository

import (
	"echocrud/internal/entity"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
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

		// necessário preload, pois o lazy loading do gorm não carrega por padrão os relacionamentos
    err := er.db.Preload("Stores").Find(&establishments).Error
    return establishments, err
}

func (er *EstablishmentRepository) GetEstablishmentById(id_establishment uint) (*entity.Establishment, error) {
	var establishment entity.Establishment
	err := er.db.Preload("Stores").First(&establishment, id_establishment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { 
			return nil, nil 
		}
		return nil, err 
	}

	return &establishment, nil
}

func (er *EstablishmentRepository) CreateEstablishment(establishment *entity.Establishment) (uint, error) {
	err := er.db.Create(establishment).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, ErrDuplicateCorporateNumber
		}
		return 0, err
	}

	return establishment.ID, nil
}

func (er *EstablishmentRepository) DeleteEstablishment(id_establishment uint) error {
	return er.db.Delete(&entity.Establishment{}, id_establishment).Error
}


func (er *EstablishmentRepository) HasStores(id_establishment uint) (bool, error) {
	// -> Fazer um delete real do db
	var count int64
	err := er.db.Model(&entity.Store{}).Where("establishment_id", id_establishment).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (er *EstablishmentRepository) UpdateEstablishment(id_establishment uint, establishment *entity.Establishment) error {
	return er.db.Model(&entity.Establishment{}).Where("id = ?", id_establishment).Updates(establishment).Error
}
