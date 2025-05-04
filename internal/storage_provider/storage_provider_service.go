package storage_provider

import (
	"context"

	"github.com/Tsugami/ftransfer/internal/protocol"
)

type StorageProviderService struct {
	Repo StorageProviderRepository
}

func NewService(repo StorageProviderRepository) *StorageProviderService {
	return &StorageProviderService{
		Repo: repo,
	}
}
func (s *StorageProviderService) Create(ctx context.Context, name string, fileSystem string, protocolConnection map[string]interface{}) (*StorageProvider, error) {
	protocolConn, err := protocol.NewConnection(protocolConnection)
	if err != nil {
		return nil, err
	}

	if err := protocolConn.Validate(); err != nil {
		return nil, err
	}

	storageProvider := StorageProvider{
		Name:               name,
		FileSystem:         FileSystemType(fileSystem),
		ProtocolConnection: protocolConn,
	}

	if err := storageProvider.Validate(); err != nil {
		return nil, err
	}

	id, err := s.Repo.Create(ctx, &storageProvider)
	if err != nil {
		return nil, err
	}

	storageProvider.ID = *id

	return &storageProvider, nil
}

func (s *StorageProviderService) List(ctx context.Context) ([]*StorageProvider, error) {
	all, err := s.Repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return all, nil
}
func (s *StorageProviderService) Get(ctx context.Context, id ID) (*StorageProvider, error) {
	storageProvider, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return storageProvider, nil
}

func (s *StorageProviderService) Update(ctx context.Context, id ID, name string, fileSystem string, protocolConnection map[string]interface{}) error {
	protocolConn, err := protocol.NewConnection(protocolConnection)
	if err != nil {
		return err
	}

	if err := protocolConn.Validate(); err != nil {
		return err
	}

	storageProvider := StorageProvider{
		ID:                 id,
		Name:               name,
		FileSystem:         FileSystemType(fileSystem),
		ProtocolConnection: protocolConn,
	}
	err = s.Repo.Update(ctx, &storageProvider)
	if err != nil {
		return err
	}

	return nil
}
func (s *StorageProviderService) Delete(ctx context.Context, id ID) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
