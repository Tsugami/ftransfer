package transfer

import (
	"context"

	"github.com/Tsugami/ftransfer/internal/storage_provider"
)

type TransferService struct {
	Repo TransferRepository
}

func NewService(repo TransferRepository) *TransferService {
	return &TransferService{
		Repo: repo,
	}
}
func (s *TransferService) Create(ctx context.Context, sourceDir Directory, destinationDir Directory, postTransferSourceDir Directory, sourceStorageProviderID storage_provider.ID, destinationStorageProviderID storage_provider.ID) (*Transfer, error) {
	transfer := Transfer{
		SourceDir:                    sourceDir,
		DestinationDir:               destinationDir,
		PostTransferSourceDir:        postTransferSourceDir,
		SourceStorageProviderID:      sourceStorageProviderID,
		DestinationStorageProviderID: destinationStorageProviderID,
	}

	if err := transfer.Validate(); err != nil {
		return nil, err
	}

	err := s.Repo.Create(ctx, &transfer)
	if err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (s *TransferService) List(ctx context.Context) ([]*Transfer, error) {
	all, err := s.Repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return all, nil
}
func (s *TransferService) Get(ctx context.Context, id ID) (*Transfer, error) {
	transfer, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrGetTransfer
	}
	return transfer, nil
}

func (s *TransferService) Update(ctx context.Context, id ID, sourceDir Directory, destinationDir Directory, postTransferSourceDir Directory, sourceStorageProviderID storage_provider.ID, destinationStorageProviderID storage_provider.ID) error {
	transfer, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return ErrGetTransfer
	}

	transfer.SourceDir = sourceDir
	transfer.DestinationDir = destinationDir
	transfer.PostTransferSourceDir = postTransferSourceDir
	transfer.SourceStorageProviderID = sourceStorageProviderID
	transfer.DestinationStorageProviderID = destinationStorageProviderID

	if err := transfer.Validate(); err != nil {
		return err
	}

	err = s.Repo.Update(ctx, transfer)
	if err != nil {
		return err
	}

	return nil
}
func (s *TransferService) Delete(ctx context.Context, id ID) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return ErrDeleteTransfer
	}

	return nil
}
