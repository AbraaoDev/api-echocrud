package db

import (
	"echocrud/internal/entity"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func NewPostgresConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), 
		logger.Config{
			SlowThreshold:             time.Second,   
			LogLevel:                  logger.Info,   
			IgnoreRecordNotFoundError: true,          
			ParameterizedQueries:      false,        
			Colorful:                  true,         
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// Auto Migrate
	log.Println("🔄 Running database migrations...")
	err = db.AutoMigrate(&entity.Establishment{}, &entity.Store{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("✅ Database migrations completed!")

	return db
}