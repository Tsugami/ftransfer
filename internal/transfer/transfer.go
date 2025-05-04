package transfer

import "github.com/Tsugami/ftransfer/internal/storage_provider"

// Transfer represents a file transfer operation
type Transfer struct {
	ID                           ID
	SourceDir                    Directory
	DestinationDir               Directory
	PostTransferSourceDir        Directory
	SourceStorageProviderID      storage_provider.ID
	DestinationStorageProviderID storage_provider.ID
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

	if t.SourceStorageProviderID == "" {
		return ErrEmptySourceStorageProviderID
	}

	if t.DestinationStorageProviderID == "" {
		return ErrEmptyDestinationStorageProviderID
	}
	return nil
}
