package request

import (
	"github.com/Tsugami/ftransfer/internal/domain/model"
	"github.com/Tsugami/ftransfer/internal/domain/model/protocol"
)

type CreateConnectorRequest struct {
	Name               string                 `json:"name" validate:"required"`
	Description        string                 `json:"description"`
	Protocol           string                 `json:"protocol" validate:"required"`
	FileSystem         string                 `json:"file_system" validate:"required"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection" validate:"required"`
	Tags               []TagRequest           `json:"tags"`
}

type UpdateConnectorRequest struct {
	Name               string                 `json:"name"`
	Description        string                 `json:"description"`
	Protocol           string                 `json:"protocol"`
	FileSystem         string                 `json:"file_system"`
	ProtocolConnection map[string]interface{} `json:"protocol_connection"`
	Tags               []TagRequest           `json:"tags"`
}

func (r *CreateConnectorRequest) ToDomain() (*model.Connector, error) {
	// Create protocol connection based on protocol type
	var protocolConn protocol.Connection
	switch model.Protocol(r.Protocol) {
	case model.ProtocolSFTP:
		protocolConn = &protocol.SFTPConnection{
			Host:          r.ProtocolConnection["host"].(string),
			Port:          int(r.ProtocolConnection["port"].(float64)),
			Username:      r.ProtocolConnection["username"].(string),
			Password:      r.ProtocolConnection["password"].(string),
			PrivateKey:    r.ProtocolConnection["private_key"].(string),
			KeyPassphrase: r.ProtocolConnection["key_passphrase"].(string),
		}
	case model.ProtocolFTP:
		protocolConn = &protocol.FTPConnection{
			Host:        r.ProtocolConnection["host"].(string),
			Port:        int(r.ProtocolConnection["port"].(float64)),
			Username:    r.ProtocolConnection["username"].(string),
			Password:    r.ProtocolConnection["password"].(string),
			PassiveMode: r.ProtocolConnection["passive_mode"].(bool),
		}
	case model.ProtocolS3:
		protocolConn = &protocol.S3Connection{
			Region:          r.ProtocolConnection["region"].(string),
			Bucket:          r.ProtocolConnection["bucket"].(string),
			AccessKeyID:     r.ProtocolConnection["access_key_id"].(string),
			SecretAccessKey: r.ProtocolConnection["secret_access_key"].(string),
			Endpoint:        r.ProtocolConnection["endpoint"].(string),
			UseSSL:          r.ProtocolConnection["use_ssl"].(bool),
		}
	}

	connector := &model.Connector{
		Base:               model.NewBase(),
		Name:               r.Name,
		Description:        r.Description,
		Protocol:           model.Protocol(r.Protocol),
		FileSystem:         model.FileSystemType(r.FileSystem),
		ProtocolConnection: protocolConn,
	}

	for _, tag := range r.Tags {
		connector.AddTag(tag.Name, tag.Value)
	}

	if err := connector.Validate(); err != nil {
		return nil, err
	}

	return connector, nil
}

func (r *UpdateConnectorRequest) ToDomain(id string) (*model.Connector, error) {
	connector, err := (&CreateConnectorRequest{
		Name:               r.Name,
		Description:        r.Description,
		Protocol:           r.Protocol,
		FileSystem:         r.FileSystem,
		ProtocolConnection: r.ProtocolConnection,
		Tags:               r.Tags,
	}).ToDomain()
	if err != nil {
		return nil, err
	}
	connector.ID = id
	return connector, nil
}
