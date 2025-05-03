package transfer

import "context"

type TransferReader interface {
	GetByID(ctx context.Context, id ID) (*Transfer, error)
	List(ctx context.Context) ([]*Transfer, error)
}

type TransferWriter interface {
	Create(ctx context.Context, transfer *Transfer) error
	Update(ctx context.Context, transfer *Transfer) error
	Delete(ctx context.Context, id ID) error
}

// TransferRepository defines the interface for transfer persistence operations
type TransferRepository interface {
	TransferReader
	TransferWriter
}
