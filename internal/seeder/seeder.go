package seeder

import (
	"echocrud/internal/entity"
	"errors"
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("Starting seeder...")

	establishments := []entity.Establishment{
		{
			Name:            "Matriz Varejo S.A.",
			CorporateNumber: "11.222.333/0001-44",
			CorporateName:   "Grande Varejista do Brasil S.A.",
			Address:         "Avenida Principal",
			City:            "São Paulo",
			State:           "SP",
			ZipCode:         "01000-000",
			Number:          "100",
			Stores: []entity.Store{
				{
					Name:        "Loja Principal Centro",
					StoreNumber: "SP-001",
					StoreName:   "Varejo Center",
					Address:     "Rua do Comércio",
					City:        "São Paulo",
					State:       "SP",
					ZipCode:     "01001-001",
					Number:      "50",
				},
				{
					Name:        "Loja Zona Sul",
					StoreNumber: "SP-002",
					StoreName:   "Varejo Sul",
					Address:     "Avenida das Nações",
					City:        "São Paulo",
					State:       "SP",
					ZipCode:     "04578-000",
					Number:      "1500",
				},
			},
		},
		{
			Name:            "Filial Nordeste Comércio",
			CorporateNumber: "44.555.666/0001-77",
			CorporateName:   "Comércio Nordestino Ltda.",
			Address:         "Avenida Litorânea",
			City:            "Recife",
			State:           "PE",
			ZipCode:         "50010-000",
			Number:          "250",
			Stores: []entity.Store{
				{
					Name:        "Loja Boa Viagem",
					StoreNumber: "PE-001",
					StoreName:   "Comércio Boa Viagem",
					Address:     "Avenida Boa Viagem",
					City:        "Recife",
					State:       "PE",
					ZipCode:     "51020-000",
					Number:      "999",
				},
			},
		},
	}

	for _, est := range establishments {
		var existingEstablishment entity.Establishment
		err := db.Where("corporate_number = ?", est.CorporateNumber).First(&existingEstablishment).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&est).Error; err != nil {
				log.Printf("Error seeding establishment %s: %v\n", est.Name, err)
			} else {
				log.Printf("Successfully seeded establishment: %s\n", est.Name)
			}
		} else if err != nil {
			log.Printf("Database error while checking establishment %s: %v\n", est.Name, err)
		} else {
			log.Printf("Establishment '%s' (CNPJ: %s) already exists. Skipping seed.\n", existingEstablishment.Name, existingEstablishment.CorporateNumber)
		}
	}

	log.Println("🌱 Seeder finished.")
}