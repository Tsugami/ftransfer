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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusCreated, storage)
}
