package storage_provider

import (
	"context"
)

// StorageProviderReader defines the interface for storage provider read operations
type StorageProviderReader interface {
	GetByID(ctx context.Context, id ID) (*StorageProvider, error)
	List(ctx context.Context) ([]*StorageProvider, error)
}

// StorageProviderWriter defines the interface for storage provider write operations
type StorageProviderWriter interface {
	Create(ctx context.Context, storageProvider *StorageProvider) (ID, error)
	Update(ctx context.Context, storageProvider *StorageProvider) error
	Delete(ctx context.Context, id ID) error
}

// StorageProviderRepository combines reader and writer interfaces
type StorageProviderRepository interface {
	StorageProviderReader
	StorageProviderWriter
}
