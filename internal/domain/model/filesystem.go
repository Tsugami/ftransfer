package model

// FileSystemType represents the type of file system
type FileSystemType string

const (
	FileSystemPOSIX   FileSystemType = "POSIX"
	FileSystemWINDOWS FileSystemType = "WINDOWS"
)

// String returns the string representation of the file system type
func (fs FileSystemType) String() string {
	return string(fs)
}

// IsValid checks if the file system type is valid
func (fs FileSystemType) IsValid() bool {
	switch fs {
	case FileSystemPOSIX, FileSystemWINDOWS:
		return true
	default:
		return false
	}
}
