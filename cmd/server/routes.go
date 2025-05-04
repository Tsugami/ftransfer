package main

import (
	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	// transferService transfer.TransferService,
	storageProviderService *storage_provider.StorageProviderService,
) *gin.Engine {

	handler := &Handler{
		storageProviderService: storageProviderService,
	}

	// Initialize the router
	router := gin.Default()
	router.Use(SetupMiddleware())

	v1 := router.Group("/api/v1")
	// Define the routes

	sp_group := v1.Group("/storage-providers")
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

	return router
}
