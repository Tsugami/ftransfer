package main

import (
	"net/http"

	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/gin-gonic/gin"
)

type CreateStorageProviderRequest struct {
	Name               string                 `json:"name" binding:"required"`
	FileSystem         string                 `json:"file_system" binding:"required"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection" binding:"required"`
}

type UpdateStorageProviderRequest struct {
	Name               string                 `json:"name" binding:"required"`
	FileSystem         string                 `json:"file_system" binding:"required"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection" binding:"required"`
}

type StorageProviderResponse struct {
	ID                 string                 `json:"id"`
	Name               string                 `json:"name"`
	FileSystem         string                 `json:"file_system"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection"`
}

func NewStorageProviderResponse(sp *storage_provider.StorageProvider) StorageProviderResponse {
	return StorageProviderResponse{
		ID:                 sp.ID.String(),
		Name:               sp.Name,
		FileSystem:         string(sp.FileSystem),
		ProtocolConnection: sp.ProtocolConnection.GetJson(),
	}
}

func NewStorageProviderResponseList(sps []*storage_provider.StorageProvider) []StorageProviderResponse {
	response := make([]StorageProviderResponse, len(sps))
	for i, sp := range sps {
		response[i] = NewStorageProviderResponse(sp)
	}
	return response
}

func (s *Handler) CreateStorageProvider(c *gin.Context) {
	var input CreateStorageProviderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sp, err := s.storageProviderService.Create(c.Request.Context(),
		input.Name,
		input.FileSystem,
		input.ProtocolConnection,
	)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, NewStorageProviderResponse(sp))
}

func (s *Handler) ListStorageProviders(c *gin.Context) {
	sps, err := s.storageProviderService.List(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, NewStorageProviderResponseList(sps))
}

func (s *Handler) GetStorageProvider(c *gin.Context) {
	id := c.Param("id")
	sp, err := s.storageProviderService.Get(c.Request.Context(), storage_provider.ID(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, NewStorageProviderResponse(sp))
}

func (s *Handler) UpdateStorageProvider(c *gin.Context) {
	id := c.Param("id")
	var input UpdateStorageProviderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.storageProviderService.Update(c.Request.Context(), storage_provider.ID(id),
		input.Name,
		input.FileSystem,
		input.ProtocolConnection,
	)

	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Handler) DeleteStorageProvider(c *gin.Context) {
	id := c.Param("id")
	err := s.storageProviderService.Delete(c.Request.Context(), storage_provider.ID(id))
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
