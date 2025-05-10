package storage_provider_client

import (
	"context"
	"fmt"
	"path"

	"github.com/Tsugami/ftransfer/internal/events"
	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/Tsugami/ftransfer/internal/transfer"
)

type StorageProviderClientProviderService struct {
	storageProviderRepository    storage_provider.StorageProviderRepository
	transferRepository           transfer.TransferRepository
	storageProviderClientFactory StorageProviderClientFactory
	eventRepository              events.EventRepository
}

func NewStorageProviderClientProviderService(
	storageProviderRepository storage_provider.StorageProviderRepository,
	transferRepository transfer.TransferRepository,
	eventRepository events.EventRepository,
) *StorageProviderClientProviderService {
	storageProviderClientFactory := NewStorageProviderClientFactory()
	return &StorageProviderClientProviderService{storageProviderRepository, transferRepository, *storageProviderClientFactory, eventRepository}
}

type ErrorList = []string

func (s *StorageProviderClientProviderService) TransferFiles(ctx context.Context) (ErrorList, error) {
	transfers, err := s.transferRepository.List(ctx)

	if err != nil {
		return nil, err
	}

	errorFiles := []string{}
	for _, transfer := range transfers {
		fmt.Printf("transferring files for transfer %s\n", transfer.ID)
		transferErrorFiles, err := s.transferFiles(ctx, transfer.ID)
		if err != nil {
			fmt.Printf("error transferring files for transfer %s: %v\n", transfer.ID, err)
			errorFiles = append(errorFiles, fmt.Sprintf("transfer %s: %v", transfer.ID, err))
			continue
		}
		errorFiles = append(errorFiles, transferErrorFiles...)
	}

	return errorFiles, nil
}

func (s *StorageProviderClientProviderService) transferFiles(ctx context.Context, transferId transfer.ID) (ErrorList, error) {
	logger := events.NewEventLogger(transferId.String(), s.eventRepository)

	transfer, err := s.transferRepository.GetByID(ctx, transferId)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("error getting transfer %s: %v", transferId.String(), err))
		return nil, err
	}

	logger.Info(ctx, fmt.Sprintf("transferring files from %s to %s", transfer.SourceDir.String(), transfer.DestinationDir.String()))

	sourceStorageProvider, err := s.storageProviderRepository.GetByID(ctx, transfer.SourceStorageProviderID)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("error getting source storage provider %s: %v", transfer.SourceStorageProviderID.String(), err))
		return nil, err
	}

	logger.Info(ctx, fmt.Sprintf("source storage provider: %s", sourceStorageProvider.ProtocolConnection))

	destinationStorageProvider, err := s.storageProviderRepository.GetByID(ctx, transfer.DestinationStorageProviderID)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("error getting destination storage provider %s: %v", transfer.DestinationStorageProviderID.String(), err))
		return nil, err
	}

	logger.Info(ctx, fmt.Sprintf("destination storage provider: %s", destinationStorageProvider.ProtocolConnection))

	sourceStorageProviderClient, err := s.storageProviderClientFactory.CreateClient(sourceStorageProvider.ProtocolConnection)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("error creating source storage provider client %s: %v", sourceStorageProvider.ProtocolConnection, err))
		return nil, err
	}

	logger.Info(ctx, fmt.Sprintf("source storage provider client: %s", sourceStorageProviderClient))

	destinationStorageProviderClient, err := s.storageProviderClientFactory.CreateClient(destinationStorageProvider.ProtocolConnection)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("error creating destination storage provider client %s: %v", destinationStorageProvider.ProtocolConnection, err))
		return nil, err
	}

	logger.Info(ctx, fmt.Sprintf("destination storage provider client: %s", destinationStorageProviderClient))

	sourceFiles, err := sourceStorageProviderClient.ListFiles(ctx, transfer.SourceDir.String())
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("error listing files in %s: %v", transfer.SourceDir.String(), err))
		return nil, err
	}

	sourceFilesCount := len(sourceFiles)
	if sourceFilesCount <= 0 {
		logger.Info(ctx, fmt.Sprintf("No files found in %s", transfer.SourceDir.String()))
		return nil, nil
	}

	logger.Info(ctx, fmt.Sprintf("%s found %d files", transfer.SourceDir.String(), len(sourceFiles)))
	errorFiles := []string{}
	for _, sourceFile := range sourceFiles {
		sourcePath := path.Join(transfer.SourceDir.String(), sourceFile)
		sourceFileReader, err := sourceStorageProviderClient.ReadFile(ctx, sourcePath)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("error reading file %s: %v", sourcePath, err))
			errorFiles = append(errorFiles, sourceFile)
			continue
		}

		destinationPath := path.Join(transfer.DestinationDir.String(), sourceFile)
		err = destinationStorageProviderClient.WriteFile(ctx, destinationPath, sourceFileReader)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("error writing file %s: %v", destinationPath, err))
			errorFiles = append(errorFiles, sourceFile)
			continue
		}

		// post
		err = sourceStorageProviderClient.MakeDir(ctx, transfer.PostTransferSourceDir.String())
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("error making directory %s: %v", transfer.PostTransferSourceDir.String(), err))
			errorFiles = append(errorFiles, sourceFile)
			continue
		}

		postTransferSourcePath := path.Join(transfer.PostTransferSourceDir.String(), sourceFile)

		logger.Info(ctx, fmt.Sprintf("moving file from %s to %s", sourcePath, postTransferSourcePath))
		err = sourceStorageProviderClient.MoveFile(ctx, sourcePath, postTransferSourcePath)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("error moving file from %s to %s: %v", sourcePath, postTransferSourcePath, err))
			errorFiles = append(errorFiles, sourceFile)
			continue
		}
	}

	err = logger.Flush(ctx)
	if err != nil {
		return nil, err
	}

	return errorFiles, nil
}
