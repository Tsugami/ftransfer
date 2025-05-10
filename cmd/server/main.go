package main

import (
	"log"
	"os"

	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/Tsugami/ftransfer/internal/transfer"
	"github.com/Tsugami/ftransfer/repositories"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")
	// Initialize repositories
	db, err := OpenDB()
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	defer db.Close()

	conn := db.GetDB()

	var storageProviderRepo storage_provider.StorageProviderRepository = repositories.NewStorageProviderRepository(conn)
	var transferRepo transfer.TransferRepository = repositories.NewTransferRepository(conn)
	// Initialize services
	storageProviderService := storage_provider.NewService(storageProviderRepo)
	transferService := transfer.NewService(transferRepo)

	handler := NewHandler(storageProviderService, transferService)
	server := gin.Default()

	// Initialize middleware
	SetupMiddleware(server)

	PUBLIC_DIR := os.Getenv("PUBLIC_DIR")
	if PUBLIC_DIR == "" {
		PUBLIC_DIR = "./public"
	}

	log.Println("Serving files from", PUBLIC_DIR)
	server.Use(static.Serve("/", static.LocalFile(PUBLIC_DIR, false)))

	api := server.Group("/api/v1")
	handler.SetupRoutes(api, storageProviderService, transferService)

	// Start server
	log.Println("Server starting on :8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
