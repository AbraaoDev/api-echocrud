package handler

import (
	"echocrud/internal/entity"
	"echocrud/internal/service"
	"errors"
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
	if err != nil {
		response := entity.Response{
			Message: "Invalid JSON format or missing required fields",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	insertEstablishment, err := e.establishmentService.CreateEstablishment(establishment)

	if err != nil {
		response := entity.Response{
			Message: "Internal Server Error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusCreated, insertEstablishment)

}

func (e *establishmentHandler) GetEstablishmentById(c echo.Context) error {
	id := c.Param("establishmentId")
	if id == "" {
		response := entity.Response{
			Message: "EstablishmentId cannot be null",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	establishmentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := entity.Response{
			Message: "EstablishmentId must be a number",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	estab, err := e.establishmentService.GetEstablishmentById(uint(establishmentId))
	if err != nil {
		if errors.Is(err, service.ErrEstablishmentNotFound) {
			response := entity.Response{
				Message: "Establishment not found.",
			}
			return c.JSON(http.StatusNotFound, response)
		}
		response := entity.Response{
			Message: "Internal Server Error",
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, estab)

}

func (e *establishmentHandler) DeleteEstablishment(c echo.Context) error {
	id := c.Param("establishmentId")
	if id == "" {
		response := entity.Response{
			Message: "EstablishmentId cannot be null",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	estab, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := entity.Response{
			Message: "EstablishmentId must be a number",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = e.establishmentService.DeleteEstablishment(uint(estab))
	if err != nil {
		if errors.Is(err, service.ErrEstablishmentNotFound) {
			response := entity.Response{
				Message: "Establishment not found.",
			}
			return c.JSON(http.StatusNotFound, response)
		} else if errors.Is(err, service.ErrEstablishmentHasStores) {
			response := entity.Response{
				Message: err.Error(),
			}
			return c.JSON(http.StatusConflict, response)
		} else {
			response := entity.Response{
				Message: "Internal Server Error",
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	return c.JSON(http.StatusOK, estab)
}


func (e *establishmentHandler) UpdateEstablishment(c echo.Context) error {
	id := c.Param("establishmentId")
	if id == "" {
		response := entity.Response{Message: "EstablishmentId cannot be null"}
		return c.JSON(http.StatusBadRequest, response)
	}

	establishmentId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := entity.Response{Message: "EstablishmentId must be a number"}
		return c.JSON(http.StatusBadRequest, response)
	}

	var establishmentToUpdate entity.Establishment
	if err := c.Bind(&establishmentToUpdate); err != nil {
		response := entity.Response{Message: "Invalid JSON format or missing required fields"}
		return c.JSON(http.StatusBadRequest, response)
	}

	updatedEstablishment, err := e.establishmentService.UpdateEstablishment(uint(establishmentId), establishmentToUpdate)
	if err != nil {
		if errors.Is(err, service.ErrEstablishmentNotFound) {
			response := entity.Response{
				Message: "Establishment not found.",
			}
			return c.JSON(http.StatusNotFound, response)
		}
		response := entity.Response{Message: "Internal Server Error"}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, updatedEstablishment)
}