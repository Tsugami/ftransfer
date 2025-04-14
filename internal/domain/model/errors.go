package model

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

	// ErrEmptyConnectorID is returned when connector ID is empty
	ErrEmptyConnectorID = errors.New("connector ID cannot be empty")

	// ErrEmptySourceFolder is returned when source folder ID is empty
	ErrEmptySourceFolder = errors.New("source folder ID cannot be empty")

	// ErrEmptyDestinationFolder is returned when destination folder ID is empty
	ErrEmptyDestinationFolder = errors.New("destination folder ID cannot be empty")

	// ErrEmptyPostTransferPath is returned when post-transfer path is empty
	ErrEmptyPostTransferPath = errors.New("post-transfer path cannot be empty")

	ErrTransferNotFound         = errors.New("transfer not found")
	ErrFolderNotFound           = errors.New("folder not found")
	ErrConnectorNotFound        = errors.New("connector not found")
	ErrInvalidConnection        = errors.New("invalid connection configuration")
	ErrInvalidPath              = errors.New("invalid path")
	ErrInvalidMatch             = errors.New("invalid match pattern")
	ErrInvalidSourceFolder      = errors.New("invalid source folder")
	ErrInvalidDestinationFolder = errors.New("invalid destination folder")
)
