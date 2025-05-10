package storage_provider_client

import (
	"context"
	"io"
)

type StorageProviderClient interface {
	ListFiles(ctx context.Context, path string) ([]string, error)
	ReadFile(ctx context.Context, path string) (io.ReadCloser, error)
	WriteFile(ctx context.Context, path string, data io.Reader) error
	MoveFile(ctx context.Context, sourcePath string, destinationPath string) error
}
