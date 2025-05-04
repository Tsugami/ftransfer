package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/Tsugami/ftransfer/internal/protocol"
	"github.com/Tsugami/ftransfer/internal/storage_provider"
	"github.com/google/uuid"
)

type StorageProviderRepository struct {
	db *sql.DB
}

type ProtocolConnectionJsonb struct {
	connection map[string]interface{}
}

func (p *ProtocolConnectionJsonb) Scan(value interface{}) error {
	b, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p.connection)
}

func (p *ProtocolConnectionJsonb) To() (protocol.Connection, error) {
	protocol, err := protocol.NewConnection(p.connection)
	return protocol, err

}

func NewStorageProviderRepository(db *sql.DB) *StorageProviderRepository {
	return &StorageProviderRepository{db: db}
}

func (r *StorageProviderRepository) Create(ctx context.Context, storageProvider *storage_provider.StorageProvider) (*storage_provider.ID, error) {
	if err := storageProvider.Validate(); err != nil {
		return nil, err
	}

	id := storage_provider.NewID(uuid.New().String())

	query := `INSERT INTO storage_providers (id, name, file_system, protocol_connection) VALUES ($1, $2, $3, $4)`

	protocolConnection, err := json.Marshal(storageProvider.ProtocolConnection.GetJson())
	if err != nil {
		return nil, err
	}

	_, err = r.db.ExecContext(ctx,
		query,
		id,
		storageProvider.Name,
		storageProvider.FileSystem,
		protocolConnection,
	)

	if err != nil {

		return nil, err
	}

	return &id, nil
}

func (r *StorageProviderRepository) List(ctx context.Context) ([]*storage_provider.StorageProvider, error) {
	query := `SELECT * FROM storage_providers`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	storageProviders := []*storage_provider.StorageProvider{}
	for rows.Next() {
		var storageProvider = storage_provider.StorageProvider{}
		var protocolConnectionJsonb ProtocolConnectionJsonb

		if err := rows.Scan(&storageProvider.ID, &storageProvider.Name, &storageProvider.FileSystem, &protocolConnectionJsonb); err != nil {
			return nil, err
		}

		protocolConnection, err := protocolConnectionJsonb.To()
		if err != nil {
			return nil, err
		}

		storageProvider.ProtocolConnection = protocolConnection

		storageProviders = append(storageProviders, &storageProvider)

	}

	return storageProviders, nil
}

func (r *StorageProviderRepository) GetByID(ctx context.Context, id storage_provider.ID) (*storage_provider.StorageProvider, error) {
	query := `SELECT * FROM storage_providers WHERE id = $1`
	var storageProvider storage_provider.StorageProvider
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&storageProvider.ID, &storageProvider.Name, &storageProvider.FileSystem, &storageProvider.ProtocolConnection); err != nil {

		return nil, err
	}

	return &storageProvider, nil
}

func (r *StorageProviderRepository) Update(ctx context.Context, storageProvider *storage_provider.StorageProvider) error {
	query := `UPDATE storage_providers SET name = $1, file_system = $2, protocol_connection = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, storageProvider.Name, storageProvider.FileSystem, storageProvider.ProtocolConnection, storageProvider.ID)
	return err
}

func (r *StorageProviderRepository) Delete(ctx context.Context, id storage_provider.ID) error {
	query := `DELETE FROM storage_providers WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
