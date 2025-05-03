package errs

import "errors"

var (
	// ErrEmptyName is returned when a name is empty
	ErrEmptyName = errors.New("name cannot be empty")

	// ErrEmptyProtocol is returned when a protocol is empty
	ErrEmptyProtocol = errors.New("protocol cannot be empty")

	// ErrInvalidProtocol is returned when a protocol is invalid
	ErrInvalidProtocol = errors.New("invalid protocol")

	// ErrEmptyProtocolConnection is returned when protocol connection details are empty
	ErrEmptyProtocolConnection = errors.New("protocol connection cannot be empty")

	// ErrEmptyFileSystem is returned when file system type is empty
	ErrEmptyFileSystem = errors.New("file system type cannot be empty")

	// ErrInvalidFileSystem is returned when file system type is invalid
	ErrInvalidFileSystem = errors.New("invalid file system type")

	// ErrEmptyDirectoryPath is returned when directory path is empty
	ErrEmptyDirectoryPath = errors.New("directory path cannot be empty")

	// ErrEmptyStorageProviderID is returned when storage provider ID is empty
	ErrEmptyStorageProviderID = errors.New("storage provider ID cannot be empty")

	// ErrEmptySourceFolder is returned when source folder ID is empty
	ErrEmptySourceFolder = errors.New("source folder ID cannot be empty")

	// ErrEmptyDestinationFolder is returned when destination folder ID is empty
	ErrEmptyDestinationFolder = errors.New("destination folder ID cannot be empty")

	// ErrEmptyPostTransferPath is returned when post-transfer path is empty
	ErrEmptyPostTransferPath = errors.New("post-transfer path cannot be empty")

	ErrTransferNotFound         = errors.New("transfer not found")
	ErrFolderNotFound           = errors.New("folder not found")
	ErrStorageProviderNotFound  = errors.New("storageProvider not found")
	ErrInvalidConnection        = errors.New("invalid connection configuration")
	ErrInvalidPath              = errors.New("invalid path")
	ErrInvalidMatch             = errors.New("invalid match pattern")
	ErrInvalidSourceFolder      = errors.New("invalid source folder")
	ErrInvalidDestinationFolder = errors.New("invalid destination folder")

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
