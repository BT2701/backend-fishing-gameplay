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

func (f *FishInstance) TakeDamage(damage int) {
	if !f.Alive {
		return
	}
	f.HP -= damage
	if f.HP <= 0 {
		f.Alive = false
		f.HP = 0
	}
}

func (f *FishInstance) IsDead() bool {
	return !f.Alive || f.HP <= 0
}

func (f *FishInstance) IsAlive() bool {
	return f.Alive && f.HP > 0
}
