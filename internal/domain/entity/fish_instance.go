package entity

type (
	FishInstance struct {
		FishUID   string `json:"fish_uid" bson:"fish_uid"`
		FishID    int    `json:"fish_id" bson:"fish_id"`
		HP        int    `json:"hp" bson:"hp"`
		SpawnTime int64  `json:"spawn_time" bson:"spawn_time"`
		PathID    int    `json:"path_id" bson:"path_id"`
		Alive     bool   `json:"alive" bson:"alive"`
	}
)
