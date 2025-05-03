package memory

import (
	"context"
	"sync"

	"github.com/Tsugami/ftransfer/internal/domain/errs"
	"github.com/Tsugami/ftransfer/internal/domain/model"
	"github.com/Tsugami/ftransfer/internal/domain/repository"
)

type storageProviderRepository struct {
	storageProviders map[string]*model.StorageProvider
	mu               sync.RWMutex
}

func NewStorageProviderRepository() repository.StorageProviderRepository {
	return &storageProviderRepository{
		storageProviders: make(map[string]*model.StorageProvider),
	}
}

func (r *storageProviderRepository) Create(ctx context.Context, storageProvider *model.StorageProvider) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storageProviders[storageProvider.ID] = storageProvider
	return nil
}

func (r *storageProviderRepository) GetByID(ctx context.Context, id string) (*model.StorageProvider, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	storageProvider, exists := r.storageProviders[id]
	if !exists {
		return nil, errs.ErrStorageProviderNotFound
	}

	return storageProvider, nil
}

func (r *storageProviderRepository) List(ctx context.Context) ([]*model.StorageProvider, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	storageProviders := make([]*model.StorageProvider, 0, len(r.storageProviders))
	for _, storageProvider := range r.storageProviders {
		storageProviders = append(storageProviders, storageProvider)
	}

	return storageProviders, nil
}

func (r *storageProviderRepository) Update(ctx context.Context, storageProvider *model.StorageProvider) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.storageProviders[storageProvider.ID]; !exists {
		return errs.ErrStorageProviderNotFound
	}

	r.storageProviders[storageProvider.ID] = storageProvider
	return nil
}

func (r *storageProviderRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.storageProviders[id]; !exists {
		return errs.ErrStorageProviderNotFound
	}

	delete(r.storageProviders, id)
	return nil
}
