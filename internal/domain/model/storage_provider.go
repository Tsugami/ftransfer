package model

import (
	"time"

	"github.com/Tsugami/ftransfer/internal/domain/errs"
	"github.com/Tsugami/ftransfer/internal/domain/model/protocol"
	"github.com/google/uuid"
)

// StorageProvider represents a connection to a file system
type StorageProvider struct {
	Base
	ID                 string
	Name               string
	ProtocolConnection protocol.Connection
	FileSystem         FileSystemType
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// NewStorageProvider creates a new StorageProvider instance
func NewStorageProvider(name string, protocolConnection protocol.Connection, fileSystem FileSystemType) *StorageProvider {
	return &StorageProvider{
		ID:                 uuid.New().String(),
		Name:               name,
		ProtocolConnection: protocolConnection,
		FileSystem:         fileSystem,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}

// Validate performs validation on the storage provider
func (c *StorageProvider) Validate() error {
	if c.Name == "" {
		return errs.ErrEmptyName
	}

	if c.ProtocolConnection == nil {
		return errs.ErrEmptyProtocolConnection
	}

	if err := c.ProtocolConnection.Validate(); err != nil {
		return err
	}

	if c.FileSystem == "" {
		return errs.ErrEmptyFileSystem
	}

	if !c.FileSystem.IsValid() {
		return errs.ErrInvalidFileSystem
	}

	return nil
}
