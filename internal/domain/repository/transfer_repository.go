package repository

import (
	"context"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

// TransferRepository defines the interface for transfer persistence operations
type TransferRepository interface {
	Create(ctx context.Context, transfer *model.Transfer) error
	GetByID(ctx context.Context, id string) (*model.Transfer, error)
	Update(ctx context.Context, transfer *model.Transfer) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Transfer, error)
	ListBySourceFolder(ctx context.Context, sourceFolderID string) ([]*model.Transfer, error)
	ListByDestinationFolder(ctx context.Context, destinationFolderID string) ([]*model.Transfer, error)
}
