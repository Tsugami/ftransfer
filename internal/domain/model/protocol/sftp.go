package protocol

import "github.com/Tsugami/ftransfer/internal/domain/errs"

// SFTPConnection represents an SFTP connection
type SFTPConnection struct {
	Host          string
	Port          int
	Username      string
	Password      string
	PrivateKey    string
	KeyPassphrase string
}

// Validate performs validation on the SFTP connection
func (c *SFTPConnection) Validate() error {
	if c.Host == "" {
		return errs.ErrEmptyHost
	}
	if c.Port == 0 {
		return errs.ErrEmptyPort
	}
	if c.Username == "" {
		return errs.ErrEmptyUsername
	}
	if c.Password == "" && c.PrivateKey == "" {
		return errs.ErrEmptyCredentials
	}
	return nil
}

// GetHost returns the host
func (c *SFTPConnection) GetHost() string {
	return c.Host
}

// GetPort returns the port
func (c *SFTPConnection) GetPort() int {
	return c.Port
}

// GetUsername returns the username
func (c *SFTPConnection) GetUsername() string {
	return c.Username
}

// GetPassword returns the password
func (c *SFTPConnection) GetPassword() string {
	return c.Password
}

func (c *SFTPConnection) GetProtocol() Protocol {
	return ProtocolSFTP
}
