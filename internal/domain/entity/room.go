package entity

type (
	Room struct {
		RoomID   string                   `json:"room_id" bson:"room_id"`
		Status   string                   `json:"status" bson:"status"`
		Players  map[string]*Player       `json:"players" bson:"players"`
		FishMap  map[string]*FishInstance `json:"fish_map" bson:"fish_map"`
		Config   RoomConfig               `json:"config" bson:"config"`
		RTPState RTPState                 `json:"rtp_state" bson:"rtp_state"`
	}
	RoomConfig struct {
		MaxPlayers int `json:"max_players" bson:"max_players"`
	}
)
