package contract

// Server defines the interface for HTTP server operations
type Server interface {
	// Start starts the HTTP server
	Start() error

	// Stop stops the HTTP server
	Stop() error

	// GetNative returns the native HTTP framework instance
	GetNative() interface{}
}

// ServerFactory defines the interface for creating HTTP servers
type ServerFactory interface {
	// CreateServer creates and returns a new server instance
	CreateServer(host string, port int, logger Logger) (Server, error)
}
