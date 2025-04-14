package protocol

import "errors"

var (
	// ErrEmptyHost is returned when the host is empty
	ErrEmptyHost = errors.New("host cannot be empty")
	// ErrEmptyPort is returned when the port is empty
	ErrEmptyPort = errors.New("port cannot be empty")
	// ErrEmptyUsername is returned when the username is empty
	ErrEmptyUsername = errors.New("username cannot be empty")
	// ErrEmptyPassword is returned when the password is empty
	ErrEmptyPassword = errors.New("password cannot be empty")
	// ErrEmptyCredentials is returned when neither password nor private key is provided
	ErrEmptyCredentials = errors.New("either password or private key must be provided")
	// ErrEmptyRegion is returned when the region is empty
	ErrEmptyRegion = errors.New("region cannot be empty")
	// ErrEmptyBucket is returned when the bucket is empty
	ErrEmptyBucket = errors.New("bucket cannot be empty")
	// ErrEmptyAccessKeyID is returned when the access key ID is empty
	ErrEmptyAccessKeyID = errors.New("access key ID cannot be empty")
	// ErrEmptySecretAccessKey is returned when the secret access key is empty
	ErrEmptySecretAccessKey = errors.New("secret access key cannot be empty")

	// ErrInvalidSFTPConnection is returned when the connection is not a valid SFTP connection
	ErrInvalidSFTPConnection = errors.New("protocol connection must be SFTPConnection for SFTP protocol")
	// ErrInvalidFTPConnection is returned when the connection is not a valid FTP connection
	ErrInvalidFTPConnection = errors.New("protocol connection must be FTPConnection for FTP protocol")
	// ErrInvalidS3Connection is returned when the connection is not a valid S3 connection
	ErrInvalidS3Connection = errors.New("protocol connection must be S3Connection for S3 protocol")
)
