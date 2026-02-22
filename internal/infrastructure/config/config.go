package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server ServerConfig
	Mongo  MongoConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	Port int
	Host string
}

type MongoConfig struct {
	URI        string
	Database   string
	Timeout    int
	MaxRetries int
	RetryDelay int // milliseconds
}

type RedisConfig struct {
	Addr       string
	Password   string
	DB         int
	CacheTTL   int // Cache TTL in seconds
	MaxRetries int
	RetryDelay int // milliseconds
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvInt("SERVER_PORT", 8080),
		},
		Mongo: MongoConfig{
			URI:        getEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database:   getEnv("MONGO_DB", "fishing_game"),
			Timeout:    getEnvInt("MONGO_TIMEOUT", 10),
			MaxRetries: getEnvInt("MONGO_MAX_RETRIES", 5),
			RetryDelay: getEnvInt("MONGO_RETRY_DELAY", 1000),
		},
		Redis: RedisConfig{
			Addr:       getEnv("REDIS_ADDR", "localhost:6379"),
			Password:   getEnv("REDIS_PASSWORD", ""),
			DB:         getEnvInt("REDIS_DB", 0),
			CacheTTL:   getEnvInt("REDIS_CACHE_TTL", 86400), // 24 hours default
			MaxRetries: getEnvInt("REDIS_MAX_RETRIES", 5),
			RetryDelay: getEnvInt("REDIS_RETRY_DELAY", 1000),
		},
	}
}

// Server configuration methods
func (c *Config) GetServerHost() string {
	return c.Server.Host
}

func (c *Config) GetServerPort() int {
	return c.Server.Port
}

// MongoDB configuration methods
func (c *Config) GetMongoURI() string {
	return c.Mongo.URI
}

func (c *Config) GetMongoDatabase() string {
	return c.Mongo.Database
}

func (c *Config) GetMongoTimeout() int {
	return c.Mongo.Timeout
}

func (c *Config) GetMongoMaxRetries() int {
	return c.Mongo.MaxRetries
}

func (c *Config) GetMongoRetryDelay() int {
	return c.Mongo.RetryDelay
}

// Redis configuration methods
func (c *Config) GetRedisAddr() string {
	return c.Redis.Addr
}

func (c *Config) GetRedisPassword() string {
	return c.Redis.Password
}

func (c *Config) GetRedisDB() int {
	return c.Redis.DB
}

func (c *Config) GetRedisCacheTTL() int {
	return c.Redis.CacheTTL
}

func (c *Config) GetRedisMaxRetries() int {
	return c.Redis.MaxRetries
}

func (c *Config) GetRedisRetryDelay() int {
	return c.Redis.RetryDelay
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}
