package router

import (
	"echocrud/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(server *echo.Echo, establishmentHandler *handler.EstablishmentHandler, storeHandler *handler.StoreHandler) {
	// prefixo
	v1 := server.Group("/api/v1")

	establishmentsGroup := v1.Group("/establishments")
	{
		establishmentsGroup.POST("", establishmentHandler.CreateEstablishment)
		establishmentsGroup.GET("", establishmentHandler.GetAll)
		establishmentsGroup.GET("/:establishmentId", establishmentHandler.GetEstablishmentById)
		establishmentsGroup.PUT("/:establishmentId", establishmentHandler.UpdateEstablishment)
		establishmentsGroup.DELETE("/:establishmentId", establishmentHandler.DeleteEstablishment)

		storesNestedGroup := establishmentsGroup.Group("/:establishmentId/stores")
		{
			storesNestedGroup.POST("", storeHandler.CreateStore)
			storesNestedGroup.GET("", storeHandler.GetAllStoresByEstablishment)
			storesNestedGroup.GET("/:storeId", storeHandler.GetStoreByID)
			storesNestedGroup.PUT("/:storeId", storeHandler.UpdateStore)
			storesNestedGroup.DELETE("/:storeId", storeHandler.DeleteStore)
		}
	}

}