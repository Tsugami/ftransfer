package storage_provider

import "errors"

var (
	ErrInvalidStorageProvider  = errors.New("invalid storage provider")
	ErrCreateStorageProvider   = errors.New("failed to create storage provider")
	ErrUpdateStorageProvider   = errors.New("failed to update storage provider")
	ErrDeleteStorageProvider   = errors.New("failed to delete storage provider")
	ErrGetStorageProvider      = errors.New("failed to get storage provider")
	ErrListStorageProviders    = errors.New("failed to list storage providers")
	ErrStorageProviderNotFound = errors.New("storage provider not found")
	ErrStorageProviderExists   = errors.New("storage provider already exists")
	ErrEmptyName               = errors.New("name is required")
	ErrEmptyProtocolConnection = errors.New("protocol connection is required")
	ErrEmptyFileSystem         = errors.New("file system is required")
	ErrInvalidFileSystem       = errors.New("invalid file system")
)
