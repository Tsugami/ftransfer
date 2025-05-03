package transfer

import "errors"

var (
	ErrEmptySourceDir       = errors.New("source directory is empty")
	ErrEmptyDestinationDir  = errors.New("destination directory is empty")
	ErrEmptyPostTransferDir = errors.New("post transfer directory is empty")
	ErrCreateTransfer       = errors.New("failed to create transfer")
	ErrListTransfers        = errors.New("failed to list transfers")
	ErrGetTransfer          = errors.New("failed to get transfer")
	ErrUpdateTransfer       = errors.New("failed to update transfer")
	ErrDeleteTransfer       = errors.New("failed to delete transfer")
)
