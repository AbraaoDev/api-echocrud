package db

import (
	"echocrud/internal/entities"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() *gorm.DB {
    dbHost := "localhost"
    dbPort := "5432"
    dbUser := "docker"
    dbPassword := "docker"
    dbName := "echocrud"

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
        dbHost, dbUser, dbPassword, dbName, dbPort,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // Auto Migrate
    err = db.AutoMigrate(&entities.Establishment{}, &entities.Store{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    return db
}
