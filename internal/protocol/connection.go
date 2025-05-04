package protocol

import (
	"fmt"

	im_util "github.com/Tsugami/ftransfer/pkg/interfacemaputil"
)

// Connection represents a generic protocol connection
type Connection interface {
	Validate() error
	GetProtocol() Protocol
}

func NewConnection(connection map[string]interface{}) (Connection, error) {
	fmt.Println("Creating new connection with connection:", connection)

	if connection == nil {
		return nil, ErrInvalidProtocolConnection
	}
	if connection["protocol"] == nil {
		return nil, ErrInvalidProtocol
	}

	protocol, err := NewProtocol(im_util.GetString(connection, "protocol"))
	if err != nil {
		return nil, err
	}

	var protocolConn Connection

	switch protocol {
	case ProtocolSFTP:
		protocolConn = &SFTPConnection{
			Host:          im_util.GetString(connection, "host"),
			Port:          im_util.GetInt(connection, "port"),
			Username:      im_util.GetString(connection, "username"),
			Password:      im_util.GetString(connection, "password"),
			PrivateKey:    im_util.GetString(connection, "private_key"),
			KeyPassphrase: im_util.GetString(connection, "key_passphrase"),
		}
	case ProtocolFTP:
		protocolConn = &FTPConnection{
			Host:        im_util.GetString(connection, "host"),
			Port:        im_util.GetInt(connection, "port"),
			Username:    im_util.GetString(connection, "username"),
			Password:    im_util.GetString(connection, "password"),
			PassiveMode: im_util.GetBool(connection, "passive_mode"),
		}
	case ProtocolS3:
		protocolConn = &S3Connection{
			Region:          im_util.GetString(connection, "region"),
			Bucket:          im_util.GetString(connection, "bucket"),
			AccessKeyID:     im_util.GetString(connection, "access_key_id"),
			SecretAccessKey: im_util.GetString(connection, "secret_access_key"),
			Endpoint:        im_util.GetString(connection, "endpoint"),
			UseSSL:          im_util.GetBool(connection, "use_ssl"),
		}
	}

	if protocolConn == nil {
		return nil, ErrInvalidProtocol
	}

	return protocolConn, nil
}
