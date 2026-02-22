# Game Configuration Data Models

## Structure Overview

All MongoDB collections follow a standardized document structure:

```go
type GameConfigDocument struct {
    ID       map[string]interface{}  // MongoDB document ID
    GameName string                  // Game identifier (e.g., "ocean_hunter_v1")
    Data     interface{}             // Game-specific configuration
}
```

## Detailed Models

### BulletConfig - Ammunition Configuration
Controls bullet properties, costs, and damage values.

```go
type BulletConfig struct {
    ID       map[string]interface{}
    GameName string
    Data     BulletData
}

type BulletData struct {
    Bullets []BulletInfo
}

type BulletInfo struct {
    BulletID int    // Unique bullet identifier
    Name     string // Display name
    Cost     int    // Cost to fire (game currency)
    Damage   int    // Damage value to target
}
```

**Example Document:**
```json
{
  "_id": ObjectId("..."),
  "game_name": "ocean_hunter_v1",
  "data": {
    "bullets": [
      {"bullet_id": 1, "name": "Pea Shot", "cost": 5, "damage": 10},
      {"bullet_id": 2, "name": "Bomb", "cost": 20, "damage": 50},
      {"bullet_id": 3, "name": "Laser", "cost": 50, "damage": 100}
    ]
  }
}
```

---

### GameConfig - Betting & Game Parameters
Defines betting levels, bet limits, and game constraints.

```go
type GameConfig struct {
    ID       map[string]interface{}
    GameName string
    Data     GameConfigData
}

type GameConfigData struct {
    MinBet       int   // Minimum bet per action
    MaxBet       int   // Maximum bet per action
    BetLevels    []int // Available bet amounts
    GameDuration int   // Game duration in seconds
    MaxPlayers   int   // Max players per room
    RoomCapacity int   // Physical room capacity
}
```

**Example Document:**
```json
{
  "_id": ObjectId("..."),
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

---

### GameFeatures - Custom Game Features
Stores special skills, rewards, and multiplier configurations.

```go
type GameFeatures struct {
    ID       map[string]interface{}
    GameName string
    Data     FeaturesData
}

type FeaturesData struct {
    SpecialSkills  []SkillFeature
    SpecialRewards []RewardInfo
    Multipliers    []Multiplier
}

type SkillFeature struct {
    SkillID   int    // Unique skill ID
    SkillName string // Display name
    Cost      int    // Cost to use skill
    Cooldown  int64  // Cooldown in milliseconds
    Effect    string // Effect description/code
}

type RewardInfo struct {
    RewardID   int    // Unique reward ID
    RewardName string // Reward name
    Amount     int    // Reward amount
    Chance     int    // Probability percentage
}

type Multiplier struct {
    FishType   int // Fish type ID
    Multiplier int // Multiplier value
}
```

**Example Document:**
```json
{
  "_id": ObjectId("..."),
  "game_name": "ocean_hunter_v1",
  "data": {
    "special_skills": [
      {
        "skill_id": 1,
        "skill_name": "Double Damage",
        "cost": 100,
        "cooldown": 5000,
        "effect": "damage_x2_10s"
      },
      {
        "skill_id": 2,
        "skill_name": "Slow Effect",
        "cost": 80,
        "cooldown": 8000,
        "effect": "fish_speed_x0.5_5s"
      }
    ],
    "special_rewards": [
      {
        "reward_id": 1,
        "reward_name": "Jackpot",
        "amount": 10000,
        "chance": 5
      }
    ],
    "multipliers": [
      {"fish_type": 5, "multiplier": 3},
      {"fish_type": 6, "multiplier": 5}
    ]
  }
}
```

---

### GamePaths - Fish Movement Paths
Defines predefined movement patterns for fish.

```go
type GamePaths struct {
    ID       map[string]interface{}
    GameName string
    Data     PathData
}

type PathData struct {
    Paths []PathInfo
}

type PathInfo struct {
    PathID      int           // Unique path identifier
    PathName    string        // Display name
    Coordinates []Coordinate  // Movement waypoints
    Duration    int           // Path duration in milliseconds
}

type Coordinate struct {
    X float64 // X position
    Y float64 // Y position
    Z float64 // Z position (depth/layer)
}
```

**Example Document:**
```json
{
  "_id": ObjectId("..."),
  "game_name": "ocean_hunter_v1",
  "data": {
    "paths": [
      {
        "path_id": 1,
        "path_name": "Top to Bottom",
        "coordinates": [
          {"x": 0, "y": 100, "z": 0},
          {"x": 0, "y": 50, "z": 0},
          {"x": 0, "y": 0, "z": 0}
        ],
        "duration": 5000
      },
      {
        "path_id": 2,
        "path_name": "Diagonal Sweep",
        "coordinates": [
          {"x": 0, "y": 100, "z": 0},
          {"x": 50, "y": 75, "z": 0},
          {"x": 100, "y": 50, "z": 0}
        ],
        "duration": 8000
      }
    ]
  }
}
```

---

### GameRTP - Return to Player Configuration
Defines Return-to-Player rates for probability calculations.

```go
type GameRTP struct {
    ID       map[string]interface{}
    GameName string
    Data     RTPData
}

type RTPData struct {
    RTPRate      int       // Overall RTP percentage
    FishRTPMap   map[int]int // Fish ID → RTP percentage
    BulletRTPMap map[int]int // Bullet ID → RTP percentage
}
```

**Example Document:**
```json
{
  "_id": ObjectId("..."),
  "game_name": "ocean_hunter_v1",
  "data": {
    "rtp_rate": 96,
    "fish_rtp_map": {
      "1": 92,
      "2": 94,
      "3": 96,
      "4": 97,
      "5": 98,
      "6": 99
    },
    "bullet_rtp_map": {
      "1": 95,
      "2": 96,
      "3": 97
    }
  }
}
```

---

### GameFishTypes - Fish Type Definitions
Defines all fish species with their properties.

```go
type GameFishTypes struct {
    ID       map[string]interface{}
    GameName string
    Data     FishTypeData
}

type FishTypeData struct {
    FishTypes []FishType
}

type FishType struct {
    FishID     int    // Unique fish type ID
    FishName   string // Display name
    HP         int    // Health points
    BaseReward int    // Base reward amount
    Rarity     string // Rarity: common/uncommon/rare/epic/legendary
    SpawnRate  int    // Spawn probability percentage
    Multiplier int    // Reward multiplier
}
```

**Example Document:**
```json
{
  "_id": ObjectId("..."),
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
        "fish_name": "Shark",
        "hp": 100,
        "base_reward": 500,
        "rarity": "rare",
        "spawn_rate": 10,
        "multiplier": 3
      },
      {
        "fish_id": 4,
        "fish_name": "Dragon Fish",
        "hp": 200,
        "base_reward": 2000,
        "rarity": "epic",
        "spawn_rate": 3,
        "multiplier": 5
      },
      {
        "fish_id": 5,
        "fish_name": "Legendary Creature",
        "hp": 500,
        "base_reward": 10000,
        "rarity": "legendary",
        "spawn_rate": 1,
        "multiplier": 10
      }
    ]
  }
}
```

## Usage in Game Logic

### Bullet Selection
```go
bulletConfig, _ := usecase.GetBulletConfig(ctx, "ocean_hunter_v1")
bullet := bulletConfig.Data.Bullets[bulletID-1]
playerBalance -= bullet.Cost
fishHP -= bullet.Damage
```

### Fish Spawning
```go
fishTypes, _ := usecase.GetGameFishTypes(ctx, "ocean_hunter_v1")
randomFish := selectBySpawnRate(fishTypes.Data.FishTypes)
```

### Skill Activation
```go
features, _ := usecase.GetGameFeatures(ctx, "ocean_hunter_v1")
skill := features.Data.SpecialSkills[skillIndex]
if playerBalance >= skill.Cost {
  applyEffect(skill.Effect)
}
```

### Reward Calculation
```go
rtp, _ := usecase.GetGameRTP(ctx, "ocean_hunter_v1")
fishRTPRate := rtp.Data.FishRTPMap[fishID]
reward := calculateReward(fishType, fishRTPRate)
```

## Redis Cache Format

All objects are cached as JSON strings in Redis with format:
```
Key: game_config:{type}:{gameName}
Value: JSON-serialized object
TTL: 86400 seconds (24 hours)
```

Example Redis entry:
```
Key: game_config:bullets:ocean_hunter_v1
Value: {"_id":..., "game_name":"ocean_hunter_v1", "data":{...}}
```
