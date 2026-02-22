package entity

import apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"

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

type RoomStatus string

const (
	RoomStatusOpen    RoomStatus = "open"
	RoomStatusRunning RoomStatus = "running"
	RoomStatusClosed  RoomStatus = "closed"
)

func (r *Room) IsValid() (ok bool, err error) {
	if r.RoomID == "" {
		return false, apperr.ErrInvalidRoomID
	}
	if r.Config.MaxPlayers <= 0 {
		return false, apperr.ErrInvalidMaxPlayers
	}
	return true, nil
}

func (r *Room) IsFull() bool {
	if r.Players == nil {
		return false
	}
	return len(r.Players) >= r.Config.MaxPlayers
}

func (r *Room) HasPlayer(playerID string) bool {
	if r.Players == nil {
		return false
	}
	_, exists := r.Players[playerID]
	return exists
}

func (r *Room) HasFish(fishUID string) bool {
	if r.FishMap == nil {
		return false
	}
	_, exists := r.FishMap[fishUID]
	return exists
}

func (r *Room) GetAliveFishCount() int {
	if r.FishMap == nil {
		return 0
	}
	count := 0
	for _, fish := range r.FishMap {
		if fish.Alive && fish.HP > 0 {
			count++
		}
	}
	return count
}
