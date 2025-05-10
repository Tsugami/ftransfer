package storage_provider_client

import (
	"fmt"

	"github.com/Tsugami/ftransfer/internal/protocol"
)

type StorageProviderClientFactory struct {
}

func NewStorageProviderClientFactory() *StorageProviderClientFactory {
	return &StorageProviderClientFactory{}
}

func (f *StorageProviderClientFactory) CreateClient(connection protocol.Connection) (StorageProviderClient, error) {
	switch connection.GetProtocol() {
	case protocol.ProtocolSFTP:
		sftpConn, ok := connection.(*protocol.SFTPConnection)
		if !ok {
			return nil, fmt.Errorf("invalid connection tyc22pe for SFTP protocol")
		}

		if err := sftpConn.Validate(); err != nil {
			return nil, fmt.Errorf("invalid SFTP connection: %w", err)
		}

		return NewStorageProviderClientSFTP(&StorageProviderClientConnection{
			Host: sftpConn.Host,
		})
	case protocol.ProtocolFTP:
		ftpConn, ok := connection.(*protocol.FTPConnection)
		if !ok {
			return nil, fmt.Errorf("invalid connection type for FTP protocol")
		}

		if err := ftpConn.Validate(); err != nil {
			return nil, fmt.Errorf("invalid FTP connection: %w", err)
		}

		return NewStorageProviderClientSFTP(&StorageProviderClientConnection{
			Host:     ftpConn.Host,
			Port:     ftpConn.Port,
			Username: ftpConn.Username,
			Password: ftpConn.Password,
		})
	default:
		return nil, fmt.Errorf("unsupported protocol: %s", connection.GetProtocol())
	}
}
