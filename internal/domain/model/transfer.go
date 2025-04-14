package model

// Transfer represents a file transfer operation
type Transfer struct {
	Base
	SourceFolderID         string
	DestinationFolderID    string
	PostTransferSourcePath string
}

// NewTransfer creates a new Transfer instance
func NewTransfer(sourceFolderID, destinationFolderID, postTransferSourcePath string) *Transfer {
	return &Transfer{
		Base:                   NewBase(),
		SourceFolderID:         sourceFolderID,
		DestinationFolderID:    destinationFolderID,
		PostTransferSourcePath: postTransferSourcePath,
	}
}

// Validate performs validation on the transfer
func (t *Transfer) Validate() error {
	if t.SourceFolderID == "" {
		return ErrEmptySourceFolder
	}
	if t.DestinationFolderID == "" {
		return ErrEmptyDestinationFolder
	}
	if t.PostTransferSourcePath == "" {
		return ErrEmptyPostTransferPath
	}
	return nil
}
