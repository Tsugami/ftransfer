package protocol

import "errors"

var (
	ErrInvalidProtocolConnection = errors.New("invalid protocol connection: the connection configuration is invalid or missing required fields")
	ErrInvalidProtocol           = errors.New("invalid protocol (field: protocol_connection.protocol): must be one of: sftp, ftp, s3")
	ErrEmptyFTPHost              = errors.New("host is required for FTP connection: specify the FTP server hostname or IP address")
	ErrEmptyFTPPort              = errors.New("port is required for FTP connection: specify the FTP server port number")
	ErrEmptyFTPUsername          = errors.New("username is required for FTP connection: specify the FTP user account name")
	ErrEmptyFTPPassword          = errors.New("password is required for FTP connection: specify the FTP user account password")
	ErrEmptySFTPUsername         = errors.New("username is required for SFTP connection: specify the SFTP user account name")
	ErrEmptySFTPCredentials      = errors.New("credentials are required for SFTP connection: specify either password or private key")
	ErrEmptyS3Bucket             = errors.New("bucket is required for S3 connection: specify the S3 bucket name")
	ErrEmptyS3AccessKeyID        = errors.New("access key id is required for S3 connection: specify the AWS access key ID")
	ErrEmptyS3SecretAccessKey    = errors.New("secret access key is required for S3 connection: specify the AWS secret access key")
	ErrEmptyS3Region             = errors.New("region is required for S3 connection: specify the AWS region (e.g. us-east-1)")
	ErrEmptyS3Endpoint           = errors.New("endpoint is required for S3 connection: specify the S3 endpoint URL")
	ErrEmptyS3UseSSL             = errors.New("use ssl is required for S3 connection: specify whether to use SSL/TLS encryption")
)
