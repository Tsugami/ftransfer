package transfer

import (
	"strings"

	"github.com/Tsugami/ftransfer/internal/storage_provider"
)

type Directory []string

func (d *Directory) String(system storage_provider.FileSystemType) string {
	switch system {
	case storage_provider.FileSystemUNIX:
		return strings.Join(*d, "/")
	case storage_provider.FileSystemWINDOWS:
		return strings.Join(*d, "\\")
	}

	return ""
}

func (d *Directory) IsValid() bool {
	return len(*d) != 0
}
