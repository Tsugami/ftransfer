package handler

import (
	"fmt"
	"net/http"

	"github.com/Tsugami/ftransfer/internal/domain/repository"
	"github.com/Tsugami/ftransfer/internal/http/dto/request"
	"github.com/Tsugami/ftransfer/internal/http/dto/response"
	"github.com/gin-gonic/gin"
)

type StorageProviderHandler struct {
	storageProviderRepo repository.StorageProviderRepository
}

func NewStorageProviderHandler(storageProviderRepo repository.StorageProviderRepository) *StorageProviderHandler {
	return &StorageProviderHandler{
		storageProviderRepo: storageProviderRepo,
	}
}

func (h *StorageProviderHandler) Create(c *gin.Context) {
	var createReq request.CreateStorageProviderRequest
	if err := c.ShouldBindJSON(&createReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storageProvider, err := createReq.ToDomain()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.storageProviderRepo.Create(c.Request.Context(), storageProvider); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response.NewStorageProviderResponse(storageProvider))
}

func (h *StorageProviderHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "StorageProvider ID is required"})
		return
	}

	storageProvider, err := h.storageProviderRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response.NewStorageProviderResponse(storageProvider))
}

func (h *StorageProviderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "StorageProvider ID is required"})
		return
	}

	if err := h.storageProviderRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *StorageProviderHandler) List(c *gin.Context) {
	storageProviders, err := h.storageProviderRepo.List(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	storageProviderResponses := make([]*response.StorageProviderResponse, len(storageProviders))
	for i, storageProvider := range storageProviders {
		storageProviderResponses[i] = response.NewStorageProviderResponse(storageProvider)
	}

	c.JSON(200, storageProviderResponses)
}
