package entity

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name            string `gorm:"type:varchar(255);not null" json:"name"`
	StoreNumber     string `json:"store_number" gorm:"unique;not null"`
	StoreName       string `json:"store_name" gorm:"not null"`
	Address         string `json:"address" gorm:"not null"`
	City            string `json:"city" gorm:"not null"`
	State           string `json:"state" gorm:"not null;size:2"`
	ZipCode         string `json:"zip_code" gorm:"not null;size:10"`
	Number          string `json:"number" gorm:"not null"`
	EstablishmentID uint   `gorm:"not null" json:"establishment_id"`
}
