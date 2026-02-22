package gameBaseModels

type GameConfigDocument struct {
	ID       map[string]interface{} `json:"_id" bson:"_id"`
	GameName string                 `json:"game_name" bson:"game_name"`
	Data     interface{}            `json:"data" bson:"data"`
}

// Bullet Configuration
type BulletConfig struct {
	ID       map[string]interface{} `json:"_id" bson:"_id"`
	GameName string                 `json:"game_name" bson:"game_name"`
	Data     BulletData             `json:"data" bson:"data"`
}

type BulletData struct {
	Bullets []BulletInfo `json:"bullets" bson:"bullets"`
}

type BulletInfo struct {
	BulletID int    `json:"bullet_id" bson:"bullet_id"`
	Name     string `json:"name" bson:"name"`
	Cost     int    `json:"cost" bson:"cost"`
	Damage   int    `json:"damage" bson:"damage"`
}

// Game Config - Betting levels and general parameters
type GameConfig struct {
	ID       map[string]interface{} `json:"_id" bson:"_id"`
	GameName string                 `json:"game_name" bson:"game_name"`
	Data     GameConfigData         `json:"data" bson:"data"`
}

type GameConfigData struct {
	MinBet       int   `json:"min_bet" bson:"min_bet"`
	MaxBet       int   `json:"max_bet" bson:"max_bet"`
	BetLevels    []int `json:"bet_levels" bson:"bet_levels"`
	GameDuration int   `json:"game_duration" bson:"game_duration"` // in seconds
	MaxPlayers   int   `json:"max_players" bson:"max_players"`
	RoomCapacity int   `json:"room_capacity" bson:"room_capacity"`
}

// Features - Custom features per game
type GameFeatures struct {
	ID       map[string]interface{} `json:"_id" bson:"_id"`
	GameName string                 `json:"game_name" bson:"game_name"`
	Data     FeaturesData           `json:"data" bson:"data"`
}

type FeaturesData struct {
	SpecialSkills  []SkillFeature `json:"special_skills" bson:"special_skills"`
	SpecialRewards []RewardInfo   `json:"special_rewards" bson:"special_rewards"`
	Multipliers    []Multiplier   `json:"multipliers" bson:"multipliers"`
}

type SkillFeature struct {
	SkillID   int    `json:"skill_id" bson:"skill_id"`
	SkillName string `json:"skill_name" bson:"skill_name"`
	Cost      int    `json:"cost" bson:"cost"`
	Cooldown  int64  `json:"cooldown" bson:"cooldown"` // in milliseconds
	Effect    string `json:"effect" bson:"effect"`
}

type RewardInfo struct {
	RewardID   int    `json:"reward_id" bson:"reward_id"`
	RewardName string `json:"reward_name" bson:"reward_name"`
	Amount     int    `json:"amount" bson:"amount"`
	Chance     int    `json:"chance" bson:"chance"` // percentage
}

type Multiplier struct {
	FishType   int `json:"fish_type" bson:"fish_type"`
	Multiplier int `json:"multiplier" bson:"multiplier"`
}

// Paths - Fish paths
type GamePaths struct {
	ID       map[string]interface{} `json:"_id" bson:"_id"`
	GameName string                 `json:"game_name" bson:"game_name"`
	Data     PathData               `json:"data" bson:"data"`
}

type PathData struct {
	Paths []PathInfo `json:"paths" bson:"paths"`
}

type PathInfo struct {
	PathID      int          `json:"path_id" bson:"path_id"`
	PathName    string       `json:"path_name" bson:"path_name"`
	Coordinates []Coordinate `json:"coordinates" bson:"coordinates"`
	Duration    int          `json:"duration" bson:"duration"` // in milliseconds
}

type Coordinate struct {
	X float64 `json:"x" bson:"x"`
	Y float64 `json:"y" bson:"y"`
	Z float64 `json:"z" bson:"z"`
}

// RTP - Return to Player
type GameRTP struct {
	ID       map[string]interface{} `json:"_id" bson:"_id"`
	GameName string                 `json:"game_name" bson:"game_name"`
	Data     RTPData                `json:"data" bson:"data"`
}

type RTPData struct {
	RTPRate      int         `json:"rtp_rate" bson:"rtp_rate"`             // percentage
	FishRTPMap   map[int]int `json:"fish_rtp_map" bson:"fish_rtp_map"`     // fish_id -> rtp percentage
	BulletRTPMap map[int]int `json:"bullet_rtp_map" bson:"bullet_rtp_map"` // bullet_id -> rtp percentage
}

// Fish Types
type GameFishTypes struct {
	ID       map[string]interface{} `json:"_id" bson:"_id"`
	GameName string                 `json:"game_name" bson:"game_name"`
	Data     FishTypeData           `json:"data" bson:"data"`
}

type FishTypeData struct {
	FishTypes []FishType `json:"fish_types" bson:"fish_types"`
}

type FishType struct {
	FishID     int    `json:"fish_id" bson:"fish_id"`
	FishName   string `json:"fish_name" bson:"fish_name"`
	HP         int    `json:"hp" bson:"hp"`
	BaseReward int    `json:"base_reward" bson:"base_reward"`
	Rarity     string `json:"rarity" bson:"rarity"`         // common, uncommon, rare, epic, legendary
	SpawnRate  int    `json:"spawn_rate" bson:"spawn_rate"` // percentage
	Multiplier int    `json:"multiplier" bson:"multiplier"`
}
