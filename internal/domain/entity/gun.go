package entity

type (
	Gun struct {
		GunID      int `json:"gun_id" bson:"gun_id"`
		BulletCost int `json:"bullet_cost" bson:"bullet_cost"`
		Damage     int `json:"damage" bson:"damage"`
		FireRateMs int `json:"fire_rate_ms" bson:"fire_rate_ms"`
	}
)
