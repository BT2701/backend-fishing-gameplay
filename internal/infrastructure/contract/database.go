package contract

import "context"

// Database defines the interface for database operations
type Database interface {
	// Ping verifies the connection to the database
	Ping(ctx context.Context) error

	// Close closes the database connection
	Close(ctx context.Context) error

	// GetDatabase returns the native database instance
	GetNative() interface{}
}

// DatabaseFactory defines the interface for creating database connections
type DatabaseFactory interface {
	// CreateDatabase creates and returns a database connection with retry logic
	CreateDatabase(cfg Config, logger Logger) (Database, error)

	// CloseDatabase closes the database connection
	CloseDatabase(db Database) error
}
