package repository

import (
	"context"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type TransferReader interface {
	GetByID(ctx context.Context, id string) (*model.Transfer, error)
	List(ctx context.Context) ([]*model.Transfer, error)
	ListByFolderID(ctx context.Context, folderID string) ([]*model.Transfer, error)
}

type TransferWriter interface {
	Create(ctx context.Context, transfer *model.Transfer) error
	Update(ctx context.Context, transfer *model.Transfer) error
	Delete(ctx context.Context, id string) error
}

// TransferRepository defines the interface for transfer persistence operations
type TransferRepository interface {
	TransferReader
	TransferWriter
}
