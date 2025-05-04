package protocol

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
		return ErrEmptySFTPUsername
	}
	if c.Port == 0 {
		return ErrEmptySFTPUsername
	}
	if c.Username == "" {
		return ErrEmptySFTPUsername
	}
	if c.Password == "" && c.PrivateKey == "" {
		return ErrEmptySFTPCredentials
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

func (c *SFTPConnection) GetJson() map[string]interface{} {
	return map[string]interface{}{
		"protocol":       c.GetProtocol(),
		"host":           c.Host,
		"port":           c.Port,
		"username":       c.Username,
		"password":       c.Password,
		"private_key":    c.PrivateKey,
		"key_passphrase": c.KeyPassphrase,
	}
}
