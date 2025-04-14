package protocol

// Connection represents a generic protocol connection
type Connection interface {
	Validate() error
}

type NetworkConnection interface {
	Connection
	GetHost() string
	GetPort() int
	GetUsername() string
	GetPassword() string
}
