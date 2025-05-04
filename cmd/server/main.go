package main

import (
	"log"

	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/Tsugami/ftransfer/repositories"
)

func main() {
	log.Println("Starting server...")
	// Initialize repositories
	// storageProviderRepo := memory.NewStorageProviderRepository()
	// transferRepo := memory.NewTransferRepository()
	db, err := OpenDB()
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	// storageProviderRepo := repositories.NewStorageProviderRepository(db)

	// transferRepo  := repositories.NewTransferRepository(db)
	var storageProviderRepo storage_provider.StorageProviderRepository = repositories.NewStorageProviderRepository(db)

	// Initialize services
	storageProviderService := storage_provider.NewService(storageProviderRepo)
	// // Initialize router
	engine := SetupRoutes(storageProviderService)
	engine.Use(SetupMiddleware())

	// // Start server
	log.Println("Server starting on :8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
