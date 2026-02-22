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
	URI      string
	Database string
	Timeout  int
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	CacheTTL int // Cache TTL in seconds
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvInt("SERVER_PORT", 8080),
		},
		Mongo: MongoConfig{
			URI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database: getEnv("MONGO_DB", "fishing_game"),
			Timeout:  getEnvInt("MONGO_TIMEOUT", 10),
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvInt("REDIS_DB", 0),
			CacheTTL: getEnvInt("REDIS_CACHE_TTL", 86400), // 24 hours default
		},
	}
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
