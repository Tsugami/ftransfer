package protocol

// Connection represents a generic protocol connection
type Connection interface {
	Validate() error
	GetProtocol() Protocol
}

func NewConnection(connection map[string]interface{}) (Connection, error) {
	protocol, err := NewProtocol(connection["protocol"].(string))
	if err != nil {
		return nil, err
	}

	var protocolConn Connection

	switch protocol {
	case ProtocolSFTP:
		protocolConn = &SFTPConnection{
			Host:          connection["host"].(string),
			Port:          int(connection["port"].(float64)),
			Username:      connection["username"].(string),
			Password:      connection["password"].(string),
			PrivateKey:    connection["private_key"].(string),
			KeyPassphrase: connection["key_passphrase"].(string),
		}
	case ProtocolFTP:
		protocolConn = &FTPConnection{
			Host:        connection["host"].(string),
			Port:        int(connection["port"].(float64)),
			Username:    connection["username"].(string),
			Password:    connection["password"].(string),
			PassiveMode: connection["passive_mode"].(bool),
		}
	case ProtocolS3:
		protocolConn = &S3Connection{
			Region:          connection["region"].(string),
			Bucket:          connection["bucket"].(string),
			AccessKeyID:     connection["access_key_id"].(string),
			SecretAccessKey: connection["secret_access_key"].(string),
			Endpoint:        connection["endpoint"].(string),
			UseSSL:          connection["use_ssl"].(bool),
		}
	}

	if protocolConn == nil {
		return nil, ErrInvalidProtocol
	}

	return protocolConn, nil
}
