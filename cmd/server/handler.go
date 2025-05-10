package main

import (
	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/Tsugami/ftransfer/internal/transfer"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	storageProviderService *storage_provider.StorageProviderService
	transferService        *transfer.TransferService
}

func NewHandler(
	storageProviderService *storage_provider.StorageProviderService,
	transferService *transfer.TransferService,
) *Handler {
	return &Handler{
		storageProviderService: storageProviderService,
		transferService:        transferService,
	}
}

func (handler *Handler) SetupRoutes(
	router *gin.RouterGroup,
	storageProviderService *storage_provider.StorageProviderService,
	transferService *transfer.TransferService,
) {
	sp_group := router.Group("/storage-providers")
	{
		sp_group.POST("", handler.CreateStorageProvider)
		sp_group.GET("", handler.ListStorageProviders)
		sp_group.GET("/:id", handler.GetStorageProvider)
		sp_group.PUT("/:id", handler.UpdateStorageProvider)
		sp_group.DELETE("/:id", handler.DeleteStorageProvider)
	}

	rg := router.Group("/transfers")
	{
		rg.GET("", handler.ListTransfers)
		rg.GET("/:id", handler.GetTransfer)
		rg.PUT("/:id", handler.UpdateTransfer)
		rg.DELETE("/:id", handler.DeleteTransfer)
	}
}
