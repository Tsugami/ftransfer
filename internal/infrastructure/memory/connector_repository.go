package memory

import (
	"context"
	"sync"

	"github.com/Tsugami/ftransfer/internal/domain/model"
)

type ConnectorRepository struct {
	mu         sync.RWMutex
	connectors map[string]*model.Connector
}

func NewConnectorRepository() *ConnectorRepository {
	return &ConnectorRepository{
		connectors: make(map[string]*model.Connector),
	}
}

func (r *ConnectorRepository) Create(ctx context.Context, connector *model.Connector) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if connector.ID == "" {
		connector.ID = generateID()
	}

	r.connectors[connector.ID] = connector
	return nil
}

func (r *ConnectorRepository) GetByID(ctx context.Context, id string) (*model.Connector, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	connector, exists := r.connectors[id]
	if !exists {
		return nil, model.ErrConnectorNotFound
	}

	return connector, nil
}

func (r *ConnectorRepository) List(ctx context.Context) ([]*model.Connector, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	connectors := make([]*model.Connector, 0, len(r.connectors))
	for _, connector := range r.connectors {
		connectors = append(connectors, connector)
	}

	return connectors, nil
}

func (r *ConnectorRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.connectors[id]; !exists {
		return model.ErrConnectorNotFound
	}

	delete(r.connectors, id)
	return nil
}

func (r *ConnectorRepository) Update(ctx context.Context, connector *model.Connector) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.connectors[connector.ID]; !exists {
		return model.ErrConnectorNotFound
	}

	r.connectors[connector.ID] = connector
	return nil
}
