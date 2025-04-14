package model

// Folder represents a directory in a file system
type Folder struct {
	Base
	ConnectorID   string
	DirectoryPath string
}

// NewFolder creates a new Folder instance
func NewFolder(connectorID string, directoryPath string) *Folder {
	return &Folder{
		Base:          NewBase(),
		ConnectorID:   connectorID,
		DirectoryPath: directoryPath,
	}
}

// Validate performs validation on the folder
func (f *Folder) Validate() error {
	if f.ConnectorID == "" {
		return ErrEmptyConnectorID
	}
	if f.DirectoryPath == "" {
		return ErrEmptyDirectoryPath
	}
	return nil
}
