package response

import (
	"github.com/Tsugami/ftransfer/internal/domain/model"
	"github.com/Tsugami/ftransfer/internal/domain/model/protocol"
)

type StorageProviderResponse struct {
	ID                 string                 `json:"id"`
	Name               string                 `json:"name"`
	Description        string                 `json:"description"`
	Protocol           string                 `json:"protocol"`
	FileSystem         string                 `json:"file_system"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection"`
	Tags               TagsResponse           `json:"tags"`
	CreatedAt          string                 `json:"created_at"`
	UpdatedAt          string                 `json:"updated_at"`
}

func NewStorageProviderResponse(storageProvider *model.StorageProvider) *StorageProviderResponse {
	// Convert protocol.Connection to map[string]interface{}
	protocolConnMap := make(map[string]interface{})
	switch conn := storageProvider.ProtocolConnection.(type) {
	case *protocol.SFTPConnection:
		protocolConnMap["host"] = conn.Host
		protocolConnMap["port"] = conn.Port
		protocolConnMap["username"] = conn.Username
		// Omit sensitive fields: password, private_key, key_passphrase
	case *protocol.FTPConnection:
		protocolConnMap["host"] = conn.Host
		protocolConnMap["port"] = conn.Port
		protocolConnMap["username"] = conn.Username
		protocolConnMap["passive_mode"] = conn.PassiveMode
		// Omit sensitive field: password
	case *protocol.S3Connection:
		protocolConnMap["region"] = conn.Region
		protocolConnMap["bucket"] = conn.Bucket
		protocolConnMap["endpoint"] = conn.Endpoint
		protocolConnMap["use_ssl"] = conn.UseSSL
		// Omit sensitive fields: access_key_id, secret_access_key
	}

	// Convert model.Tags to TagsResponse
	tags := make(TagsResponse, len(storageProvider.Tags))
	for i, tag := range storageProvider.Tags {
		tags[i] = TagResponse{
			Name:  tag.Name,
			Value: tag.Value,
		}
	}

	return &StorageProviderResponse{
		ID:                 storageProvider.ID,
		Name:               storageProvider.Name,
		FileSystem:         string(storageProvider.FileSystem),
		ProtocolConnection: protocolConnMap,
		Tags:               tags,
		CreatedAt:          storageProvider.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:          storageProvider.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
