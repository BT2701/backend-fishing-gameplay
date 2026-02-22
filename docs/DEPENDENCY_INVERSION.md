# Dependency Inversion Principle Implementation

## Overview

Proyek ini mengimplementasikan Dependency Inversion Principle (DIP) melalui serangkaian interfaces yang mendefinisikan contracts untuk berbagai layanan infrastruktur. Ini memastikan bahwa high-level modules (business logic) tidak bergantung pada low-level modules (concrete implementations), tetapi keduanya bergantung pada abstractions (interfaces).

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Application Layer                        │
│                   (Usecase, Handlers)                       │
└─────────────────────────────────────────────────────────────┘
                            ▲
                            │ depends on
                            │
┌─────────────────────────────────────────────────────────────┐
│              Contract/Interface Layer                       │
│  ┌──────────────┐  ┌──────────┐  ┌────────┐  ┌──────────┐  │
│  │ Logger       │  │ Config   │  │ Server │  │ Database │  │
│  │ Cache        │  │ Factory  │  │ Factory│  │ Factory  │  │
│  └──────────────┘  └──────────┘  └────────┘  └──────────┘  │
└─────────────────────────────────────────────────────────────┘
                            ▲
                            │ implements
                            │
┌─────────────────────────────────────────────────────────────┐
│              Adapter/Implementation Layer                   │
│  ┌──────────────┐  ┌──────────┐  ┌────────┐  ┌──────────┐  │
│  │ ZapLogger    │  │ Config   │  │ Fiber  │  │ Mongo    │  │
│  │ Adapter      │  │ Getters  │  │ Server │  │ Database │  │
│  │ RedisCache   │  │          │  │ Fiber  │  │ Redis    │  │
│  │ Adapter      │  │          │  │ Factory│  │ Cache    │  │
│  └──────────────┘  └──────────┘  └────────┘  └──────────┘  │
└─────────────────────────────────────────────────────────────┘
```

## Contract Interfaces

### 1. Logger Interface
**File**: `internal/infrastructure/contract/logger.go`

```go
type Logger interface {
    Debug(msg string, fields ...interface{})
    Info(msg string, fields ...interface{})
    Warn(msg string, fields ...interface{})
    Error(msg string, fields ...interface{})
    Fatal(msg string, fields ...interface{})
    Sync() error
}
```

**Implementation**: `internal/infrastructure/logger/adapter.go`
- `ZapLoggerAdapter` - Wraps `*zap.Logger`

**Benefits**:
- Easy to swap Logger implementation
- No coupling to zap library in business logic
- Can mock for testing

### 2. Config Interface
**File**: `internal/infrastructure/contract/config.go`

```go
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
```

**Implementation**: `internal/infrastructure/config/config.go`
- `Config` struct with getter methods

**Benefits**:
- Can inject different configurations per environment
- Type-safe access to configuration values
- Easy to test with mock configs

### 3. Database Interface
**File**: `internal/infrastructure/contract/database.go`

```go
type Database interface {
    Ping(ctx context.Context) error
    Close(ctx context.Context) error
    GetNative() interface{}
}

type DatabaseFactory interface {
    CreateDatabase(cfg Config, logger Logger) (Database, error)
    CloseDatabase(db Database) error
}
```

**Implementation**: `adapter/database/mongo.go`
- `MongoDatabase` - Wraps `*mongo.Client`
- `MongoDatabaseFactory` - Creates MongoDB connections with retry logic

**Benefits**:
- Can switch between different databases
- Retry logic abstracted away
- No direct dependency on MongoDB driver

### 4. Cache Interface
**File**: `internal/infrastructure/contract/cache.go`

```go
type Cache interface {
    Ping(ctx context.Context) error
    Close() error
    GetNative() interface{}
}

type CacheFactory interface {
    CreateCache(cfg Config, logger Logger) (Cache, error)
    CloseCache(cache Cache) error
}
```

**Implementation**: `adapter/database/redis.go`
- `RedisCache` - Wraps `*redis.Client`
- `RedisCacheFactory` - Creates Redis connections with retry logic

**Benefits**:
- Can switch between different caching solutions
- Redis details abstracted away
- Testable with mock caches

### 5. Server Interface
**File**: `internal/infrastructure/contract/server.go`

```go
type Server interface {
    Start() error
    Stop() error
    GetNative() interface{}
}

type ServerFactory interface {
    CreateServer(host string, port int, logger Logger) (Server, error)
}
```

**Implementation**: `adapter/server/fiber.go`
- `FiberServer` - Wraps `*fiber.App`
- `FiberServerFactory` - Creates Fiber servers

**Benefits**:
- Can switch HTTP frameworks
- Server lifecycle management abstracted

## How to Use

### Basic Usage

```go
// Load configuration
cfg := config.Load()

// Create adapters
loggerAdapter := logger.NewZapLoggerAdapter(zapLogger)
dbFactory := &mongo.MongoDatabaseFactory{}
cacheFactory := &redis.RedisCacheFactory{}

// Create database and cache using factories
db, err := dbFactory.CreateDatabase(cfg, loggerAdapter)
if err != nil {
    loggerAdapter.Fatal("Failed to create database", err)
}

cache, err := cacheFactory.CreateCache(cfg, loggerAdapter)
if err != nil {
    loggerAdapter.Fatal("Failed to create cache", err)
}

// Get native clients for repository initialization
mongoClient := db.GetNative().(*mongo.Client)
redisClient := cache.GetNative().(*redis.Client)

// Create repositories using native clients
roomRepo := mongo.NewRoomRepository(mongoClient.Database(cfg.GetMongoDatabase()))
rtpRepo := redis.NewRTPRepository(redisClient)
```

## Benefits of DIP

1. **Testability**: Easy to mock dependencies for unit testing
2. **Flexibility**: Can swap implementations without changing business logic
3. **Maintainability**: Clear contracts make the code easier to understand
4. **Decoupling**: High-level modules don't depend on low-level details
5. **Extensibility**: Easy to add new implementations

## Example: Adding a New Logger Implementation

```go
// 1. Create a new logger (e.g., Structured)
type StructuredLogger struct {
    // ...
}

// 2. Implement the Logger interface
func (s *StructuredLogger) Debug(msg string, fields ...interface{}) { /* ... */ }
func (s *StructuredLogger) Info(msg string, fields ...interface{})  { /* ... */ }
// ... other methods

// 3. Use it in main.go
loggerAdapter := &StructuredLogger{}
dbFactory.CreateDatabase(cfg, loggerAdapter)
```

## Example: Adding a New Database Implementation

```go
// 1. Create a new database adapter (e.g., PostgreSQL)
type PostgreSQLDatabase struct {
    client *sql.DB
}

// 2. Implement the Database interface
func (p *PostgreSQLDatabase) Ping(ctx context.Context) error { /* ... */ }
func (p *PostgreSQLDatabase) Close(ctx context.Context) error { /* ... */ }
// ... other methods

// 3. Create a factory
type PostgreSQLFactory struct{}
func (f *PostgreSQLFactory) CreateDatabase(cfg Config, logger Logger) (Database, error) {
    // ... create and return PostgreSQL database
}

// 4. Use it in main.go
dbFactory := &PostgreSQLFactory{}
db, err := dbFactory.CreateDatabase(cfg, loggerAdapter)
```

## Current Implementation Status

✅ **Implemented**:
- Logger interface + ZapLogger adapter
- Config interface + getter methods
- Database interface + MongoDB factory  
- Cache interface + Redis factory
- Server interface (basic Fiber implementation)
- Retry logic abstraction in persistence layer

⚠️ **Next Steps**:
- Refactor main.go to fully use factory pattern
- Add more sophisticated server factory
- Add repository factory interfaces
- Complete server adapter implementation

## Testing

All interfaces are testable using mock implementations:

```go
// Mock logger for testing
type MockLogger struct{}
func (m *MockLogger) Debug(msg string, fields ...interface{}) {}
func (m *MockLogger) Info(msg string, fields ...interface{})  {}
// ... other methods

// Mock config for testing
type MockConfig struct{}
func (m *MockConfig) GetServerHost() string { return "localhost" }
func (m *MockConfig) GetServerPort() int    { return 8080 }
// ... other methods
```

## References

- [Dependency Inversion Principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle)
- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go best practices](https://golang.org/doc/effective_go)
