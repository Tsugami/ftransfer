package transfer

// Transfer represents a file transfer operation
type Transfer struct {
	ID                    ID
	SourceDir             Directory
	DestinationDir        Directory
	PostTransferSourceDir Directory
}

// NewTransfer creates a new Transfer instance
func NewTransfer(sourceDir Directory, destinationDir Directory, postTransferSourceDir Directory) *Transfer {
	return &Transfer{
		SourceDir:             sourceDir,
		DestinationDir:        destinationDir,
		PostTransferSourceDir: postTransferSourceDir,
	}
}

// Validate performs validation on the transfer
func (t *Transfer) Validate() error {
	if !t.SourceDir.IsValid() {
		return ErrEmptySourceDir
	}

	if !t.DestinationDir.IsValid() {
		return ErrEmptyDestinationDir
	}

	if !t.PostTransferSourceDir.IsValid() {
		return ErrEmptyPostTransferDir
	}
	return nil
}
