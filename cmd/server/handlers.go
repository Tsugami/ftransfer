package main

import (
	"net/http"

	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	storageProviderService *storage_provider.StorageProviderService
}

func (h *Handler) CreateStorageProvider(c *gin.Context) {
	var input CreateStorageProviderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is not valid"})
		return
	}

	storage, err := h.storageProviderService.Create(
		c.Request.Context(),
		input.Name,
		input.FileSystem,
		input.ProtocolConnection,
	)

	if err != nil {
		c.Error(err)
		return
	}

	var response StorageProviderResponse = NewStorageProviderResponse(storage)
	c.JSON(http.StatusCreated, response)
}

func (h *Handler) ListStorageProviders(c *gin.Context) {
	storageProviders, err := h.storageProviderService.List(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	response := NewStorageProviderResponseList(storageProviders)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetStorageProvider(c *gin.Context) {
	id := storage_provider.ID(c.Param("id"))

	storageProvider, err := h.storageProviderService.Get(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	response := NewStorageProviderResponse(storageProvider)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateStorageProvider(c *gin.Context) {
	id := storage_provider.ID(c.Param("id"))
	var input UpdateStorageProviderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is not valid"})
		return
	}

	err := h.storageProviderService.Update(
		c.Request.Context(),
		id,
		input.Name,
		input.FileSystem,
		input.ProtocolConnection,
	)

	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteStorageProvider(c *gin.Context) {
	id := storage_provider.ID(c.Param("id"))

	err := h.storageProviderService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
