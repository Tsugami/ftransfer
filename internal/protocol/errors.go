package protocol

import "errors"

var (
	ErrInvalidProtocolConnection = errors.New("invalid protocol connection")
	ErrInvalidProtocol           = errors.New("invalid protocol (field: protocol_connection.protocol)")
	ErrEmptyFTPHost              = errors.New("host is required")
	ErrEmptyFTPPort              = errors.New("port is required")
	ErrEmptyFTPUsername          = errors.New("username is required")
	ErrEmptyFTPPassword          = errors.New("password is required")
	ErrEmptySFTPUsername         = errors.New("username is required")
	ErrEmptySFTPCredentials      = errors.New("credentials are required")
	ErrEmptyS3Bucket             = errors.New("bucket is required")
	ErrEmptyS3AccessKeyID        = errors.New("access key id is required")
	ErrEmptyS3SecretAccessKey    = errors.New("secret access key is required")
	ErrEmptyS3Region             = errors.New("region is required")
	ErrEmptyS3Endpoint           = errors.New("endpoint is required")
	ErrEmptyS3UseSSL             = errors.New("use ssl is required")
)
