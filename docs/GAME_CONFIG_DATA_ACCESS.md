# Game Configuration Data Access Layer

## Overview
Comprehensive data access layer for managing game configuration data with Redis caching and MongoDB fallback. All game configuration collections are read-only to serve game logic.

## Database Collections

### 1. **bullets**
Stores bullet/ammunition configurations
```json
{
  "_id": {...},
  "game_name": "ocean_hunter_v1",
  "data": {
    "bullets": [
      {
        "bullet_id": 1,
        "name": "Normal Bullet",
        "cost": 10,
        "damage": 20
      }
    ]
  }
}
```

### 2. **config**
Stores betting levels and general game parameters
```json
{
  "_id": {...},
  "game_name": "ocean_hunter_v1",
  "data": {
    "min_bet": 10,
    "max_bet": 1000,
    "bet_levels": [10, 25, 50, 100],
    "game_duration": 300,
    "max_players": 8,
    "room_capacity": 100
  }
}
```

### 3. **features**
Stores custom features per game (special skills, rewards, multipliers)
```json
{
  "_id": {...},
  "game_name": "ocean_hunter_v1",
  "data": {
    "special_skills": [
      {
        "skill_id": 1,
        "skill_name": "Double Damage",
        "cost": 100,
        "cooldown": 5000,
        "effect": "2x_damage_10s"
      }
    ],
    "special_rewards": [...],
    "multipliers": [...]
  }
}
```

### 4. **paths**
Stores fish movement paths
```json
{
  "_id": {...},
  "game_name": "ocean_hunter_v1",
  "data": {
    "paths": [
      {
        "path_id": 1,
        "path_name": "Top Left to Bottom Right",
        "coordinates": [
          {"x": 0, "y": 100, "z": 0},
          {"x": 50, "y": 50, "z": 0}
        ],
        "duration": 5000
      }
    ]
  }
}
```

### 5. **rtps**
Stores Return-to-Player (RTP) configuration by game
```json
{
  "_id": {...},
  "game_name": "ocean_hunter_v1",
  "data": {
    "rtp_rate": 96,
    "fish_rtp_map": {
      "1": 94,
      "2": 96,
      "3": 97
    },
    "bullet_rtp_map": {
      "1": 95,
      "2": 96
    }
  }
}
```

### 6. **types**
Stores fish type configurations
```json
{
  "_id": {...},
  "game_name": "ocean_hunter_v1",
  "data": {
    "fish_types": [
      {
        "fish_id": 1,
        "fish_name": "Goldfish",
        "hp": 10,
        "base_reward": 50,
        "rarity": "common",
        "spawn_rate": 40,
        "multiplier": 1
      }
    ]
  }
}
```

## Architecture

### 1. **Domain Models** (`internal/domain/games/game_base/models/db_model.go`)
- `BulletConfig` - Bullet configuration
- `GameConfig` - Game betting and parameters
- `GameFeatures` - Custom game features
- `GamePaths` - Fish movement paths
- `GameRTP` - RTP configuration
- `GameFishTypes` - Fish type definitions

### 2. **Port Interface** (`internal/domain/port/game_config_repository.go`)
```go
type GameConfigRepository interface {
    GetBulletConfig(ctx context.Context, gameName string) (*BulletConfig, error)
    GetGameConfig(ctx context.Context, gameName string) (*GameConfig, error)
    GetGameFeatures(ctx context.Context, gameName string) (*GameFeatures, error)
    GetGamePaths(ctx context.Context, gameName string) (*GamePaths, error)
    GetGameRTP(ctx context.Context, gameName string) (*GameRTP, error)
    GetGameFishTypes(ctx context.Context, gameName string) (*GameFishTypes, error)
}
```

### 3. **MongoDB Adapter** (`adapter/repository/mongo/game_config_repo.go`)
Direct access to MongoDB collections with error mapping:
- Returns `apperr.ErrNotFound` if document not found
- Returns database error if query fails

### 4. **Redis Cache Adapter** (`adapter/repository/redis/game_config_cache.go`)
Two-tier caching strategy:
1. **Check Redis** - If cached, return immediately
2. **MongoDB Fallback** - If not cached:
   - Query MongoDB
   - Return error if not found
   - Cache in Redis (TTL: 24 hours)
   - Return data

## Data Flow

```
User Request
    ↓
GameConfigUsecase
    ↓
Redis Cache Repository
    ├─ Redis Hit → Return cached data
    └─ Redis Miss → Query MongoDB
                      ├─ MongoDB Hit → Cache in Redis + Return
                      └─ MongoDB Miss → Return error code
```

## API Endpoints

### Get Bullet Configuration
```
GET /api/v1/game-config/:gameName/bullets
Response: BulletConfig with bullets array
```

### Get Game Config
```
GET /api/v1/game-config/:gameName/config
Response: GameConfig with betting levels and parameters
```

### Get Game Features
```
GET /api/v1/game-config/:gameName/features
Response: GameFeatures with skills, rewards, multipliers
```

### Get Game Paths
```
GET /api/v1/game-config/:gameName/paths
Response: GamePaths with path definitions
```

### Get Game RTP
```
GET /api/v1/game-config/:gameName/rtp
Response: GameRTP with RTP rates and mappings
```

### Get Fish Types
```
GET /api/v1/game-config/:gameName/fish-types
Response: GameFishTypes with fish type definitions
```

## Error Handling

All errors are returned with specific error codes:

| Error Code | Meaning |
|-----------|---------|
| `BULLET_CONFIG_NOT_FOUND` | Bullet config not found for game |
| `GAME_CONFIG_NOT_FOUND` | Game config not found |
| `GAME_FEATURES_NOT_FOUND` | Game features not found |
| `GAME_PATHS_NOT_FOUND` | Game paths not found |
| `GAME_RTP_NOT_FOUND` | Game RTP config not found |
| `GAME_FISH_TYPES_NOT_FOUND` | Fish types not found |

## Caching Strategy

- **TTL**: 24 hours (86400 seconds)
- **Cache Key Format**: `game_config:{configType}:{gameName}`
- **Invalidation**: Automatic after 24 hours
- **Graceful Degradation**: If Redis fails, directly queries MongoDB

## Example Usage

```go
// In usecase
config, err := uc.gameConfigUsecase.GetGameConfig(ctx, "ocean_hunter_v1")
if err != nil {
    // Handle error - returns error with specific code
    return err
}

// Use config data
minBet := config.Data.MinBet
maxBet := config.Data.MaxBet
```

## Performance Benefits

1. **Redis Caching**: Sub-millisecond response for cached data
2. **MongoDB Fallback**: Reliable fallback if cache fails
3. **24-Hour TTL**: Balances freshness with performance
4. **Lazy Loading**: Configuration loaded on-demand
5. **Automatic Caching**: Transparent caching after first MongoDB hit

## Integration Points

- Used by game logic services for rule extraction
- Referenced by skill system for special effects
- Used by fish spawning for type definitions
- Referenced by RTP calculation for probability mapping
- Used by path system for fish movement definitions
