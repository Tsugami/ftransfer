package repository

import (
	"context"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type FolderReader interface {
	GetByID(ctx context.Context, id string) (*model.Folder, error)
	List(ctx context.Context) ([]*model.Folder, error)
	ListByConnectorID(ctx context.Context, connectorID string) ([]*model.Folder, error)
}

type FolderWriter interface {
	Create(ctx context.Context, folder *model.Folder) error
	Update(ctx context.Context, folder *model.Folder) error
	Delete(ctx context.Context, id string) error
}

// FolderRepository defines the interface for folder persistence operations
type FolderRepository interface {
	FolderReader
	FolderWriter
}
