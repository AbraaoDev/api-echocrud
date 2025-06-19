package handler

import (
	"echocrud/internal/entity"
	"echocrud/internal/repository"
	"echocrud/internal/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StoreHandler struct {
	storeService service.StoreService
}

func NewStoreHandler(service service.StoreService) *StoreHandler {
	return &StoreHandler{
		storeService: service,
	}
}

func (h *StoreHandler) CreateStore(c echo.Context) error {
	establishmentIdStr := c.Param("establishmentId")
	establishmentId, err := strconv.ParseUint(establishmentIdStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "establishmentId invalid"})
	}

	var store entity.Store
	if err := c.Bind(&store); err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "Format JSON invalid"})
	}

	store.EstablishmentID = uint(establishmentId)
	createdStore, err := h.storeService.CreateStore(store)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateStoreCorporateNumber) {
			return c.JSON(http.StatusConflict, entity.Response{Message: err.Error()})
		}

		if errors.Is(err, service.ErrEstablishmentNotFound) {
			return c.JSON(http.StatusNotFound, entity.Response{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, entity.Response{Message: "Error creating store"})
	}

	return c.JSON(http.StatusCreated, createdStore)
}

func (h *StoreHandler) GetAllStoresByEstablishment(c echo.Context) error {
	establishmentIdStr := c.Param("establishmentId")
	establishmentId, err := strconv.ParseUint(establishmentIdStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "establishmentId invalid"})
	}

	stores, err := h.storeService.GetAllStoresByEstablishment(uint(establishmentId))
	if err != nil {
		if errors.Is(err, service.ErrEstablishmentNotFound) {
			return c.JSON(http.StatusNotFound, entity.Response{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, entity.Response{Message: "Error searching stores"})
	}

	return c.JSON(http.StatusOK, stores)
}

func (h *StoreHandler) GetStoreByID(c echo.Context) error {
	establishmentId, err := strconv.ParseUint(c.Param("establishmentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "establishmentId invalid"})
	}

	storeId, err := strconv.ParseUint(c.Param("storeId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "storeId invalid"})
	}

	store, err := h.storeService.GetStoreByID(uint(establishmentId), uint(storeId))
	if err != nil {
		if errors.Is(err, service.ErrStoreNotFound) {
			return c.JSON(http.StatusNotFound, entity.Response{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, entity.Response{Message: "Error searching store"})
	}

	return c.JSON(http.StatusOK, store)
}

func (h *StoreHandler) UpdateStore(c echo.Context) error {
	establishmentId, err := strconv.ParseUint(c.Param("establishmentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "establishmentId invalid"})
	}

	storeId, err := strconv.ParseUint(c.Param("storeId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "storeId invalid"})
	}

	var storeToUpdate entity.Store
	if err := c.Bind(&storeToUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "Format JSON invalid"})
	}

	updatedStore, err := h.storeService.UpdateStore(uint(establishmentId), uint(storeId), storeToUpdate)
	if err != nil {
		if errors.Is(err, service.ErrStoreNotFound) {
			return c.JSON(http.StatusNotFound, entity.Response{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, entity.Response{Message: "Error updating store"})
	}

	return c.JSON(http.StatusOK, updatedStore)
}

func (h *StoreHandler) DeleteStore(c echo.Context) error {
	establishmentId, err := strconv.ParseUint(c.Param("establishmentId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "establishmentId invalid"})
	}

	storeId, err := strconv.ParseUint(c.Param("storeId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{Message: "storeId invalid"})
	}

	err = h.storeService.DeleteStore(uint(establishmentId), uint(storeId))
	if err != nil {
		if errors.Is(err, service.ErrStoreNotFound) {
			return c.JSON(http.StatusNotFound, entity.Response{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, entity.Response{Message: "Error deleting store"})
	}

	return c.JSON(http.StatusOK, entity.Response{Message: "Store deleted successfully"})
}