package router

import (
	"github.com/Tsugami/ftransfer/internal/domain/repository"
	"github.com/Tsugami/ftransfer/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	connectorReader repository.ConnectorReader
	connectorWriter repository.ConnectorWriter
	folderReader    repository.FolderReader
	folderWriter    repository.FolderWriter
	transferReader  repository.TransferReader
	transferWriter  repository.TransferWriter
}

func NewRouter(
	connectorReader repository.ConnectorReader,
	connectorWriter repository.ConnectorWriter,
	folderReader repository.FolderReader,
	folderWriter repository.FolderWriter,
	transferReader repository.TransferReader,
	transferWriter repository.TransferWriter,
) *Router {
	return &Router{
		connectorReader: connectorReader,
		connectorWriter: connectorWriter,
		folderReader:    folderReader,
		folderWriter:    folderWriter,
		transferReader:  transferReader,
		transferWriter:  transferWriter,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Connector routes
	connectorHandler := handler.NewConnectorHandler(r.connectorReader, r.connectorWriter)
	router.POST("/api/v1/connectors", connectorHandler.Create)
	router.GET("/api/v1/connectors/:id", connectorHandler.Get)
	router.DELETE("/api/v1/connectors/:id", connectorHandler.Delete)

	// Folder routes
	folderHandler := handler.NewFolderHandler(r.folderReader, r.folderWriter)
	router.POST("/api/v1/folders", folderHandler.Create)
	router.GET("/api/v1/folders/:id", folderHandler.Get)
	router.DELETE("/api/v1/folders/:id", folderHandler.Delete)

	// Transfer routes
	transferHandler := handler.NewTransferHandler(r.transferReader, r.transferWriter)
	router.POST("/api/v1/transfers", transferHandler.Create)
	router.GET("/api/v1/transfers/:id", transferHandler.Get)
	router.DELETE("/api/v1/transfers/:id", transferHandler.Delete)

	return router
}
