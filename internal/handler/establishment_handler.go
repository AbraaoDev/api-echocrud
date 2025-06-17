package handler

import (
	"echocrud/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type establishmentHandler struct {
	establishmentService service.EstablishmentService
}

func NewEstablishmentHandler(service service.EstablishmentService) establishmentHandler {
	return establishmentHandler{
		establishmentService: service,
	}
}

func (e *establishmentHandler) GetAll(c echo.Context) error {
	ests, err := e.establishmentService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, ests)

}
