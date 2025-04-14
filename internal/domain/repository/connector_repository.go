package repository

import (
	"context"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

// ConnectorReader defines the interface for connector read operations
type ConnectorReader interface {
	GetByID(ctx context.Context, id string) (*model.Connector, error)
	List(ctx context.Context) ([]*model.Connector, error)
}

// ConnectorWriter defines the interface for connector write operations
type ConnectorWriter interface {
	Create(ctx context.Context, connector *model.Connector) error
	Update(ctx context.Context, connector *model.Connector) error
	Delete(ctx context.Context, id string) error
}

// ConnectorRepository combines reader and writer interfaces
type ConnectorRepository interface {
	ConnectorReader
	ConnectorWriter
}
