package storage_provider_client

import (
	"context"
	"fmt"

	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/Tsugami/ftransfer/internal/transfer"
)

type StorageProviderClientProviderService struct {
	storageProviderRepository    storage_provider.StorageProviderRepository
	transferRepository           transfer.TransferRepository
	storageProviderClientFactory StorageProviderClientFactory
}

func NewStorageProviderClientProviderService(
	storageProviderRepository storage_provider.StorageProviderRepository,
	transferRepository transfer.TransferRepository,
) *StorageProviderClientProviderService {
	storageProviderClientFactory := NewStorageProviderClientFactory()
	return &StorageProviderClientProviderService{storageProviderRepository, transferRepository, *storageProviderClientFactory}
}

type ErrorList = []string

func (s *StorageProviderClientProviderService) TransferFiles(ctx context.Context) (ErrorList, error) {
	transfers, err := s.transferRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	errorFiles := []string{}
	for _, transfer := range transfers {
		transferErrorFiles, err := s.transferFiles(ctx, transfer.ID)
		if err != nil {
			errorFiles = append(errorFiles, fmt.Sprintf("transfer %s: %v", transfer.ID, err))
			continue
		}
		errorFiles = append(errorFiles, transferErrorFiles...)
	}

	return errorFiles, nil
}

func (s *StorageProviderClientProviderService) transferFiles(ctx context.Context, transferId transfer.ID) (ErrorList, error) {

	transfer, err := s.transferRepository.GetByID(ctx, transferId)
	if err != nil {
		return nil, err
	}

	sourceStorageProvider, err := s.storageProviderRepository.GetByID(ctx, transfer.SourceStorageProviderID)
	if err != nil {
		return nil, err
	}

	destinationStorageProvider, err := s.storageProviderRepository.GetByID(ctx, transfer.DestinationStorageProviderID)
	if err != nil {
		return nil, err
	}

	sourceStorageProviderClient, err := s.storageProviderClientFactory.CreateClient(sourceStorageProvider.ProtocolConnection)
	if err != nil {
		return nil, err
	}

	destinationStorageProviderClient, err := s.storageProviderClientFactory.CreateClient(destinationStorageProvider.ProtocolConnection)
	if err != nil {
		return nil, err
	}

	sourceFiles, err := sourceStorageProviderClient.ListFiles(ctx, transfer.SourceDir.String())
	if err != nil {
		return nil, err
	}

	errorFiles := []string{}
	for _, sourceFile := range sourceFiles {
		sourceFileReader, err := sourceStorageProviderClient.ReadFile(ctx, sourceFile)
		if err != nil {
			errorFiles = append(errorFiles, sourceFile)
			continue
		}

		err = destinationStorageProviderClient.WriteFile(ctx, transfer.DestinationDir.String(), sourceFileReader)
		if err != nil {
			errorFiles = append(errorFiles, sourceFile)
			continue
		}
	}

	return errorFiles, nil
}
