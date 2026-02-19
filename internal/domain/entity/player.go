package entity

type (
	Player struct {
		PlayerID     string `json:"player_id" bson:"player_id"`
		Balance      int64  `json:"balance" bson:"balance"`
		SeatID       int    `json:"seat_id" bson:"seat_id"`
		GunID        int    `json:"gun_id" bson:"gun_id"`
		RoomID       string `json:"room_id" bson:"room_id"`
		SessionID    string `json:"session_id" bson:"session_id"`
		IsOnline     bool   `json:"is_online" bson:"is_online"`
		LastActionAt int64  `json:"last_action_at" bson:"last_action_at"`
	}
)
