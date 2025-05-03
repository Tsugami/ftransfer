package repository

import (
	"context"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

// StorageProviderReader defines the interface for storage provider read operations
type StorageProviderReader interface {
	GetByID(ctx context.Context, id string) (*model.StorageProvider, error)
	List(ctx context.Context) ([]*model.StorageProvider, error)
}

// StorageProviderWriter defines the interface for storage provider write operations
type StorageProviderWriter interface {
	Create(ctx context.Context, storageProvider *model.StorageProvider) error
	Update(ctx context.Context, storageProvider *model.StorageProvider) error
	Delete(ctx context.Context, id string) error
}

// StorageProviderRepository combines reader and writer interfaces
type StorageProviderRepository interface {
	StorageProviderReader
	StorageProviderWriter
}
