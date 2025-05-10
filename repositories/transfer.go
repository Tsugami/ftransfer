package repositories

import (
	"context"
	"database/sql"

	"github.com/Tsugami/ftransfer/internal/transfer"
)

type TransferRepository struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) *TransferRepository {
	return &TransferRepository{db: db}
}

func (r *TransferRepository) Create(ctx context.Context, transfer *transfer.Transfer) error {
	query := `INSERT INTO transfers (source_dir, destination_dir, post_transfer_source_dir, source_storage_provider_id, destination_storage_provider_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, transfer.SourceDir, transfer.DestinationDir, transfer.PostTransferSourceDir, transfer.SourceStorageProviderID.String(), transfer.DestinationStorageProviderID.String())
	return err
}

func (r *TransferRepository) List(ctx context.Context) ([]*transfer.Transfer, error) {
	query := `SELECT id, source_dir, destination_dir, post_transfer_source_dir, source_storage_provider_id, destination_storage_provider_id FROM transfers`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transfers := []*transfer.Transfer{}
	for rows.Next() {
		var transfer transfer.Transfer
		if err := rows.Scan(&transfer.ID, &transfer.SourceDir, &transfer.DestinationDir, &transfer.PostTransferSourceDir, &transfer.SourceStorageProviderID, &transfer.DestinationStorageProviderID); err != nil {
			return nil, err
		}
		transfers = append(transfers, &transfer)
	}

	return transfers, nil
}

func (r *TransferRepository) GetByID(ctx context.Context, id transfer.ID) (*transfer.Transfer, error) {
	query := `SELECT id, source_dir, destination_dir, post_transfer_source_dir, source_storage_provider_id, destination_storage_provider_id FROM transfers WHERE id = $1`
	var transfer transfer.Transfer
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&transfer.ID, &transfer.SourceDir, &transfer.DestinationDir, &transfer.PostTransferSourceDir, &transfer.SourceStorageProviderID, &transfer.DestinationStorageProviderID); err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (r *TransferRepository) Update(ctx context.Context, transfer *transfer.Transfer) error {
	query := `UPDATE transfers SET source_dir = $1, destination_dir = $2, post_transfer_source_dir = $3, source_storage_provider_id = $4, destination_storage_provider_id = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query, transfer.SourceDir, transfer.DestinationDir, transfer.PostTransferSourceDir, transfer.SourceStorageProviderID, transfer.DestinationStorageProviderID, transfer.ID)
	return err
}

func (r *TransferRepository) Delete(ctx context.Context, id transfer.ID) error {
	query := `DELETE FROM transfers WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
