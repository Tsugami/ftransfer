package protocol

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
		return ErrEmptyFTPHost
	}
	if c.Port == 0 {
		return ErrEmptyFTPPort
	}
	if c.Username == "" {
		return ErrEmptyFTPUsername
	}
	if c.Password == "" {
		return ErrEmptyFTPPassword
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

func (c *FTPConnection) GetJson() map[string]interface{} {
	return map[string]interface{}{
		"protocol":     c.GetProtocol(),
		"host":         c.Host,
		"port":         c.Port,
		"username":     c.Username,
		"password":     c.Password,
		"passive_mode": c.PassiveMode,
	}
}
