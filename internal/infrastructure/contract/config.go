package contract

// Config defines the interface for application configuration
type Config interface {
	// Server configuration
	GetServerHost() string
	GetServerPort() int

	// MongoDB configuration
	GetMongoURI() string
	GetMongoDatabase() string
	GetMongoTimeout() int
	GetMongoMaxRetries() int
	GetMongoRetryDelay() int

	// Redis configuration
	GetRedisAddr() string
	GetRedisPassword() string
	GetRedisDB() int
	GetRedisCacheTTL() int
	GetRedisMaxRetries() int
	GetRedisRetryDelay() int
}
