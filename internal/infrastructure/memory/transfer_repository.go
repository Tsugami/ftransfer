package memory

import (
	"context"
	"sync"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type TransferRepository struct {
	mu        sync.RWMutex
	transfers map[string]*model.Transfer
}

func NewTransferRepository() *TransferRepository {
	return &TransferRepository{
		transfers: make(map[string]*model.Transfer),
	}
}

func (r *TransferRepository) Create(ctx context.Context, transfer *model.Transfer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if transfer.ID == "" {
		transfer.ID = generateID()
	}

	r.transfers[transfer.ID] = transfer
	return nil
}

func (r *TransferRepository) GetByID(ctx context.Context, id string) (*model.Transfer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	transfer, exists := r.transfers[id]
	if !exists {
		return nil, model.ErrTransferNotFound
	}

	return transfer, nil
}

func (r *TransferRepository) List(ctx context.Context) ([]*model.Transfer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	transfers := make([]*model.Transfer, 0, len(r.transfers))
	for _, transfer := range r.transfers {
		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (r *TransferRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.transfers[id]; !exists {
		return model.ErrTransferNotFound
	}

	delete(r.transfers, id)
	return nil
}

func (r *TransferRepository) Update(ctx context.Context, transfer *model.Transfer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.transfers[transfer.ID]; !exists {
		return model.ErrTransferNotFound
	}

	r.transfers[transfer.ID] = transfer
	return nil
}

func (r *TransferRepository) ListBySourceFolder(ctx context.Context, sourceFolderID string) ([]*model.Transfer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var transfers []*model.Transfer
	for _, transfer := range r.transfers {
		if transfer.SourceFolderID == sourceFolderID {
			transfers = append(transfers, transfer)
		}
	}

	return transfers, nil
}

func (r *TransferRepository) ListByFolderID(ctx context.Context, folderID string) ([]*model.Transfer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var transfers []*model.Transfer
	for _, t := range r.transfers {
		if t.SourceFolderID == folderID || t.DestinationFolderID == folderID {
			transfers = append(transfers, t)
		}
	}
	return transfers, nil
}

func generateID() string {
	return "transfer-" + generateUUID()
}

func generateUUID() string {
	// Implementação simples de UUID para testes
	return "123e4567-e89b-12d3-a456-426614174000"
}
