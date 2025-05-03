package protocol

import "github.com/Tsugami/ftransfer/internal/domain/errs"

// FTPConnection represents an FTP connection
type FTPConnection struct {
	Host        string
	Port        int
	Username    string
	Password    string
	PassiveMode bool
}

// Validate performs validation on the FTP connection
func (c *FTPConnection) Validate() error {
	if c.Host == "" {
		return errs.ErrEmptyHost
	}
	if c.Port == 0 {
		return errs.ErrEmptyPort
	}
	if c.Username == "" {
		return errs.ErrEmptyUsername
	}
	if c.Password == "" {
		return errs.ErrEmptyPassword
	}
	return nil
}

// GetHost returns the host
func (c *FTPConnection) GetHost() string {
	return c.Host
}

// GetPort returns the port
func (c *FTPConnection) GetPort() int {
	return c.Port
}

// GetUsername returns the username
func (c *FTPConnection) GetUsername() string {
	return c.Username
}

// GetPassword returns the password
func (c *FTPConnection) GetPassword() string {
	return c.Password
}

func (c *FTPConnection) GetProtocol() Protocol {
	return ProtocolFTP
}
