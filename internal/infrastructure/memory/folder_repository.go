package memory

import (
	"context"
	"sync"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type FolderRepository struct {
	mu      sync.RWMutex
	folders map[string]*model.Folder
}

func NewFolderRepository() *FolderRepository {
	return &FolderRepository{
		folders: make(map[string]*model.Folder),
	}
}

func (r *FolderRepository) Create(ctx context.Context, folder *model.Folder) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if folder.ID == "" {
		folder.ID = generateID()
	}

	r.folders[folder.ID] = folder
	return nil
}

func (r *FolderRepository) GetByID(ctx context.Context, id string) (*model.Folder, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	folder, exists := r.folders[id]
	if !exists {
		return nil, model.ErrFolderNotFound
	}

	return folder, nil
}

func (r *FolderRepository) List(ctx context.Context) ([]*model.Folder, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	folders := make([]*model.Folder, 0, len(r.folders))
	for _, folder := range r.folders {
		folders = append(folders, folder)
	}

	return folders, nil
}

func (r *FolderRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.folders[id]; !exists {
		return model.ErrFolderNotFound
	}

	delete(r.folders, id)
	return nil
}

func (r *FolderRepository) Update(ctx context.Context, folder *model.Folder) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.folders[folder.ID]; !exists {
		return model.ErrFolderNotFound
	}

	r.folders[folder.ID] = folder
	return nil
}

func (r *FolderRepository) ListByConnectorID(ctx context.Context, connectorID string) ([]*model.Folder, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var folders []*model.Folder
	for _, folder := range r.folders {
		if folder.ConnectorID == connectorID {
			folders = append(folders, folder)
		}
	}

	return folders, nil
}
