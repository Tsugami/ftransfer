package request

import (
	"github.com/Tsugami/ftransfer/internal/domain/model"
	"github.com/Tsugami/ftransfer/internal/domain/model/protocol"
)

type CreateStorageProviderRequest struct {
	Name               string                 `json:"name" validate:"required"`
	Description        string                 `json:"description"`
	FileSystem         string                 `json:"file_system" validate:"required"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection" validate:"required"`
	Tags               []TagRequest           `json:"tags"`
}

type UpdateStorageProviderRequest struct {
	Name               string                 `json:"name"`
	Description        string                 `json:"description"`
	FileSystem         string                 `json:"file_system"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection"`
	Tags               []TagRequest           `json:"tags"`
}

func (r *CreateStorageProviderRequest) ToDomain() (*model.StorageProvider, error) {
	// Create protocol connection based on protocol type
	var protocolConn protocol.Connection

	protocolInput, err := protocol.NewProtocol(r.ProtocolConnection["protocol"].(string))
	if err != nil {
		return nil, err
	}

	switch protocolInput {
	case protocol.ProtocolSFTP:
		protocolConn = &protocol.SFTPConnection{
			Host:          r.ProtocolConnection["host"].(string),
			Port:          int(r.ProtocolConnection["port"].(float64)),
			Username:      r.ProtocolConnection["username"].(string),
			Password:      r.ProtocolConnection["password"].(string),
			PrivateKey:    r.ProtocolConnection["private_key"].(string),
			KeyPassphrase: r.ProtocolConnection["key_passphrase"].(string),
		}
	case protocol.ProtocolFTP:
		protocolConn = &protocol.FTPConnection{
			Host:        r.ProtocolConnection["host"].(string),
			Port:        int(r.ProtocolConnection["port"].(float64)),
			Username:    r.ProtocolConnection["username"].(string),
			Password:    r.ProtocolConnection["password"].(string),
			PassiveMode: r.ProtocolConnection["passive_mode"].(bool),
		}
	case protocol.ProtocolS3:
		protocolConn = &protocol.S3Connection{
			Region:          r.ProtocolConnection["region"].(string),
			Bucket:          r.ProtocolConnection["bucket"].(string),
			AccessKeyID:     r.ProtocolConnection["access_key_id"].(string),
			SecretAccessKey: r.ProtocolConnection["secret_access_key"].(string),
			Endpoint:        r.ProtocolConnection["endpoint"].(string),
			UseSSL:          r.ProtocolConnection["use_ssl"].(bool),
		}
	}

	storageProvider := model.NewStorageProvider(
		r.Name,
		protocolConn,
		model.FileSystemType(r.FileSystem),
	)

	for _, tag := range r.Tags {
		storageProvider.AddTag(tag.Name, tag.Value)
	}

	if err := storageProvider.Validate(); err != nil {
		return nil, err
	}

	return storageProvider, nil
}

func (r *UpdateStorageProviderRequest) ToDomain(id string) (*model.StorageProvider, error) {
	storageProvider, err := (&CreateStorageProviderRequest{
		Name:               r.Name,
		Description:        r.Description,
		FileSystem:         r.FileSystem,
		ProtocolConnection: r.ProtocolConnection,
		Tags:               r.Tags,
	}).ToDomain()
	if err != nil {
		return nil, err
	}
	storageProvider.ID = id
	return storageProvider, nil
}
