package router

import (
	"github.com/Tsugami/ftransfer/internal/domain/repository"
	"github.com/Tsugami/ftransfer/internal/http/handler"
	"github.com/Tsugami/ftransfer/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine                 *gin.Engine
	storageProviderHandler *handler.StorageProviderHandler
	transferHandler        *handler.TransferHandler
}

func NewRouter(
	storageProviderRepo repository.StorageProviderRepository,
	transferRepo repository.TransferRepository,
) *Router {
	return &Router{
		engine:                 gin.Default(),
		storageProviderHandler: handler.NewStorageProviderHandler(storageProviderRepo),
		transferHandler:        handler.NewTransferHandler(transferRepo),
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.ErrorHandling())

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Storage Provider routes
		storageProviders := v1.Group("/storage-providers")
		{
			storageProviders.POST("", r.storageProviderHandler.Create)
			storageProviders.GET("", r.storageProviderHandler.List)
			storageProviders.GET("/:id", r.storageProviderHandler.Get)
			storageProviders.DELETE("/:id", r.storageProviderHandler.Delete)
		}

		// Transfer routes
		transfers := v1.Group("/transfers")
		{
			transfers.POST("", r.transferHandler.Create)
			transfers.GET("/:id", r.transferHandler.Get)
			transfers.DELETE("/:id", r.transferHandler.Delete)
		}
	}

	return router
}
