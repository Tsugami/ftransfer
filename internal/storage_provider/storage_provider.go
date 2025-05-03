package storage_provider

import (
	"github.com/Tsugami/ftransfer/internal/protocol"
)

// StorageProvider represents a connection to a file system
type StorageProvider struct {
	ID                 ID
	Name               string
	ProtocolConnection protocol.Connection
	FileSystem         FileSystemType
}

// Validate performs validation on the storage provider
func (c *StorageProvider) Validate() error {
	if c.Name == "" {
		return ErrEmptyName
	}

	if c.ProtocolConnection == nil {
		return ErrEmptyProtocolConnection
	}

	if err := c.ProtocolConnection.Validate(); err != nil {
		return err
	}

	if c.FileSystem == "" {
		return ErrEmptyFileSystem
	}

	if !c.FileSystem.IsValid() {
		return ErrInvalidFileSystem
	}

	return nil
}
