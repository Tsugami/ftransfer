package model

import (
	"github.com/Tsugami/ftransfer/internal/domain/model/protocol"
)

// Connector represents a connection to a file system
type Connector struct {
	Base
	Name               string
	Description        string
	Protocol           Protocol
	ProtocolConnection protocol.Connection
	FileSystem         FileSystemType
}

// NewConnector creates a new Connector instance
func NewConnector(name string, protocol Protocol, protocolConnection protocol.Connection, fileSystem FileSystemType) *Connector {
	return &Connector{
		Base:               NewBase(),
		Name:               name,
		Protocol:           protocol,
		ProtocolConnection: protocolConnection,
		FileSystem:         fileSystem,
	}
}

// Validate performs validation on the connector
func (c *Connector) Validate() error {
	if c.Name == "" {
		return ErrEmptyName
	}
	if c.Protocol == "" {
		return ErrEmptyProtocol
	}
	if !c.Protocol.IsValid() {
		return ErrInvalidProtocol
	}
	if c.ProtocolConnection == nil {
		return ErrEmptyProtocolConnection
	}

	// Validate that the ProtocolConnection matches the Protocol
	switch c.Protocol {
	case ProtocolSFTP:
		if _, ok := c.ProtocolConnection.(*protocol.SFTPConnection); !ok {
			return protocol.ErrInvalidSFTPConnection
		}
	case ProtocolFTP:
		if _, ok := c.ProtocolConnection.(*protocol.FTPConnection); !ok {
			return protocol.ErrInvalidFTPConnection
		}
	case ProtocolS3:
		if _, ok := c.ProtocolConnection.(*protocol.S3Connection); !ok {
			return protocol.ErrInvalidS3Connection
		}
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
