package main

import (
	"context"
	"log"
	"time"

	"github.com/Tsugami/ftransfer/internal/storage_provider_client"
)

type CronJob struct {
	storageProviderClientService *storage_provider_client.StorageProviderClientProviderService
}

func NewCronJob(storageProviderClientService *storage_provider_client.StorageProviderClientProviderService) *CronJob {
	return &CronJob{
		storageProviderClientService: storageProviderClientService,
	}
}

func (c *CronJob) Run() {
	log.Println("CronJob started")
	for {
		errorFiles, err := c.storageProviderClientService.TransferFiles(context.Background())
		if err != nil {
			log.Printf("failed to transfer files: %v", err)
		}

		log.Printf("transferred %d files with errors: %v", len(errorFiles), errorFiles)

		time.Sleep(1 * time.Second)
	}
}
