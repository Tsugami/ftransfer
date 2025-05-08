package main

import (
	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.RouterGroup,
	// transferService transfer.TransferService,
	storageProviderService *storage_provider.StorageProviderService,
) {
	handler := &Handler{
		storageProviderService: storageProviderService,
	}

	sp_group := router.Group("/storage-providers")
	{
		sp_group.POST("", handler.CreateStorageProvider)
		sp_group.GET("", handler.ListStorageProviders)
		sp_group.GET("/:id", handler.GetStorageProvider)
		sp_group.PUT("/:id", handler.UpdateStorageProvider)
		sp_group.DELETE("/:id", handler.DeleteStorageProvider)
	}

	// rg.GET("/transfers", listTransfers)
	// rg.GET("/transfers/:id", getTransfer)
	// rg.PUT("/transfers/:id", updateTransfer)
}
