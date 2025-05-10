package storage_provider_client

import (
	"context"
	"fmt"
	"path"

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

	fmt.Printf("%s found %d files\n", transfer.SourceDir.String(), len(sourceFiles))
	errorFiles := []string{}
	for _, sourceFile := range sourceFiles {
		sourcePath := path.Join(transfer.SourceDir.String(), sourceFile)
		sourceFileReader, err := sourceStorageProviderClient.ReadFile(ctx, sourcePath)
		if err != nil {
			fmt.Printf("error reading file %s: %v\n", sourcePath, err)
			errorFiles = append(errorFiles, sourceFile)
			continue
		}

		destinationPath := path.Join(transfer.DestinationDir.String(), sourceFile)
		err = destinationStorageProviderClient.WriteFile(ctx, destinationPath, sourceFileReader)
		if err != nil {
			fmt.Printf("error writing file %s: %v\n", destinationPath, err)
			errorFiles = append(errorFiles, sourceFile)
			continue
		}

		// post
		err = sourceStorageProviderClient.MakeDir(ctx, transfer.PostTransferSourceDir.String())
		if err != nil {
			fmt.Printf("error making directory %s: %v\n", transfer.PostTransferSourceDir.String(), err)
			errorFiles = append(errorFiles, sourceFile)
			continue
		}

		postTransferSourcePath := path.Join(transfer.PostTransferSourceDir.String(), sourceFile)
		err = sourceStorageProviderClient.MoveFile(ctx, sourcePath, postTransferSourcePath)
		if err != nil {
			fmt.Printf("error moving file from %s to %s: %v\n", sourcePath, postTransferSourcePath, err)
			errorFiles = append(errorFiles, sourceFile)
			continue
		}
	}

	return errorFiles, nil
}
