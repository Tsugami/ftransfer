package model

import "github.com/Tsugami/ftransfer/internal/domain/errs"

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
		return errs.ErrEmptySourceFolder
	}
	if t.DestinationFolderID == "" {
		return errs.ErrEmptyDestinationFolder
	}
	if t.PostTransferSourcePath == "" {
		return errs.ErrEmptyPostTransferPath
	}
	return nil
}
