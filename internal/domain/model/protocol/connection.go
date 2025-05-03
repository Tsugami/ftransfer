package protocol

import "github.com/Tsugami/ftransfer/internal/domain/errs"

// Connection represents a generic protocol connection
type Connection interface {
	Validate() error
	GetProtocol() Protocol
}

func NewConnection(connection map[string]interface{}) (Connection, error) {
	protocolInput := connection["protocol"].(string)
	protocol, err := NewProtocol(protocolInput)
	if err != nil {
		return nil, err
	}

	switch protocol {
	case ProtocolSFTP:
		return &SFTPConnection{}, nil
	case ProtocolFTP:
		return &FTPConnection{}, nil
	case ProtocolS3:
		return &S3Connection{}, nil
	}

	return nil, errs.ErrInvalidProtocol
}
