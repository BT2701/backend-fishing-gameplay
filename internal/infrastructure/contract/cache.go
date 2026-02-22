package contract

import "context"

// Cache defines the interface for cache operations
type Cache interface {
	// Ping verifies the connection to the cache
	Ping(ctx context.Context) error

	// Close closes the cache connection
	Close() error

	// GetNative returns the native cache client instance
	GetNative() interface{}
}

// CacheFactory defines the interface for creating cache connections
type CacheFactory interface {
	// CreateCache creates and returns a cache connection with retry logic
	CreateCache(cfg Config, logger Logger) (Cache, error)

	// CloseCache closes the cache connection
	CloseCache(cache Cache) error
}
