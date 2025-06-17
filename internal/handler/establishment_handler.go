package handler

import (
	"echocrud/internal/entity"
	"echocrud/internal/service"
	"net/http"
	"strconv"

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
		response := entity.Response{
			Message: "Internal Server Error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, ests)

}

func (e *establishmentHandler) CreateEstablishment(c echo.Context) error {
	var establishment entity.Establishment

	err := c.Bind(&establishment)
	if(err != nil){
		response := entity.Response{
			Message: "Invalid JSON format or missing required fields",
		}
		return c.JSON(http.StatusBadRequest, response)	 
	}

	insertEstablishment, err := e.establishmentService.CreateEstablishment(establishment)

	if(err != nil){
		response := entity.Response{
			Message: "Internal Server Error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusCreated, insertEstablishment)

}

func (e *establishmentHandler) GetEstablishmentById(c echo.Context) error {
	id := c.Param("establishmentId")
	if (id == "") {
		response := entity.Response{
			Message: "EstablishmentId cannot be null",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	establishmentId, err := strconv.ParseUint(id, 10, 32)
	if(err != nil ){
		response := entity.Response{
			Message: "EstablishmentId must be a number",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	estab, err := e.establishmentService.GetEstablishmentById(uint(establishmentId))
	if err != nil {
		response := entity.Response{
			Message: "Internal Server Error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	if estab == nil {
		response := entity.Response{
			Message: "Establishment not found.",
		}
		return c.JSON(http.StatusNotFound, response)
	}
	return c.JSON(http.StatusOK, estab)

}