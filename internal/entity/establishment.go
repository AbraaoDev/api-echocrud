package entity

import (
	"gorm.io/gorm"
)

type Establishment struct {
	gorm.Model             //percebi que o Id não é uma boa prática, pois gera ID's sequenciais
	Name            string `gorm:"type:varchar(255);not null" json:"name"`
	CorporateNumber string `json:"corporate_number" gorm:"unique;not null"`
	CorporateName   string `json:"corporate_name" gorm:"not null"`
	Address         string `json:"address" gorm:"not null"`
	City            string `json:"city" gorm:"not null"`
	State           string `json:"state" gorm:"not null;size:2"`
	ZipCode         string `json:"zip_code" gorm:"not null;size:10"`
	Number          string `json:"number" gorm:"not null"`
	Stores          []Store        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"stores"`
}
