package entity

type (
	Shot struct {
		BulletID string `json:"bullet_id" bson:"bullet_id"`
		PlayerID string `json:"player_id" bson:"player_id"`
		FishUID  string `json:"fish_uid" bson:"fish_uid"`
		GunID    int    `json:"gun_id" bson:"gun_id"`
		FireTime int64  `json:"fire_time" bson:"fire_time"`
	}
)
