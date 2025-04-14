package repository

import (
	"context"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

// ConnectorRepository defines the interface for connector persistence operations
type ConnectorRepository interface {
	Create(ctx context.Context, connector *model.Connector) error
	GetByID(ctx context.Context, id string) (*model.Connector, error)
	GetByName(ctx context.Context, name string) (*model.Connector, error)
	Update(ctx context.Context, connector *model.Connector) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Connector, error)
}
