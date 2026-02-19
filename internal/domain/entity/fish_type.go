package entity

type (
	FishType struct {
		FishID  int     `json:"fish_id" bson:"fish_id"`
		BaseHP  int     `json:"base_hp" bson:"base_hp"`
		Reward  int     `json:"reward" bson:"reward"`
		HitRate float64 `json:"hit_rate" bson:"hit_rate"`
		Speed   float64 `json:"speed" bson:"speed"`
		IsBoss  bool    `json:"is_boss" bson:"is_boss"`
	}
)
