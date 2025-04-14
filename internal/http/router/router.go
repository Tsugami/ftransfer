package router

import (
	"github.com/Tsugami/ftransfer/internal/domain/repository"
	"github.com/Tsugami/ftransfer/internal/http/handler"
	"github.com/Tsugami/ftransfer/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	connectorHandler *handler.ConnectorHandler
	folderHandler    *handler.FolderHandler
	transferHandler  *handler.TransferHandler
}

func NewRouter(
	connectorRepo repository.ConnectorRepository,
	folderRepo repository.FolderRepository,
	transferRepo repository.TransferRepository,
) *Router {
	return &Router{
		connectorHandler: handler.NewConnectorHandler(connectorRepo),
		folderHandler:    handler.NewFolderHandler(folderRepo),
		transferHandler:  handler.NewTransferHandler(transferRepo),
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.ErrorHandling())

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Connector routes
		connectors := v1.Group("/connectors")
		{
			connectors.POST("", r.connectorHandler.Create)
			connectors.GET("", r.connectorHandler.List)
			connectors.GET("/:id", r.connectorHandler.Get)
			connectors.DELETE("/:id", r.connectorHandler.Delete)
		}

		// Folder routes
		folders := v1.Group("/folders")
		{
			folders.POST("", r.folderHandler.Create)
			folders.GET("/:id", r.folderHandler.Get)
			folders.DELETE("/:id", r.folderHandler.Delete)
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
