package main

import (
	"echocrud/internal/db"
	"echocrud/internal/handler"
	"echocrud/internal/repository"
	"echocrud/internal/service"
	"log"

	"github.com/labstack/echo/v4"
)

func main(){
	server := echo.New()

	// Database
	database := db.NewPostgresConnection()

	// Repositories
	EstablishmentRepository := repository.NewEstablishmentRepository(database)

	// Services
	EstablishmentService := service.NewEstablishmentService(EstablishmentRepository)

	// Handlers
	EstablishmentHandler := handler.NewEstablishmentHandler(EstablishmentService)

	//Routes
	server.GET("/establishments", EstablishmentHandler.GetAll)

	log.Println("🚀 Server Running")
	server.Logger.Fatal(server.Start(":3333"))
}