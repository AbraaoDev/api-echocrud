package handler

import (
	"echocrud/internal/entities"
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

func (e *establishmentHandler) CreateEstablishment(c echo.Context) error {
	var establishment entities.Establishment

	err := c.Bind(&establishment)
	if(err != nil){
		return c.JSON(http.StatusBadRequest, err)	 
	}

	insertEstablishment, err := e.establishmentService.CreateEstablishment(establishment)

	if(err != nil){
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, insertEstablishment)

}