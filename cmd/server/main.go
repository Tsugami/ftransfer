package main

import (
	"log"

	"github.com/Tsugami/ftransfer/internal/http/router"
	"github.com/Tsugami/ftransfer/internal/infrastructure/memory"
)

func main() {
	// Initialize repositories
	connectorRepo := memory.NewConnectorRepository()
	folderRepo := memory.NewFolderRepository()
	transferRepo := memory.NewTransferRepository()

	// Initialize router
	r := router.NewRouter(connectorRepo, folderRepo, transferRepo)
	engine := r.SetupRoutes()

	// Start server
	log.Println("Server starting on :8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
