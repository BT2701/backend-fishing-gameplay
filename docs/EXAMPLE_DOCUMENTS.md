# Game Configuration - Example MongoDB Documents

These are complete example documents for each collection in MongoDB.

## 1. Bullets Collection

```json
{
  "_id": ObjectId("60a1b2c3d4e5f6g7h8i9j0k1"),
  "game_name": "ocean_hunter_v1",
  "data": {
    "bullets": [
      {
        "bullet_id": 1,
        "name": "Pea Shot",
        "cost": 5,
        "damage": 10
      },
      {
        "bullet_id": 2,
        "name": "Cannon Ball",
        "cost": 15,
        "damage": 40
      },
      {
        "bullet_id": 3,
        "name": "Laser Beam",
        "cost": 30,
        "damage": 80
      },
      {
        "bullet_id": 4,
        "name": "Nuclear Bomb",
        "cost": 60,
        "damage": 150
      },
      {
        "bullet_id": 5,
        "name": "Holy Light",
        "cost": 100,
        "damage": 250
      }
    ]
  }
}
```

## 2. Config Collection

```json
{
  "_id": ObjectId("60a1b2c3d4e5f6g7h8i9j0k2"),
  "game_name": "ocean_hunter_v1",
  "data": {
    "min_bet": 10,
    "max_bet": 1000,
    "bet_levels": [10, 25, 50, 100, 250, 500, 1000],
    "game_duration": 300,
    "max_players": 8,
    "room_capacity": 100
  }
}
```

## 3. Features Collection

```json
{
  "_id": ObjectId("60a1b2c3d4e5f6g7h8i9j0k3"),
  "game_name": "ocean_hunter_v1",
  "data": {
    "special_skills": [
      {
        "skill_id": 1,
        "skill_name": "Double Damage",
        "cost": 100,
        "cooldown": 5000,
        "effect": "damage_multiplier_x2_10s"
      },
      {
        "skill_id": 2,
        "skill_name": "Slow Effect",
        "cost": 80,
        "cooldown": 8000,
        "effect": "fish_speed_x0.5_5s"
      },
      {
        "skill_id": 3,
        "skill_name": "Freeze",
        "cost": 120,
        "cooldown": 10000,
        "effect": "fish_immobilized_3s"
      },
      {
        "skill_id": 4,
        "skill_name": "Lucky Shot",
        "cost": 150,
        "cooldown": 12000,
        "effect": "guaranteed_capture_next_shoot"
      }
    ],
    "special_rewards": [
      {
        "reward_id": 1,
        "reward_name": "Jackpot",
        "amount": 10000,
        "chance": 5
      },
      {
        "reward_id": 2,
        "reward_name": "Bonus Round",
        "amount": 5000,
        "chance": 10
      },
      {
        "reward_id": 3,
        "reward_name": "Free Bullets",
        "amount": 50,
        "chance": 15
      }
    ],
    "multipliers": [
      {
        "fish_type": 5,
        "multiplier": 3
      },
      {
        "fish_type": 6,
        "multiplier": 5
      },
      {
        "fish_type": 7,
        "multiplier": 10
      }
    ]
  }
}
```

## 4. Paths Collection

```json
{
  "_id": ObjectId("60a1b2c3d4e5f6g7h8i9j0k4"),
  "game_name": "ocean_hunter_v1",
  "data": {
    "paths": [
      {
        "path_id": 1,
        "path_name": "Top to Bottom",
        "coordinates": [
          {"x": 0, "y": 100, "z": 0},
          {"x": 0, "y": 75, "z": 0},
          {"x": 0, "y": 50, "z": 0},
          {"x": 0, "y": 25, "z": 0},
          {"x": 0, "y": 0, "z": 0}
        ],
        "duration": 5000
      },
      {
        "path_id": 2,
        "path_name": "Diagonal Sweep",
        "coordinates": [
          {"x": 0, "y": 100, "z": 0},
          {"x": 25, "y": 75, "z": 0},
          {"x": 50, "y": 50, "z": 0},
          {"x": 75, "y": 25, "z": 0},
          {"x": 100, "y": 0, "z": 0}
        ],
        "duration": 8000
      },
      {
        "path_id": 3,
        "path_name": "Figure Eight",
        "coordinates": [
          {"x": 0, "y": 50, "z": 0},
          {"x": 30, "y": 80, "z": 0},
          {"x": 50, "y": 50, "z": 0},
          {"x": 70, "y": 20, "z": 0},
          {"x": 50, "y": 50, "z": 0},
          {"x": 30, "y": 80, "z": 0},
          {"x": 0, "y": 50, "z": 0}
        ],
        "duration": 10000
      },
      {
        "path_id": 4,
        "path_name": "Spiral Down",
        "coordinates": [
          {"x": 50, "y": 100, "z": 0},
          {"x": 70, "y": 80, "z": 0},
          {"x": 70, "y": 60, "z": 0},
          {"x": 50, "y": 40, "z": 0},
          {"x": 30, "y": 40, "z": 0},
          {"x": 30, "y": 20, "z": 0},
          {"x": 50, "y": 0, "z": 0}
        ],
        "duration": 12000
      }
    ]
  }
}
```

## 5. RTP Collection

```json
{
  "_id": ObjectId("60a1b2c3d4e5f6g7h8i9j0k5"),
  "game_name": "ocean_hunter_v1",
  "data": {
    "rtp_rate": 96,
    "fish_rtp_map": {
      "1": 92,
      "2": 94,
      "3": 95,
      "4": 96,
      "5": 97,
      "6": 98,
      "7": 99,
      "8": 99
    },
    "bullet_rtp_map": {
      "1": 95,
      "2": 96,
      "3": 96,
      "4": 97,
      "5": 98
    }
  }
}
```

## 6. Types Collection

```json
{
  "_id": ObjectId("60a1b2c3d4e5f6g7h8i9j0k6"),
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
      },
      {
        "fish_id": 2,
        "fish_name": "Catfish",
        "hp": 25,
        "base_reward": 150,
        "rarity": "uncommon",
        "spawn_rate": 25,
        "multiplier": 2
      },
      {
        "fish_id": 3,
        "fish_name": "Tuna",
        "hp": 40,
        "base_reward": 300,
        "rarity": "uncommon",
        "spawn_rate": 20,
        "multiplier": 2
      },
      {
        "fish_id": 4,
        "fish_name": "Shark",
        "hp": 100,
        "base_reward": 500,
        "rarity": "rare",
        "spawn_rate": 10,
        "multiplier": 3
      },
      {
        "fish_id": 5,
        "fish_name": "Dragon Fish",
        "hp": 200,
        "base_reward": 2000,
        "rarity": "epic",
        "spawn_rate": 3,
        "multiplier": 5
      },
      {
        "fish_id": 6,
        "fish_name": "Phoenix Fish",
        "hp": 300,
        "base_reward": 5000,
        "rarity": "epic",
        "spawn_rate": 2,
        "multiplier": 8
      },
      {
        "fish_id": 7,
        "fish_name": "Leviathan",
        "hp": 500,
        "base_reward": 10000,
        "rarity": "legendary",
        "spawn_rate": 1,
        "multiplier": 10
      },
      {
        "fish_id": 8,
        "fish_name": "Ancient God",
        "hp": 1000,
        "base_reward": 50000,
        "rarity": "legendary",
        "spawn_rate": 0.5,
        "multiplier": 20
      }
    ]
  }
}
```

## Creating Documents in MongoDB

Use these MongoDB shell commands to insert the documents:

```javascript
// Insert all documents
db.bullets.insertOne({
  "_id": ObjectId("60a1b2c3d4e5f6g7h8i9j0k1"),
  "game_name": "ocean_hunter_v1",
  "data": { /* ... */ }
});

db.config.insertOne({
  "_id": ObjectId("60a1b2c3d4e5f6g7h8i9j0k2"),
  "game_name": "ocean_hunter_v1",
  "data": { /* ... */ }
});

// ... repeat for other collections

// Create indexes for better query performance
db.bullets.createIndex({ "game_name": 1 });
db.config.createIndex({ "game_name": 1 });
db.features.createIndex({ "game_name": 1 });
db.paths.createIndex({ "game_name": 1 });
db.rtps.createIndex({ "game_name": 1 });
db.types.createIndex({ "game_name": 1 });
```

## Redis Cache Examples

After first request to each endpoint, Redis will contain:

```
# After GET /api/v1/game-config/ocean_hunter_v1/bullets
Key: game_config:bullets:ocean_hunter_v1
Value: {"_id":...,"game_name":"ocean_hunter_v1","data":{"bullets":[...]}}
TTL: 86400 seconds

# After GET /api/v1/game-config/ocean_hunter_v1/config
Key: game_config:config:ocean_hunter_v1
Value: {"_id":...,"game_name":"ocean_hunter_v1","data":{...}}
TTL: 86400 seconds

# And so on for each config type...
```

## Testing the Data Access Layer

### Test API Calls

```bash
# Get bullet configuration
curl -X GET "http://localhost:8080/api/v1/game-config/ocean_hunter_v1/bullets"

# Get game configuration
curl -X GET "http://localhost:8080/api/v1/game-config/ocean_hunter_v1/config"

# Get special features
curl -X GET "http://localhost:8080/api/v1/game-config/ocean_hunter_v1/features"

# Get fish paths
curl -X GET "http://localhost:8080/api/v1/game-config/ocean_hunter_v1/paths"

# Get RTP configuration
curl -X GET "http://localhost:8080/api/v1/game-config/ocean_hunter_v1/rtp"

# Get fish types
curl -X GET "http://localhost:8080/api/v1/game-config/ocean_hunter_v1/fish-types"
```

### Expected Responses

**Success (200 OK):**
```json
{
  "_id": {...},
  "game_name": "ocean_hunter_v1",
  "data": {...}
}
```

**Not Found (404):**
```json
{
  "error": "game config not found for game: ocean_hunter_v1"
}
```

**Invalid Request (400):**
```json
{
  "error": "game_name is required"
}
```

## Performance Verification

### First Request (Cache Miss)
- Time: ~20-50ms
- Flow: Redis (miss) → MongoDB (hit) → Redis (cache) → Return
- Next request will be < 1ms

### Subsequent Requests (Cache Hit)
- Time: < 1ms
- Flow: Redis (hit) → Return
- No MongoDB query needed

### Not Found Case
- Time: ~20-50ms
- Flow: Redis (miss) → MongoDB (miss) → Error
- Error information returned with specific code
