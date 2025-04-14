package model

// Protocol represents the transfer protocol type
type Protocol string

const (
	ProtocolSFTP  Protocol = "SFTP"
	ProtocolFTP   Protocol = "FTP"
	ProtocolS3    Protocol = "S3"
	ProtocolLOCAL Protocol = "LOCAL"
)

// String returns the string representation of the protocol
func (p Protocol) String() string {
	return string(p)
}

// IsValid checks if the protocol is valid
func (p Protocol) IsValid() bool {
	switch p {
	case ProtocolSFTP, ProtocolFTP, ProtocolS3, ProtocolLOCAL:
		return true
	default:
		return false
	}
}
